package handlers

import (
	"net/http"

	amodels "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/article/models"
	sbErr "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/syberErrors"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
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
	ctx := c.Request().Context()
	ChunkData, err := api.UseCase.Fetch(ctx, rec, chunkSize)
	if err != nil {
		return errors.Wrap(err, "articlesHandler/GetFeed")
	}
	// Возвращаем записи

	// формируем ответ
	response := amodels.ChunkResponse{
		Status:    http.StatusOK,
		ChunkData: ChunkData,
	}
	return c.JSON(http.StatusOK, response)
}

func (api *ArticlesHandler) Update(c echo.Context) error {
	newArticle := new(amodels.Article)
	err := c.Bind(newArticle)
	if err != nil {
		return sbErr.ErrUnpackingJSON{
			Reason:   err.Error(),
			Function: "articlesHandler/Update",
		}
	}
	ctx := c.Request().Context()
	err = api.UseCase.Update(ctx, newArticle)
	if err != nil {
		return errors.Wrap(err, "articlesHandler/Update")
	}

	response := "UPDATED"
	return c.JSON(http.StatusOK, response)
}

func (api *ArticlesHandler) Create(c echo.Context) error {
	newArticle := new(amodels.Article)
	err := c.Bind(newArticle)
	if err != nil {
		return sbErr.ErrUnpackingJSON{
			Reason:   err.Error(),
			Function: "articlesHandler/Create",
		}
	}
	ctx := c.Request().Context()
	err = api.UseCase.Store(ctx, newArticle)
	if err != nil {
		return errors.Wrap(err, "articlesHandler/Create")
	}

	response := "CREATED"
	return c.JSON(http.StatusOK, response)
}

func (api *ArticlesHandler) Delete(c echo.Context) error {
	id := c.QueryParam("id")
	ctx := c.Request().Context()
	err := api.UseCase.Delete(ctx, id)
	if err != nil {
		return errors.Wrap(err, "articlesHandler/Delete")
	}
	// формируем ответ
	response := "DELETED"
	return c.JSON(http.StatusOK, response)
}
