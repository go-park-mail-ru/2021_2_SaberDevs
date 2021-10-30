package server

import (
	syberMiddleware "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/middleware"
	shandler "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/session/handler"
	srepo "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/session/repository"
	susecase "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/session/usecase"
	uhandler "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/user/handler"
	urepo "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/user/repisitory"
	uusecase "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/user/usecase"
	"github.com/tarantool/go-tarantool"

	"net/http"

	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"

	"github.com/labstack/echo/v4"
)

func router(e *echo.Echo) {
	// us := ausecase.NewArticleUsecase()
	// articlesAPI := ahandler.NewArticlesHandler(e, us)

	opts := tarantool.Opts{User: "admin", Pass: "pass"}
	sessionsDbConn, err := tarantool.Connect(":3302", opts)
	if err != nil {
		panic("error connetcting to session DB: " + err.Error())
	}

	_, err = sessionsDbConn.Ping()
	if err != nil {
		panic("error pinging session DB: " + err.Error())
	}

	userRepo := urepo.NewUserRepository()
	sessionRepo := srepo.NewSessionRepository(sessionsDbConn)

	userUsecase := uusecase.NewUserUsecase(userRepo, sessionRepo)
	userAPI := uhandler.NewUserHandler(userUsecase)

	sessionUsecase := susecase.NewsessionUsecase(userRepo, sessionRepo)
	sessionAPI := shandler.NewSessionHandler(sessionUsecase)

	// e.Use(syberMiddleware.ValidateRequestBody)
	e.HTTPErrorHandler = syberMiddleware.ErrorHandler

	// e.GET("/feed", articlesAPI.GetFeed)

	e.POST("/login", userAPI.Login)
	e.POST("/signup", userAPI.Register)
	e.POST("/logout", userAPI.Logout)

	e.POST("/", sessionAPI.CheckSession)

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

	// e.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{
	// 	TokenLookup: "header:X-XSRF-TOKEN",
	// }))

	router(e)

	e.Logger.Fatal(e.Start(address))
}
