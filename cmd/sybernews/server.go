package server

import (
	"net/http"

	"github.com/go-park-mail-ru/2021_2_SaberDevs/internal/router"
	"github.com/labstack/echo/v4/middleware"

	"github.com/labstack/echo/v4"
)

func Run(address string) {
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"http://87.228.2.178:8080"},
		AllowMethods:     []string{http.MethodGet, http.MethodPost},
		AllowCredentials: true,
	}))

	router.Router(e)

	e.Logger.Fatal(e.Start(address))
}
