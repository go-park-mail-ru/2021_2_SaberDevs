package handlers

import (
	"fmt"
	"net/http"
	"strconv"

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
func NewArticlesHandler(us amodels.ArticleUsecase) ArticlesHandler {
	handler := &ArticlesHandler{
		UseCase: us,
	}
	return *handler
}

const del = "DELETED"
const up = "UPDATED"
const chunkSize = 5

func SanitizeArticle(a *amodels.Article) *amodels.Article {
	s := bluemonday.StrictPolicy()
	l := bluemonday.UGCPolicy()
	a.AuthorAvatar = s.Sanitize(a.AuthorAvatar)
	a.AuthorName = s.Sanitize(a.AuthorName)
	a.AuthorUrl = s.Sanitize(a.AuthorUrl)
	a.CommentsUrl = s.Sanitize(a.CommentsUrl)
	a.Id = s.Sanitize(a.Id)
	a.PreviewUrl = s.Sanitize(a.PreviewUrl)
	for i := range a.Tags {
		a.Tags[i] = l.Sanitize(a.Tags[i])
	}
	a.Text = s.Sanitize(a.Text)
	a.Title = s.Sanitize(a.Title)
	return a
}
func SanitizeCreate(a *amodels.ArticleCreate) *amodels.ArticleCreate {
	s := bluemonday.StrictPolicy()
	l := bluemonday.UGCPolicy()
	for i := range a.Tags {
		a.Tags[i] = l.Sanitize(a.Tags[i])
	}
	a.Text = s.Sanitize(a.Text)
	a.Title = s.Sanitize(a.Title)
	return a
}
func SanitizeUpdate(a *amodels.ArticleUpdate) *amodels.ArticleUpdate {
	s := bluemonday.StrictPolicy()
	l := bluemonday.UGCPolicy()
	a.Id = s.Sanitize(a.Id)
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

func (api *ArticlesHandler) GetByID(c echo.Context) error {
	strId := c.QueryParam("id")
	ctx := c.Request().Context()
	if strId == "" {
		strId = "0"
	}
	if strId == "end" {
		strId = "12"
	}

	id, err := strconv.Atoi(strId)
	if err != nil {
		return errors.Wrap(err, "articleHandler/getbyid")
	}
	Data, err := api.UseCase.GetByID(ctx, int64(id))
	if err != nil {
		return errors.Wrap(err, "articlesHandler/GetbyID")
	}
	response := Data

	return c.JSON(http.StatusOK, response)
}

func (api *ArticlesHandler) GetByAuthor(c echo.Context) error {
	login := c.QueryParam("login")
	ctx := c.Request().Context()
	ChunkData, err := api.UseCase.GetByAuthor(ctx, login)
	if err != nil {
		return errors.Wrap(err, "articlesHandler/GetByAuthor")
	}
	response := amodels.ChunkResponse{
		Status:    http.StatusOK,
		ChunkData: ChunkData,
	}
	return c.JSON(http.StatusOK, response)
}

func (api *ArticlesHandler) Update(c echo.Context) error {
	newArticle := new(amodels.ArticleUpdate)
	err := c.Bind(newArticle)
	if err != nil {
		return sbErr.ErrUnpackingJSON{
			Reason:   err.Error(),
			Function: "articlesHandler/Update",
		}
	}
	newArticle = SanitizeUpdate(newArticle)
	ctx := c.Request().Context()
	err = api.UseCase.Update(ctx, newArticle)
	if err != nil {
		return errors.Wrap(err, "articlesHandler/Update")
	}

	response := up
	return c.JSON(http.StatusOK, response)
}

func (api *ArticlesHandler) GetByTag(c echo.Context) error {
	tag := c.QueryParam("tag")
	ctx := c.Request().Context()
	ChunkData, err := api.UseCase.GetByTag(ctx, tag)
	if err != nil {
		return errors.Wrap(err, "articlesHandler/GetByTag")
	}
	response := amodels.ChunkResponse{
		Status:    http.StatusOK,
		ChunkData: ChunkData,
	}
	return c.JSON(http.StatusOK, response)
}

func (api *ArticlesHandler) Create(c echo.Context) error {
	tempArticle := new(amodels.ArticleCreate)
	err := c.Bind(tempArticle)
	if err != nil {
		return sbErr.ErrUnpackingJSON{
			Reason:   err.Error(),
			Function: "articlesHandler/Create",
		}
	}
	cookie, err := c.Cookie("session")
	if err != nil {
		return sbErr.ErrAuthorised{
			Reason:   err.Error(),
			Function: "articlesHandler/Create",
		}
	}
	tempArticle = SanitizeCreate(tempArticle)
	ctx := c.Request().Context()
	Id, err := api.UseCase.Store(ctx, cookie, tempArticle)
	if err != nil {
		return errors.Wrap(err, "articlesHandler/Create")
	}

	response := fmt.Sprint(Id)
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
	response := del
	return c.JSON(http.StatusOK, response)
}
