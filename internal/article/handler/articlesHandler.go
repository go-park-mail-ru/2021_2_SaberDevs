package handlers

import (
	"context"
	"fmt"
	"net/http"
	"regexp"
	"strconv"

	app "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/article/article_app"
	"github.com/go-park-mail-ru/2021_2_SaberDevs/internal/article/models"
	amodels "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/article/models"
	sbErr "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/syberErrors"
	"github.com/labstack/echo/v4"
	"github.com/microcosm-cc/bluemonday"
	"github.com/pkg/errors"
	"github.com/prometheus/client_golang/prometheus"
	"google.golang.org/grpc/metadata"
)

type ArticlesHandler struct {
	UseCase app.ArticleDeliveryClient
}

var maxNum = "999999"

var Hits = prometheus.NewCounterVec(prometheus.CounterOpts{
	Name: "hits",
}, []string{"layer", "path"})

var Errors = prometheus.NewCounterVec(prometheus.CounterOpts{
	Name: "hits",
}, []string{"status", "path"})

var Duration = prometheus.NewCounterVec(prometheus.CounterOpts{
	Name: "hits",
}, []string{"status", "path"})

var layer = "delivery"

func reverseConv(a *app.Preview) *models.Preview {
	val := new(models.Preview)
	val.Category = a.Category
	val.Comments = uint(a.Comments)
	val.CommentsUrl = a.CommentsUrl
	val.DateTime = a.DateTime
	val.Id = a.Id
	val.Likes = uint(a.Likes)
	val.PreviewUrl = a.PreviewUrl
	val.Tags = a.Tags
	val.Text = a.Text
	val.Title = a.Title
	val.Author.Id = int(a.Author.Id)
	val.Author.AvatarUrl = a.Author.AvatarUrl
	val.Author.Description = a.Author.Description
	val.Author.Email = a.Author.Email
	val.Author.Login = a.Author.Login
	val.Author.Name = a.Author.Name
	val.Author.Password = a.Author.Password
	val.Author.Score = int(a.Author.Score)
	val.Author.Surname = a.Author.Surname
	return val
}

func revFullConv(a *app.FullArticle) *models.FullArticle {
	val := new(models.FullArticle)
	val.Category = a.Category
	val.Comments = uint(a.Comments)
	val.CommentsUrl = a.CommentsUrl
	val.DateTime = a.DateTime
	val.Id = a.Id
	val.Likes = uint(a.Likes)
	val.PreviewUrl = a.PreviewUrl
	val.Tags = a.Tags
	val.Text = a.Text
	val.Title = a.Title
	val.Author.Id = int(a.Author.Id)
	val.Author.AvatarUrl = a.Author.AvatarUrl
	val.Author.Description = a.Author.Description
	val.Author.Email = a.Author.Email
	val.Author.Login = a.Author.Login
	val.Author.Name = a.Author.Name
	val.Author.Password = a.Author.Password
	val.Author.Score = int(a.Author.Score)
	val.Author.Surname = a.Author.Surname
	return val
}

