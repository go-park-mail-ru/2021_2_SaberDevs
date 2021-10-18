package server

import (
	ahandler "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/article/handler"
	srepo "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/session/repository"
	uhandler "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/user/handler"
	urepo "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/user/repisitory"
	uuscase "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/user/usecase"
	"net/http"

	"github.com/labstack/echo/v4/middleware"

	"github.com/labstack/echo/v4"
)

func router(e *echo.Echo) {
	userRepo := urepo.NewUserRepository()
	sessionRepo := srepo.NewSessionRepository()

	userUsecase := uuscase.NewUserUsecase(userRepo, sessionRepo)
	userApi := uhandler.NewUserHandler(userUsecase)

	articlesApi := ahandler.NewArticlesHandler()

	e.POST("/login", userApi.Login)
	e.POST("/signup", userApi.Register)
	e.POST("/logout", userApi.Logout)

	e.GET("/feed", articlesApi.Getfeed)
}

func Run(address string) {
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"http://87.228.2.178:8080"},
		AllowMethods:     []string{http.MethodGet, http.MethodPost},
		AllowCredentials: true,
	}))

	router(e)

	e.Logger.Fatal(e.Start(address))
}
