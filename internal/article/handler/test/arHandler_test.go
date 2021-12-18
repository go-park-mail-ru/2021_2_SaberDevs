package testing

import (
	"net/http/httptest"
	"testing"

	apps "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/article/article_app"
	mock "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/article/article_app/mock"
	hand "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/article/handler"
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
	})

}
