package server

import (
	"net/http"

	"github.com/go-park-mail-ru/2021_2_SaberDevs/server/handlers"
	"github.com/labstack/echo/v4/middleware"

	"github.com/labstack/echo/v4"
)

func router(e *echo.Echo) {
	api := handlers.NewHandler()

	e.POST("/login", api.Login)
	e.POST("/signup", api.Register)
	e.POST("/logout", api.Logout)
	e.GET("/feed", api.Getfeed)
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
