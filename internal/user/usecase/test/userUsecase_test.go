package test

import (
	"context"
	amocks "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/article/models/mock"
	kmodels "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/keys/models/mock"
	smocks "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/session/models/mock"
	umodels "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/user/models"
	umocks "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/user/models/mock"
	usecase "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/user/usecase"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetAuthorProfile(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	sesRepo := smocks.NewMockSessionRepository(ctrl)
	aRepo := amocks.NewMockArticleRepository(ctrl)
	uRepo := umocks.NewMockUserRepository(ctrl)
	kRepo := kmodels.NewMockKeyRepository(ctrl)

	t.Run("succes", func(t *testing.T) {
		user := umodels.User{}
		uRepo.EXPECT().GetByLogin(gomock.Any(), gomock.Any()).Return(user, nil).AnyTimes()

		u := usecase.NewUserUsecase(uRepo, sesRepo, kRepo, aRepo)

		_, err := u.GetAuthorProfile(context.TODO(), "a")
		assert.NoError(t, err)
	})
}

func TestGetUserProfile(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	sesRepo := smocks.NewMockSessionRepository(ctrl)
	aRepo := amocks.NewMockArticleRepository(ctrl)
	uRepo := umocks.NewMockUserRepository(ctrl)
	kRepo := kmodels.NewMockKeyRepository(ctrl)

	t.Run("succes", func(t *testing.T) {
		sesRepo.EXPECT().GetSessionLogin(gomock.Any(), gomock.Any()).Return("a", nil).AnyTimes()

		user := umodels.User{}
		uRepo.EXPECT().GetByLogin(gomock.Any(), gomock.Any()).Return(user, nil).AnyTimes()

		u := usecase.NewUserUsecase(uRepo, sesRepo, kRepo, aRepo)

		_, err := u.GetUserProfile(context.TODO(), "a")
		assert.NoError(t, err)
	})
}

func TestUpdateProfile(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	sesRepo := smocks.NewMockSessionRepository(ctrl)
	aRepo := amocks.NewMockArticleRepository(ctrl)
	uRepo := umocks.NewMockUserRepository(ctrl)
	kRepo := kmodels.NewMockKeyRepository(ctrl)

	t.Run("succes", func(t *testing.T) {
		sesRepo.EXPECT().GetSessionLogin(gomock.Any(), gomock.Any()).Return("a", nil).AnyTimes()

		uRepo.EXPECT().UpdateUser(gomock.Any(), gomock.Any()).Return(umodels.User{}, nil).AnyTimes()

		user := umodels.User{}
		uRepo.EXPECT().GetByLogin(gomock.Any(), gomock.Any()).Return(user, nil).AnyTimes()

		u := usecase.NewUserUsecase(uRepo, sesRepo, kRepo, aRepo)

		_, err := u.UpdateProfile(context.TODO(), &umodels.User{}, "a")
		assert.NoError(t, err)
	})
}

func TestLoginUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	sesRepo := smocks.NewMockSessionRepository(ctrl)
	aRepo := amocks.NewMockArticleRepository(ctrl)
	uRepo := umocks.NewMockUserRepository(ctrl)
	kRepo := kmodels.NewMockKeyRepository(ctrl)

	t.Run("succes", func(t *testing.T) {
		user := umodels.User{Password: "mollen"}
		uRepo.EXPECT().GetByLogin(gomock.Any(), gomock.Any()).Return(user, nil).AnyTimes()

		sesRepo.EXPECT().CreateSession(gomock.Any(), gomock.Any()).Return(",", nil).AnyTimes()

		u := usecase.NewUserUsecase(uRepo, sesRepo, kRepo, aRepo)

		_, _, err := u.LoginUser(context.TODO(), &umodels.User{Password: "mollen"})

		assert.NoError(t, err)
	})
}

func TestSignup(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	sesRepo := smocks.NewMockSessionRepository(ctrl)
	aRepo := amocks.NewMockArticleRepository(ctrl)
	uRepo := umocks.NewMockUserRepository(ctrl)
	kRepo := kmodels.NewMockKeyRepository(ctrl)

	t.Run("succes", func(t *testing.T) {
		user := umodels.User{Password: "mollen"}
		uRepo.EXPECT().Store(gomock.Any(), gomock.Any()).Return(user, nil).AnyTimes()

		sesRepo.EXPECT().CreateSession(gomock.Any(), gomock.Any()).Return(",", nil).AnyTimes()

		u := usecase.NewUserUsecase(uRepo, sesRepo, kRepo, aRepo)

		_, _, err := u.Signup(context.TODO(), &umodels.User{Password: "mollen"})

		assert.NoError(t, err)
	})
}

func TestLogout(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	sesRepo := smocks.NewMockSessionRepository(ctrl)
	aRepo := amocks.NewMockArticleRepository(ctrl)
	uRepo := umocks.NewMockUserRepository(ctrl)
	kRepo := kmodels.NewMockKeyRepository(ctrl)

	t.Run("succes", func(t *testing.T) {
		sesRepo.EXPECT().DeleteSession(gomock.Any(), gomock.Any()).Return( nil).AnyTimes()

		u := usecase.NewUserUsecase(uRepo, sesRepo, kRepo, aRepo)

		err := u.Logout(context.TODO()," &umodels.User{Password:}")

		assert.NoError(t, err)
	})
}
