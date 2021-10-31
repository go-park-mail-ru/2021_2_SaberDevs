package server

import (
	ahandler "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/article/handler"
	ausecase "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/article/usecase"
	syberMiddleware "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/middleware"
	shandler "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/session/handler"
	srepo "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/session/repository"
	susecase "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/session/usecase"
	uhandler "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/user/handler"
	urepo "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/user/repisitory"
	uusecase "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/user/usecase"
	"github.com/jmoiron/sqlx"
	"github.com/tarantool/go-tarantool"

	"net/http"

	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"

	"github.com/labstack/echo/v4"
)

func DbConnect() (*sqlx.DB, error) {
	connStr := "user=postgres dbname=postgres password=yura11011 host=localhost sslmode=disable"
	db, err := sqlx.Open("postgres", connStr)
	if err != nil {
		return db, err
	}
	err = db.Ping()
	if err != nil {
		return db, err
	}
	return db, err
}

func DbClose(db *sqlx.DB) error {
	err := db.Close()
	if err != nil {
		return err
	}
	return err
}

func router(e *echo.Echo) {

	// us := ausecase.NewArticleUsecase()
	// articlesAPI := ahandler.NewArticlesHandler(e, us)

}

func Run(address string) {
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"http://127.0.0.1:8080", "http://87.228.2.178:8080"},
		AllowMethods:     []string{http.MethodGet, http.MethodPost},
		AllowCredentials: true,
	}))

	e.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{
		StackSize: 1 << 10, // 1 KB
		LogLevel:  log.ERROR,
	}))

	db, err := DbConnect()
	if err != nil {
		e.Logger.Fatal(err)
	}
	opts := tarantool.Opts{User: "admin", Pass: "pass"}
	sessionsDbConn, err := tarantool.Connect(":3302", opts)
	if err != nil {
		panic("error connetcting to session DB: " + err.Error())
	}

	_, err = sessionsDbConn.Ping()
	if err != nil {
		panic("error pinging session DB: " + err.Error())
	}

	userRepo := urepo.NewUserRepository(db)
	sessionRepo := srepo.NewSessionRepository(sessionsDbConn)
	userUsecase := uusecase.NewUserUsecase(userRepo, sessionRepo)
	userAPI := uhandler.NewUserHandler(userUsecase)

	sessionUsecase := susecase.NewsessionUsecase(userRepo, sessionRepo)
	sessionAPI := shandler.NewSessionHandler(sessionUsecase)

	// e.Use(syberMiddleware.ValidateRequestBody)
	e.HTTPErrorHandler = syberMiddleware.ErrorHandler
	e.Use(syberMiddleware.AddId)
	e.Use(syberMiddleware.AccessLogger)
	e.Logger.SetLevel(log.INFO)
	e.POST("/login", userAPI.Login)
	e.POST("/signup", userAPI.Register)
	e.POST("/logout", userAPI.Logout)
	e.POST("/", sessionAPI.CheckSession)
	articles := e.Group("/feed")
	//articles.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{}))
	us := ausecase.NewArticleUsecase(db)
	articlesAPI := ahandler.NewArticlesHandler(e, us)

	articles.GET("", articlesAPI.GetFeed)
	articles.POST("/create", articlesAPI.Create)
	articles.POST("/update", articlesAPI.Update)
	articles.DELETE("/delete", articlesAPI.Delete)
	defer DbClose(db)
	// e.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{
	// 	TokenLookup: "header:X-XSRF-TOKEN",
	// }))

	router(e)

	e.Logger.Fatal(e.Start(address))
}
