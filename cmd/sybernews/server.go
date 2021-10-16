package server

import (
	"net/http"

	ahandler "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/article/delivery"
	uhandler "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/user/handler"

	"github.com/labstack/echo/v4/middleware"

	"github.com/labstack/echo/v4"
)

func router(e *echo.Echo) {
	userApi := uhandler.NewUserHandler()
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
