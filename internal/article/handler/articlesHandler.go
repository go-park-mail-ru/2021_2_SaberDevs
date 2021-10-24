package handlers

import (
	"net/http"

	amodels "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/article/models"
	errResp "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/errResponses"
	"github.com/labstack/echo/v4"
)

type ArticlesHandler struct {
	UseCase amodels.ArticleUseCase
}

// NewArticleHandler will initialize the articles/ resources endpoint
func NewArticlesHandler(e *echo.Echo, us amodels.ArticleUseCase) ArticlesHandler {
	handler := &ArticlesHandler{
		UseCase: us,
	}
	// e.GET("/articles", handler.GetFeed)
	return *handler
}

const chunkSize = 5

func (api *ArticlesHandler) GetFeed(c echo.Context) error {
	rec := c.QueryParam("idLastLoaded")
	ChunkData, err := api.UseCase.Fetch(c, rec, chunkSize)
	if err != nil {
		c.Logger().Printf("Error: %s", err.Error())
		return c.JSON(http.StatusNotFound, errResp.ErrNotFeedNumber)
	}
	// Возвращаем записи

	// формируем ответ
	response := amodels.ChunkResponse{
		Status:    http.StatusOK,
		ChunkData: ChunkData,
	}
	return c.JSON(http.StatusOK, response)
}
