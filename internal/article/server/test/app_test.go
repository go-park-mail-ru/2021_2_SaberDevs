package testing

import (
	"context"
	"testing"

	app "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/article/article_app"
	"github.com/go-park-mail-ru/2021_2_SaberDevs/internal/article/models"
	mocks "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/article/models/mock"
	ser "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/article/server/serve"
	gomock "github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestGetByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockUcase := mocks.NewMockArticleUsecase(ctrl)
	mockArticle := models.FullArticle{
		Id:          "1",
		PreviewUrl:  "#",
		Tags:        []string{"fishing", "boat"},
		Title:       "fishing",
		Text:        "fishing is bad for business",
		Author:      models.Author{Id: 1},
		CommentsUrl: "#",
		Comments:    11,
		Likes:       12,
	}
	resArticle := &app.FullArticle{
		Id:          "1",
		PreviewUrl:  "#",
		Tags:        []string{"fishing", "boat"},
		Title:       "fishing",
		Text:        "fishing is bad for business",
		Author:      &app.Author{Id: 1},
		CommentsUrl: "#",
		Comments:    11,
		Likes:       12,
	}
	u := ser.NewArticleManager(mockUcase)
	login := "Iam"
	id := &app.Id{Id: "1", Value: login}
	t.Run("success", func(t *testing.T) {
		mockUcase.EXPECT().GetByID(gomock.Eq(context.TODO()), gomock.Eq(login), gomock.Eq(int64(1))).Return(mockArticle, nil).AnyTimes()
		a, err := u.GetByID(context.TODO(), id)
		assert.NoError(t, err)
		assert.Equal(t, a, resArticle)

	})
}

func TestFetch(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockUcase := mocks.NewMockArticleUsecase(ctrl)
	mockArticle := models.Preview{
		Id:          "1",
		PreviewUrl:  "#",
		Tags:        []string{"fishing", "boat"},
		Title:       "fishing",
		Text:        "fishing is bad for business",
		Author:      models.Author{Id: 1},
		CommentsUrl: "#",
		Comments:    11,
		Likes:       12,
	}
	preArticle := []models.Preview{mockArticle}
	resArticle := &app.Preview{
		Id:          "1",
		PreviewUrl:  "#",
		Tags:        []string{"fishing", "boat"},
		Title:       "fishing",
		Text:        "fishing is bad for business",
		Author:      &app.Author{Id: 1},
		CommentsUrl: "#",
		Comments:    11,
		Likes:       12,
	}
	repArticle := &app.Repview{Preview: []*app.Preview{resArticle}}
	login := "Iam"
	ch := &app.Chunk{IdLastLoaded: "999999", ChunkSize: 5, Value: login}
	u := ser.NewArticleManager(mockUcase)
	t.Run("success", func(t *testing.T) {
		mockUcase.EXPECT().Fetch(gomock.Eq(context.TODO()), gomock.Eq(login), "999999", 5).Return(preArticle, nil).AnyTimes()
		a, err := u.Fetch(context.TODO(), ch)
		assert.NoError(t, err)
		assert.Equal(t, a, repArticle)

	})
}
