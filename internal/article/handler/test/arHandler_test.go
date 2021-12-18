package testing

import (
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
		ctx := e.AcquireContext()
		myid := apps.Id{Id: "1"}
		app.EXPECT().GetByID(ctx, &myid).Return(nil).AnyTimes()
		err := handler.GetByID(ctx)
		assert.NoError(t, err)
	})

}
