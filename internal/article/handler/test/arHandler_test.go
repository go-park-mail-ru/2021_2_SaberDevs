package testing

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	apps "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/article/article_app"
	mock "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/article/article_app/mock"
	hand "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/article/handler"
	models "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/article/models"
	gomock "github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestGetByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	app := mock.NewMockArticleDeliveryClient(ctrl)

	handler := hand.NewArticlesHandler(app)

	t.Run("success", func(t *testing.T) {
		e := echo.New()

		req := httptest.NewRequest(echo.GET, "http://localhost:8081/api/v1/articles?Id=1", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		rec := httptest.NewRecorder()

		c := e.NewContext(req, rec)
		new := apps.FullArticle{}
		new.Author = &apps.Author{}
		app.EXPECT().GetByID(gomock.Any(), gomock.Any()).Return(&new, nil).AnyTimes()
		err := handler.GetByID(c)
		assert.NoError(t, err)
		assert.Equal(t, rec.Result().StatusCode, 200)
	})

}

func TestGetByFeed(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	app := mock.NewMockArticleDeliveryClient(ctrl)

	handler := hand.NewArticlesHandler(app)

	t.Run("success", func(t *testing.T) {
		e := echo.New()

		req := httptest.NewRequest(echo.GET, "http://localhost:8081/api/v1/articles/feed?idLastLoaded=9&login=all", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		rec := httptest.NewRecorder()

		c := e.NewContext(req, rec)
		new := apps.Repview{}
		new.Preview = []*apps.Preview{}

		app.EXPECT().Fetch(gomock.Any(), gomock.Any()).Return(&new, nil).AnyTimes()
		err := handler.GetFeed(c)
		assert.NoError(t, err)
		assert.Equal(t, rec.Result().StatusCode, 200)
	})

}

func TestGetByAuthor(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	app := mock.NewMockArticleDeliveryClient(ctrl)

	handler := hand.NewArticlesHandler(app)

	t.Run("success", func(t *testing.T) {
		e := echo.New()

		req := httptest.NewRequest(echo.GET, "http://localhost:8081/api/v1/articles/author?idLastLoaded=&login=DenisTest", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		rec := httptest.NewRecorder()

		c := e.NewContext(req, rec)
		new := apps.Repview{}
		new.Preview = []*apps.Preview{}

		app.EXPECT().GetByAuthor(gomock.Any(), gomock.Any()).Return(&new, nil).AnyTimes()
		err := handler.GetByAuthor(c)
		assert.NoError(t, err)
		assert.Equal(t, rec.Result().StatusCode, 200)
	})

}

func TestGetByCategory(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	app := mock.NewMockArticleDeliveryClient(ctrl)

	handler := hand.NewArticlesHandler(app)

	t.Run("success", func(t *testing.T) {
		e := echo.New()

		req := httptest.NewRequest(echo.GET, "http://localhost:8081/api/v1/articles/category?idLastLoaded=&category=Маркетинг", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		rec := httptest.NewRecorder()

		c := e.NewContext(req, rec)
		new := apps.Repview{}
		new.Preview = []*apps.Preview{}

		app.EXPECT().GetByCategory(gomock.Any(), gomock.Any()).Return(&new, nil).AnyTimes()
		err := handler.GetByCategory(c)
		assert.NoError(t, err)
		assert.Equal(t, rec.Result().StatusCode, 200)
	})

}

func TestGetByTag(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	app := mock.NewMockArticleDeliveryClient(ctrl)

	handler := hand.NewArticlesHandler(app)

	t.Run("success", func(t *testing.T) {
		e := echo.New()

		req := httptest.NewRequest(echo.GET, "http://localhost:8081/api/v1/articles/tags?idLastLoaded=&tag=abc", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		rec := httptest.NewRecorder()

		c := e.NewContext(req, rec)
		new := apps.Repview{}
		new.Preview = []*apps.Preview{}

		app.EXPECT().GetByTag(gomock.Any(), gomock.Any()).Return(&new, nil).AnyTimes()
		err := handler.GetByTag(c)
		assert.NoError(t, err)
		assert.Equal(t, rec.Result().StatusCode, 200)
	})

}

func TestFindArticles(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	app := mock.NewMockArticleDeliveryClient(ctrl)

	handler := hand.NewArticlesHandler(app)

	t.Run("success", func(t *testing.T) {
		e := echo.New()

		req := httptest.NewRequest(echo.GET, "http://localhost:8081/api/v1/search/articles?q=asdf&idLastLoaded=", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		rec := httptest.NewRecorder()

		c := e.NewContext(req, rec)
		new := apps.Repview{}
		new.Preview = []*apps.Preview{}

		app.EXPECT().FindArticles(gomock.Any(), gomock.Any()).Return(&new, nil).AnyTimes()
		err := handler.FindArticles(c)
		assert.NoError(t, err)
		assert.Equal(t, rec.Result().StatusCode, 200)
	})

}
func TestFindAuthors(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	app := mock.NewMockArticleDeliveryClient(ctrl)

	handler := hand.NewArticlesHandler(app)

	t.Run("success", func(t *testing.T) {
		e := echo.New()

		req := httptest.NewRequest(echo.GET, "http://localhost:8081/api/v1/search/authors?q=asdf&idLastLoaded=", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		rec := httptest.NewRecorder()

		c := e.NewContext(req, rec)
		new := apps.AView{}
		auth := &apps.Author{AvatarUrl: "who"}
		new.Author = []*apps.Author{auth}

		app.EXPECT().FindAuthors(gomock.Any(), gomock.Any()).Return(&new, nil).AnyTimes()
		err := handler.FindAuthors(c)
		assert.NoError(t, err)
		assert.Equal(t, rec.Result().StatusCode, 200)
	})

}
func TestFindbyTag(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	app := mock.NewMockArticleDeliveryClient(ctrl)

	handler := hand.NewArticlesHandler(app)

	t.Run("success", func(t *testing.T) {
		e := echo.New()

		req := httptest.NewRequest(echo.GET, "http://localhost:8081/api/v1/search/tags?q=asdf&idLastLoaded=", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		rec := httptest.NewRecorder()

		c := e.NewContext(req, rec)
		new := apps.Repview{}
		new.Preview = []*apps.Preview{}

		app.EXPECT().FindByTag(gomock.Any(), gomock.Any()).Return(&new, nil).AnyTimes()
		err := handler.FindByTag(c)
		assert.NoError(t, err)
		assert.Equal(t, rec.Result().StatusCode, 200)
	})

}

func TestUpdate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	app := mock.NewMockArticleDeliveryClient(ctrl)

	handler := hand.NewArticlesHandler(app)

	t.Run("success", func(t *testing.T) {
		e := echo.New()
		art := models.ArticleUpdate{}
		body, _ := json.Marshal(art)
		req := httptest.NewRequest(echo.POST, "http://localhost:8081/api/v1/articles/update", bytes.NewReader(body))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		cc := &http.Cookie{Name: "session", Value: "123"}
		req.AddCookie(cc)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		new := apps.Nothing{}
		app.EXPECT().Update(gomock.Any(), gomock.Any()).Return(&new, nil).AnyTimes()
		err := handler.Update(c)
		assert.NoError(t, err)
		assert.Equal(t, rec.Result().StatusCode, 200)
	})

}

func TestCreate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	app := mock.NewMockArticleDeliveryClient(ctrl)

	handler := hand.NewArticlesHandler(app)

	t.Run("success", func(t *testing.T) {
		e := echo.New()
		art := models.ArticleCreate{}
		body, _ := json.Marshal(art)
		req := httptest.NewRequest(echo.POST, "http://localhost:8081/api/v1/articles/create", bytes.NewReader(body))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		cc := &http.Cookie{Name: "session", Value: "123"}
		req.AddCookie(cc)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		new := apps.Created{Id: 1}
		app.EXPECT().Store(gomock.Any(), gomock.Any()).Return(&new, nil).AnyTimes()
		err := handler.Create(c)
		assert.NoError(t, err)
		assert.Equal(t, rec.Result().StatusCode, 200)
	})

}

func TestDelete(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	app := mock.NewMockArticleDeliveryClient(ctrl)

	handler := hand.NewArticlesHandler(app)

	t.Run("success", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(echo.POST, "http://localhost:8081/api/v1/articles/delete?Id=1", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		cc := &http.Cookie{Name: "session", Value: "123"}
		req.AddCookie(cc)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		new := apps.Nothing{}
		app.EXPECT().Delete(gomock.Any(), gomock.Any()).Return(&new, nil).AnyTimes()
		err := handler.Delete(c)
		assert.NoError(t, err)
		assert.Equal(t, rec.Result().StatusCode, 200)
	})

}