// NewArticleHandler will initialize the articles/ resources endpoint
func NewArticlesHandler(us app.ArticleDeliveryClient) ArticlesHandler {
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
		strId = maxNum
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

func upConv(a *models.ArticleUpdate) *app.ArticleUpdate {
	ar := new(app.ArticleUpdate)
	ar.Category = a.Category
	ar.Img = a.Img
	ar.Tags = a.Tags
	ar.Text = a.Text
	ar.Title = a.Title
	ar.Id = a.Id
	return ar
}

func auConv(a app.Author) *models.Author {
	thor := new(models.Author)
	thor.Id = int(a.Id)
	thor.AvatarUrl = a.AvatarUrl
	thor.Description = a.Description
	thor.Email = a.Email
	thor.Login = a.Login
	thor.Name = a.Name
	thor.Password = a.Password
	thor.Score = int(a.Score)
	thor.Surname = thor.Surname
	return thor
}

func arConv(a *models.ArticleCreate) *app.ArticleCreate {
	ar := new(app.ArticleCreate)
	ar.Category = a.Category
	ar.Img = a.Img
	ar.Tags = a.Tags
	ar.Text = a.Text
	ar.Title = a.Title
	return ar
}

func (api *ArticlesHandler) GetFeed(c echo.Context) error {
	id := c.QueryParam("idLastLoaded")
	fPath := "/api/v1/articles/feed"
	Hits.WithLabelValues(layer, fPath).Inc()
	reqID := c.Request().Header.Get(echo.HeaderXRequestID)
	md := metadata.New(map[string]string{"X-Request-ID": reqID})
	ctx := metadata.NewOutgoingContext(context.Background(), md)
	a := &app.Chunk{ChunkSize: chunkSize, IdLastLoaded: id}
	Data, err := api.UseCase.Fetch(ctx, a)
	if err != nil {
		return errors.Wrap(err, "articlesHandler/GetFeed")
	}
	// Возвращаем записи
	var ChunkData []amodels.Preview
	for _, a := range Data.Preview {
		val := reverseConv(a)
		ChunkData = append(ChunkData, *val)
	}
	// формируем ответ
	response := amodels.ChunkResponse{
		Status:    http.StatusOK,
		ChunkData: ChunkData,
	}
	return c.JSON(http.StatusOK, response)
}

func (api *ArticlesHandler) GetByID(c echo.Context) error {
	strId := c.QueryParam("id")
	fPath := "/api/v1/article/:id"
	Hits.WithLabelValues(layer, fPath).Inc()
	reqID := c.Request().Header.Get(echo.HeaderXRequestID)
	md := metadata.New(map[string]string{"X-Request-ID": reqID})
	ctx := metadata.NewOutgoingContext(context.Background(), md)
	myid := app.Id{Id: strId}
	Data, err := api.UseCase.GetByID(ctx, &myid)
	if err != nil {
		return errors.Wrap(err, "articlesHandler/GetbyID")
	}
	response := amodels.ArticleResponse{
		Status: http.StatusOK,
		Data:   *revFullConv(Data),
	}

	return c.JSON(http.StatusOK, response)
}

func (api *ArticlesHandler) GetByAuthor(c echo.Context) error {
	login := c.QueryParam("login")
	id := c.QueryParam("idLastLoaded")
	fPath := "/api/v1/articles/author"
	Hits.WithLabelValues(layer, fPath).Inc()
	reqID := c.Request().Header.Get(echo.HeaderXRequestID)
	md := metadata.New(map[string]string{"X-Request-ID": reqID})
	ctx := metadata.NewOutgoingContext(context.Background(), md)
	a := &app.Chunk{ChunkSize: chunkSize, IdLastLoaded: id}
	authors := &app.Authors{Author: login, Chunk: a}
	Data, err := api.UseCase.GetByAuthor(ctx, authors)
	if err != nil {
		return errors.Wrap(err, "articlesHandler/GetByAuthor")
	}
	var ChunkData []amodels.Preview
	for _, a := range Data.Preview {
		val := reverseConv(a)
		ChunkData = append(ChunkData, *val)
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
	fPath := "/api/v1/articles/category"
	Hits.WithLabelValues(layer, fPath).Inc()
	if cat == "" {
		return sbErr.ErrNoContent{
			Reason:   "empty Category",
			Function: "articlesHandler/GetByCategory",
		}
	}
	reqID := c.Request().Header.Get(echo.HeaderXRequestID)
	md := metadata.New(map[string]string{"X-Request-ID": reqID})
	ctx := metadata.NewOutgoingContext(context.Background(), md)
	a := &app.Chunk{ChunkSize: chunkSize, IdLastLoaded: id}
	categories := &app.Categories{Category: cat, Chunk: a}
	Data, err := api.UseCase.GetByCategory(ctx, categories)
	if err != nil {
		return errors.Wrap(err, "articlesHandler/GetByAuthor")
	}
	var ChunkData []amodels.Preview
	for _, a := range Data.Preview {
		val := reverseConv(a)
		ChunkData = append(ChunkData, *val)
	}
	response := amodels.ChunkResponse{
		Status:    http.StatusOK,
		ChunkData: ChunkData,
	}
	return c.JSON(http.StatusOK, response)
}

func (api *ArticlesHandler) Update(c echo.Context) error {
	fPath := "/api/v1/articles/update"
	Hits.WithLabelValues(layer, fPath).Inc()
	newArticle := new(amodels.ArticleUpdate)
	err := c.Bind(newArticle)
	if (err != nil) || (newArticle == new(amodels.ArticleUpdate)) {
		return sbErr.ErrUnpackingJSON{
			Reason:   err.Error(),
			Function: "articlesHandler/Update",
		}
	}
	newArticle = SanitizeUpdate(newArticle)
	reqID := c.Request().Header.Get(echo.HeaderXRequestID)
	md := metadata.New(map[string]string{"X-Request-ID": reqID})
	ctx := metadata.NewOutgoingContext(context.Background(), md)
	_, err = api.UseCase.Update(ctx, upConv(newArticle))
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
	fPath := "/api/v1/articles/tag"
	Hits.WithLabelValues(layer, fPath).Inc()
	reqID := c.Request().Header.Get(echo.HeaderXRequestID)
	md := metadata.New(map[string]string{"X-Request-ID": reqID})
	ctx := metadata.NewOutgoingContext(context.Background(), md)
	a := &app.Chunk{ChunkSize: chunkSize, IdLastLoaded: id}
	tags := &app.Tags{Tag: tag, Chunk: a}
	Data, err := api.UseCase.GetByTag(ctx, tags)
	if err != nil {
		return errors.Wrap(err, "articlesHandler/GetByTag")
	}
	var ChunkData []amodels.Preview
	for _, a := range Data.Preview {
		val := reverseConv(a)
		ChunkData = append(ChunkData, *val)
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
	fPath := "/api/v1/articles/query"
	Hits.WithLabelValues(layer, fPath).Inc()
	reqID := c.Request().Header.Get(echo.HeaderXRequestID)
	md := metadata.New(map[string]string{"X-Request-ID": reqID})
	ctx := metadata.NewOutgoingContext(context.Background(), md)
	a := &app.Chunk{ChunkSize: chunkSize, IdLastLoaded: id}
	query := &app.Queries{Query: q, Chunk: a}
	Data, err := api.UseCase.FindArticles(ctx, query)
	if err != nil {
		return errors.Wrap(err, "articlesHandler/GetByTag")
	}
	var ChunkData []amodels.Preview
	for _, a := range Data.Preview {
		val := reverseConv(a)
		ChunkData = append(ChunkData, *val)
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
	fPath := "/api/v1/articles/query"
	Hits.WithLabelValues(layer, fPath).Inc()
	reqID := c.Request().Header.Get(echo.HeaderXRequestID)
	md := metadata.New(map[string]string{"X-Request-ID": reqID})
	ctx := metadata.NewOutgoingContext(context.Background(), md)
	a := &app.Chunk{ChunkSize: chunkSize, IdLastLoaded: id}
	query := &app.Queries{Query: q, Chunk: a}
	Data, err := api.UseCase.FindAuthors(ctx, query)
	if err != nil {
		return errors.Wrap(err, "articlesHandler/GetByTag")
	}
	var ChunkData []amodels.Author
	for _, a := range Data.Author {
		val := *auConv(*a)
		ChunkData = append(ChunkData, val)
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
	fPath := "/api/v1/articles/query"
	Hits.WithLabelValues(layer, fPath).Inc()
	reqID := c.Request().Header.Get(echo.HeaderXRequestID)
	md := metadata.New(map[string]string{"X-Request-ID": reqID})
	ctx := metadata.NewOutgoingContext(context.Background(), md)
	a := &app.Chunk{ChunkSize: chunkSize, IdLastLoaded: id}
	query := &app.Queries{Query: q, Chunk: a}
	Data, err := api.UseCase.FindByTag(ctx, query)
	if err != nil {
		return errors.Wrap(err, "articlesHandler/GetByTag")
	}
	var ChunkData []amodels.Preview
	for _, a := range Data.Preview {
		val := reverseConv(a)
		ChunkData = append(ChunkData, *val)
	}
	response := amodels.ChunkResponse{
		Status:    http.StatusOK,
		ChunkData: ChunkData,
	}
	return c.JSON(http.StatusOK, response)
}

func (api *ArticlesHandler) Create(c echo.Context) error {
	tempArticle := new(amodels.ArticleCreate)
	fPath := "/api/v1/articles/create"
	Hits.WithLabelValues(layer, fPath).Inc()
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
	cook := cookie.Value
	reqID := c.Request().Header.Get(echo.HeaderXRequestID)
	md := metadata.New(map[string]string{"X-Request-ID": reqID})
	ctx := metadata.NewOutgoingContext(context.Background(), md)
	art := arConv(tempArticle)
	cr := &app.Create{Art: art, Value: cook}
	c.Logger().Error("!!!!!!", tempArticle.Category)
	Id, err := api.UseCase.Store(ctx, cr)
	if err != nil {
		return errors.Wrap(err, "articlesHandler/Create")
	}

	response := amodels.GenericResponse{
		Status: http.StatusOK,
		Data:   fmt.Sprint(Id.Id),
	}

	return c.JSON(http.StatusOK, response)
}

func (api *ArticlesHandler) Delete(c echo.Context) error {
	id := c.QueryParam("id")
	fPath := "/api/v1/articles/delete"
	Hits.WithLabelValues(layer, fPath).Inc()
	reqID := c.Request().Header.Get(echo.HeaderXRequestID)
	md := metadata.New(map[string]string{"X-Request-ID": reqID})
	ctx := metadata.NewOutgoingContext(context.Background(), md)
	d := &app.Id{Id: id}
	_, err := api.UseCase.Delete(ctx, d)
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
