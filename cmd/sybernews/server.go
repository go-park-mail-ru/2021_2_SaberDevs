package server

import (
	srepo "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/session/repository"
	uhandler "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/user/handler"
	urepo "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/user/repisitory"
	uusecase "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/user/usecase"
	susecase "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/session/usecase"
	shandler "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/session/handler"
  
    "net/http"

	ahandler "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/article/handler"
	ausecase "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/article/usecase"

	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"

	"github.com/labstack/echo/v4"
)

func router(e *echo.Echo) {
	us := ausecase.NewArticleUsecase()
	articlesApi := ahandler.NewArticlesHandler(e, us)

    userRepo := urepo.NewUserRepository()
	sessionRepo := srepo.NewSessionRepository()

	userUsecase := uusecase.NewUserUsecase(userRepo, sessionRepo)
	userApi := uhandler.NewUserHandler(userUsecase)

	sessionUsecase := susecase.NewsessionUsecase(userRepo, sessionRepo)
	sessionApi := shandler.NewSessionHandler(sessionUsecase)

	e.GET("/feed", articlesApi.GetFeed)

	e.POST("/login", userApi.Login)
	e.POST("/signup", userApi.Register)
	e.POST("/logout", userApi.Logout)

	e.POST("/", sessionApi.CheckSession)

}

func Run(address string) {
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"http://87.228.2.178:8080"},
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
