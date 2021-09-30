package server

import (
	"net/http"

	"github.com/labstack/echo/v4/middleware"
	"saverdevs.com/2021_2_Saberdevs/server/handlers"

	"github.com/labstack/echo/v4"
)

func router(e *echo.Echo) {
	api := handlers.NewMyHandler()

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
