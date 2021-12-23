package test

import (
	"context"
	smocks "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/session/models/mock"
	uc "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/session/usecase"
	umodels "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/user/models"
	umocks "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/user/models/mock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateComment(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	sesRepo := smocks.NewMockSessionRepository(ctrl)
	uRepo := umocks.NewMockUserRepository(ctrl)

	usecase := uc.NewsessionUsecase(uRepo, sesRepo)

	t.Run("succes", func(t *testing.T) {
		sesRepo.EXPECT().GetSessionLogin(gomock.Any(), gomock.Any()).Return("a", nil).AnyTimes()

		user := umodels.User{}
		uRepo.EXPECT().GetByLogin(gomock.Any(), gomock.Any()).Return(user, nil).AnyTimes()

		_, err := usecase.IsSession(context.TODO(), "a")

		assert.NoError(t, err)
	})
}
