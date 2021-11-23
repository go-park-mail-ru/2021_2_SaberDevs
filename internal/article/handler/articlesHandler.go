package handlers

import (
	"fmt"
	"net/http"
	"regexp"
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

func IdToStr(strId string) (int, error) {
	if strId == "" {
		strId = "0"
	}
	id, err := strconv.Atoi(strId)
	return id, err
}

func SanitizeArticle(a *amodels.Article) *amodels.Article {
	s := bluemonday.StrictPolicy()
	l := bluemonday.UGCPolicy()
	a.AuthorAvatar = s.Sanitize(a.AuthorAvatar)
	a.AuthorName = s.Sanitize(a.AuthorName)
	a.AuthorUrl = s.Sanitize(a.AuthorUrl)
	a.CommentsUrl = s.Sanitize(a.CommentsUrl)
	a.PreviewUrl = s.Sanitize(a.PreviewUrl)
	for i := range a.Tags {
		a.Tags[i] = l.Sanitize(a.Tags[i])
	}
	a.Text = s.Sanitize(a.Text)
	a.Title = s.Sanitize(a.Title)
	r := regexp.MustCompile("\\s+")
	a.Title = r.ReplaceAllString(a.Title, " ")

	return a
}
func SanitizeCreate(a *amodels.ArticleCreate) *amodels.ArticleCreate {
	s := bluemonday.StrictPolicy()
	l := bluemonday.UGCPolicy()
	for i := range a.Tags {
		a.Tags[i] = l.Sanitize(a.Tags[i])
	}
	a.Category = s.Sanitize(a.Category)
	a.Img = s.Sanitize(a.Img)
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
	a.Category = s.Sanitize(a.Category)
	a.Img = s.Sanitize(a.Img)
	a.Text = s.Sanitize(a.Text)
	a.Title = s.Sanitize(a.Title)
	return a
}

func (api *ArticlesHandler) GetFeed(c echo.Context) error {
	id := c.QueryParam("idLastLoaded")
	ctx := c.Request().Context()
	ChunkData, err := api.UseCase.Fetch(ctx, id, chunkSize)
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
	id, err := IdToStr(strId)
	if err != nil {
		return errors.Wrap(err, "articleHandler/getbyid")
	}
	Data, err := api.UseCase.GetByID(ctx, int64(id))
	if err != nil {
		return errors.Wrap(err, "articlesHandler/GetbyID")
	}
	response := amodels.ArticleResponse{
		Status: http.StatusOK,
		Data:   Data,
	}

	return c.JSON(http.StatusOK, response)
}

func (api *ArticlesHandler) GetByAuthor(c echo.Context) error {
	login := c.QueryParam("login")
	id := c.QueryParam("idLastLoaded")
	ctx := c.Request().Context()
	ChunkData, err := api.UseCase.GetByAuthor(ctx, login, id, chunkSize)
	if err != nil {
		return errors.Wrap(err, "articlesHandler/GetByAuthor")
	}
	response := amodels.ChunkResponse{
		Status:    http.StatusOK,
		ChunkData: ChunkData,
	}
	return c.JSON(http.StatusOK, response)
}
func (api *ArticlesHandler) GetByCategory(c echo.Context) error {
	id := c.QueryParam("idLastLoaded")
	cat := c.QueryParam("category")
	c.Logger().Info("!!!!!!!!!!!!!Id = ", id)
	ctx := c.Request().Context()
	ChunkData, err := api.UseCase.GetByCategory(ctx, cat, id, chunkSize)
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
	if (err != nil) || (newArticle == new(amodels.ArticleUpdate)) {
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

	response := amodels.GenericResponse{
		Status: http.StatusOK,
		Data:   up,
	}

	return c.JSON(http.StatusOK, response)
}

func (api *ArticlesHandler) GetByTag(c echo.Context) error {
	tag := c.QueryParam("tag")
	id := c.QueryParam("idLastLoaded")
	ctx := c.Request().Context()
	ChunkData, err := api.UseCase.GetByTag(ctx, tag, id, chunkSize)
	if err != nil {
		return errors.Wrap(err, "articlesHandler/GetByTag")
	}
	response := amodels.ChunkResponse{
		Status:    http.StatusOK,
		ChunkData: ChunkData,
	}
	return c.JSON(http.StatusOK, response)
}

func (api *ArticlesHandler) FindArticles(c echo.Context) error {
	q := c.QueryParam("q")
	id := c.QueryParam("idLastLoaded")
	s := bluemonday.StrictPolicy()
	id = s.Sanitize(id)
	q = s.Sanitize(q)
	ctx := c.Request().Context()
	ChunkData, err := api.UseCase.FindArticles(ctx, q, id, chunkSize)
	if err != nil {
		return errors.Wrap(err, "articlesHandler/GetByTag")
	}
	response := amodels.ChunkResponse{
		Status:    http.StatusOK,
		ChunkData: ChunkData,
	}
	return c.JSON(http.StatusOK, response)
}

func (api *ArticlesHandler) FindAuthors(c echo.Context) error {
	q := c.QueryParam("q")
	id := c.QueryParam("idLastLoaded")
	s := bluemonday.StrictPolicy()
	id = s.Sanitize(id)
	q = s.Sanitize(q)
	ctx := c.Request().Context()
	ChunkData, err := api.UseCase.FindAuthors(ctx, q, id, chunkSize)
	if err != nil {
		return errors.Wrap(err, "articlesHandler/GetByTag")
	}
	response := amodels.AuthorsChunks{
		Status:    http.StatusOK,
		ChunkData: ChunkData,
	}
	return c.JSON(http.StatusOK, response)
}

func (api *ArticlesHandler) FindByTag(c echo.Context) error {
	q := c.QueryParam("q")
	id := c.QueryParam("idLastLoaded")
	s := bluemonday.StrictPolicy()
	id = s.Sanitize(id)
	q = s.Sanitize(q)
	ctx := c.Request().Context()
	ChunkData, err := api.UseCase.FindByTag(ctx, q, id, chunkSize)
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

	response := amodels.GenericResponse{
		Status: http.StatusOK,
		Data:   fmt.Sprint(Id),
	}

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

	response := amodels.GenericResponse{
		Status: http.StatusOK,
		Data:   del,
	}
	return c.JSON(http.StatusOK, response)
}
