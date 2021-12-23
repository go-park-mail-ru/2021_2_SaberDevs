package main

import (
	"context"
	"github.com/go-park-mail-ru/2021_2_SaberDevs/internal/user/models"
	app "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/user/user_app"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
	mock "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/user/models/mock"
)

func TestUpdateProfile(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	usecase := mock.NewMockUserUsecase(ctrl)
	m := NewUserManager(usecase)

	t.Run("success", func(t *testing.T) {
		resp := models.LoginResponse{}
		usecase.EXPECT().UpdateProfile(gomock.Any(), gomock.Any(), gomock.Any()).Return(resp, nil).AnyTimes()


		_, err := m.UpdateProfile(context.TODO(), &app.UpdateInput{User: &app.User{Score: 5}})
		assert.NoError(t, err)
	})
}

func TestGetAuthorProfile(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	usecase := mock.NewMockUserUsecase(ctrl)
	m := NewUserManager(usecase)

	t.Run("success", func(t *testing.T) {
		resp := models.GetUserResponse{}
		usecase.EXPECT().GetAuthorProfile(gomock.Any(), gomock.Any()).Return(resp, nil).AnyTimes()


		_, err := m.GetAuthorProfile(context.TODO(), &app.Author{})
		assert.NoError(t, err)
	})
}

func TestGetGetUserProfile(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	usecase := mock.NewMockUserUsecase(ctrl)
	m := NewUserManager(usecase)

	t.Run("success", func(t *testing.T) {
		resp := models.GetUserResponse{}
		usecase.EXPECT().GetUserProfile(gomock.Any(), gomock.Any()).Return(resp, nil).AnyTimes()


		_, err := m.GetUserProfile(context.TODO(), &app.SessionID{})
		assert.NoError(t, err)
	})
}

func TestLoginUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	usecase := mock.NewMockUserUsecase(ctrl)
	m := NewUserManager(usecase)

	t.Run("success", func(t *testing.T) {
		resp := models.LoginResponse{}
		usecase.EXPECT().LoginUser(gomock.Any(), gomock.Any()).Return(resp, "", nil).AnyTimes()


		_, err := m.LoginUser(context.TODO(), &app.User{Score: 9})
		assert.NoError(t, err)
	})
}

func TestSignup(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	usecase := mock.NewMockUserUsecase(ctrl)
	m := NewUserManager(usecase)

	t.Run("success", func(t *testing.T) {
		resp := models.SignupResponse{}
		usecase.EXPECT().Signup(gomock.Any(), gomock.Any()).Return(resp, "", nil).AnyTimes()

		_, err := m.Signup(context.TODO(), &app.User{Score: 9})
		assert.NoError(t, err)
	})
}

func TestLogout(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	usecase := mock.NewMockUserUsecase(ctrl)
	m := NewUserManager(usecase)

	t.Run("success", func(t *testing.T) {
		usecase.EXPECT().Logout(gomock.Any(), gomock.Any()).Return( nil).AnyTimes()

		_, err := m.Logout(context.TODO(), &app.CookieValue{})
		assert.NoError(t, err)
	})
}
