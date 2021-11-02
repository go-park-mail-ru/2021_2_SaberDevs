package testing

import (
	"context"
	"testing"

	amodels "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/article/models"
	mocks "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/article/models/mock"
	repo "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/article/usecase"
	"github.com/go-park-mail-ru/2021_2_SaberDevs/internal/session/models"
	gomock "github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestGetByID(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	mockArticleRepo := mocks.NewMockArticleRepository()(ctrl)
	mockArticle := amodels.Article{
		Id:           "1",
		PreviewUrl:   "#",
		Tags:         []string{"fishing", "boat"},
		Title:        "fishing",
		Text:         "fishing is bad for business",
		AuthorUrl:    "$$$",
		AuthorName:   "Fisher",
		AuthorAvatar: "ElonMusk",
		CommentsUrl:  "#",
		Comments:     11,
		Likes:        12,
	}

	t.Run("success", func(t *testing.T) {
		mockArticleRepo.EXPECT().GetByID(gomock.Eq(context.TODO()), gomock.Eq(int64(1))).Return(mockArticle, nil).AnyTimes()

		u := repo.NewArticleUsecase(mockArticleRepo, models.SessionRepository)
		a, err := u.GetByID(context.TODO(), 1)
		assert.NoError(t, err)
		assert.NotNil(t, a)

	})

}
