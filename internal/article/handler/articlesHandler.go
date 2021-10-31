package handlers

import (
	"net/http"

	amodels "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/article/models"
	sbErr "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/syberErrors"
	"github.com/labstack/echo/v4"
	"github.com/microcosm-cc/bluemonday"
	"github.com/pkg/errors"
)

type ArticlesHandler struct {
	UseCase amodels.ArticleUsecase
}

// NewArticleHandler will initialize the articles/ resources endpoint
func NewArticlesHandler(e *echo.Echo, us amodels.ArticleUsecase) ArticlesHandler {
	handler := &ArticlesHandler{
		UseCase: us,
	}
	// e.GET("/articles", handler.GetFeed)
	return *handler
}

const chunkSize = 5

func SanitizeArticle(a *amodels.Article) *amodels.Article {
	// s := bluemonday.NewPolicy()
	//s.AllowStandardURLs()
	s := bluemonday.StrictPolicy()
	l := bluemonday.UGCPolicy()
	a.AuthorAvatar = s.Sanitize(a.AuthorAvatar)
	a.AuthorName = s.Sanitize(a.AuthorName)
	a.AuthorUrl = s.Sanitize(a.AuthorUrl)
	//a.Comments = s.Sanitize(a.Comments) //not a string
	a.CommentsUrl = s.Sanitize(a.CommentsUrl)
	a.Id = s.Sanitize(a.Id)
	// a.Likes = s.Sanitize(a.Likes)//not a string
	a.PreviewUrl = s.Sanitize(a.PreviewUrl)
	for i := range a.Tags {
		a.Tags[i] = l.Sanitize(a.Tags[i])
	}
	a.Text = s.Sanitize(a.Text)
	a.Title = s.Sanitize(a.Title)
	return a
}

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
	newArticle = SanitizeArticle(newArticle)
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
	newArticle = SanitizeArticle(newArticle)
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
