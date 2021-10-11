package router

import (
	"github.com/go-park-mail-ru/2021_2_SaberDevs/internal/handlers"
	"github.com/labstack/echo/v4"
)

func Router(e *echo.Echo) {
	userApi := handlers.NewUserHandler()
	articlesApi := handlers.NewArticlesHandler()

	e.POST("/login", userApi.Login)
	e.POST("/signup", userApi.Register)
	e.POST("/logout", userApi.Logout)

	e.GET("/feed", articlesApi.Getfeed)
}
