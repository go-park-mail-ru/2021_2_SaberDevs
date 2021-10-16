package server

import (
	"net/http"

	ahandler "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/article/delivery"
	usecase "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/article/usecase"
	uhandler "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/user/handler"

	"github.com/labstack/echo/v4/middleware"

	echo "github.com/labstack/echo/v4"
)

func router(e *echo.Echo) {
	userApi := uhandler.NewUserHandler()
	us := usecase.NewArticleUsecase()
	articlesApi := ahandler.NewArticlesHandler(e, us)
	e.GET("/feed", articlesApi.GetFeed)

	e.POST("/login", userApi.Login)
	e.POST("/signup", userApi.Register)
	e.POST("/logout", userApi.Logout)

}

func Run(address string) {
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"http://localhost:8080"},
		AllowMethods:     []string{http.MethodGet, http.MethodPost},
		AllowCredentials: true,
	}))

	router(e)

	e.Logger.Fatal(e.Start(address))
}
