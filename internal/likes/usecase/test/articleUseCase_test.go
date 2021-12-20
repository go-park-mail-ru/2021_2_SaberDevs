package testing

import (
	"context"
	"testing"

	amodels "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/likes/models"
	mocks "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/likes/models/mock"
	repo "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/likes/usecase"
	smocks "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/session/models/mock"
	gomock "github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestRating(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()
	mockSesRepo := smocks.NewMockSessionRepository(ctrl)
	mockLikesRepo := mocks.NewMockLikesRepository(ctrl)
	login := "Iam"

	t.Run("success", func(t *testing.T) {
		mocklike := amodels.LikeData{
			Ltype: 0,
			Sign:  1,
			Id:    1,
		}
		mockSesRepo.EXPECT().GetSessionLogin(gomock.Eq(context.TODO()), gomock.Eq(login)).Return(login, nil).AnyTimes()
		mockLikesRepo.EXPECT().Like(gomock.Any(), gomock.Any()).Return(1, nil).AnyTimes()
		u := repo.NewArLikeUsecase(mockLikesRepo, mockSesRepo)
		a, err := u.Rating(context.TODO(), &mocklike, login)
		assert.NoError(t, err)
		assert.Equal(t, a, 1)
	})
	t.Run("success", func(t *testing.T) {
		mocklike := amodels.LikeData{
			Ltype: 0,
			Sign:  -1,
			Id:    1,
		}
		mockSesRepo.EXPECT().GetSessionLogin(gomock.Eq(context.TODO()), gomock.Eq(login)).Return(login, nil).AnyTimes()
		mockLikesRepo.EXPECT().Dislike(gomock.Any(), gomock.Any()).Return(-1, nil).AnyTimes()
		u := repo.NewArLikeUsecase(mockLikesRepo, mockSesRepo)
		a, err := u.Rating(context.TODO(), &mocklike, login)
		assert.NoError(t, err)
		assert.Equal(t, a, -1)
	})

	t.Run("success", func(t *testing.T) {
		mocklike := amodels.LikeData{
			Ltype: 0,
			Sign:  0,
			Id:    1,
		}
		mockSesRepo.EXPECT().GetSessionLogin(gomock.Eq(context.TODO()), gomock.Eq(login)).Return(login, nil).AnyTimes()
		mockLikesRepo.EXPECT().Cancel(gomock.Any(), gomock.Any()).Return(0, nil).AnyTimes()
		u := repo.NewArLikeUsecase(mockLikesRepo, mockSesRepo)
		a, err := u.Rating(context.TODO(), &mocklike, login)
		assert.NoError(t, err)
		assert.Equal(t, a, 0)
	})

	t.Run("success", func(t *testing.T) {
		mocklike := amodels.LikeData{
			Ltype: 1,
			Sign:  1,
			Id:    1,
		}
		mockSesRepo.EXPECT().GetSessionLogin(gomock.Eq(context.TODO()), gomock.Eq(login)).Return(login, nil).AnyTimes()
		mockLikesRepo.EXPECT().Like(gomock.Any(), gomock.Any()).Return(1, nil).AnyTimes()
		u := repo.NewComLikeUsecase(mockLikesRepo, mockSesRepo)
		a, err := u.Rating(context.TODO(), &mocklike, login)
		assert.NoError(t, err)
		assert.Equal(t, a, 1)
	})
	t.Run("success", func(t *testing.T) {
		mocklike := amodels.LikeData{
			Ltype: 1,
			Sign:  -1,
			Id:    1,
		}
		mockSesRepo.EXPECT().GetSessionLogin(gomock.Eq(context.TODO()), gomock.Eq(login)).Return(login, nil).AnyTimes()
		mockLikesRepo.EXPECT().Dislike(gomock.Any(), gomock.Any()).Return(-1, nil).AnyTimes()
		u := repo.NewComLikeUsecase(mockLikesRepo, mockSesRepo)
		a, err := u.Rating(context.TODO(), &mocklike, login)
		assert.NoError(t, err)
		assert.Equal(t, a, -1)
	})
	t.Run("success", func(t *testing.T) {
		mocklike := amodels.LikeData{
			Ltype: 1,
			Sign:  0,
			Id:    1,
		}
		mockSesRepo.EXPECT().GetSessionLogin(gomock.Eq(context.TODO()), gomock.Eq(login)).Return(login, nil).AnyTimes()
		mockLikesRepo.EXPECT().Cancel(gomock.Any(), gomock.Any()).Return(0, nil).AnyTimes()
		u := repo.NewComLikeUsecase(mockLikesRepo, mockSesRepo)
		a, err := u.Rating(context.TODO(), &mocklike, login)
		assert.NoError(t, err)
		assert.Equal(t, a, 0)
	})

}
