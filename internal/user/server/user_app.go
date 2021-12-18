package main

import (
	"context"

	umodels "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/user/models"
	app "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/user/user_app"
	"github.com/pkg/errors"
)

type UserManager struct {
	usecase umodels.UserUsecase
}

func NewUserManager(handler umodels.UserUsecase) *UserManager {
	return &UserManager{
		usecase: handler,
	}
}

func convertToUserModel(user app.User) *umodels.User {
	return &umodels.User{
		Login:       user.Login,
		Name:        user.FirstName,
		Surname:     user.LastName,
		Email:       user.Email,
		Password:    user.Password,
		Score:       int(user.Score),
		AvatarURL:   user.AvatarUrl,
		Description: user.Description,
	}
}

func (m *UserManager) UpdateProfile(ctx context.Context, updateInput *app.UpdateInput) (*app.LoginResponse, error) {
	response, err := m.usecase.UpdateProfile(ctx, convertToUserModel(*updateInput.User), updateInput.SesionID)

	return &app.LoginResponse{
		Status: uint32(response.Status),
		Data: &app.LoginData{
			Login:       response.Data.Login,
			FirstName:   response.Data.Name,
			LastName:    response.Data.Surname,
			Email:       response.Data.Email,
			Score:       int32(response.Data.Score),
			AvatarUrl:   response.Data.AvatarURL,
			Description: response.Data.Description,
		},
		Msg: response.Msg,
	}, errors.Wrap(err, "user_app/UpdateProfile")
}

func (m *UserManager) GetAuthorProfile(ctx context.Context, author *app.Author) (*app.GetUserResponse, error) {
	response, err := m.usecase.GetAuthorProfile(ctx, author.Author)

	return &app.GetUserResponse{
		Status: uint32(response.Status),
		Data: &app.GetUserData{
			Login:       response.Data.Login,
			FirstName:   response.Data.Name,
			LastName:    response.Data.Surname,
			Score:       int32(response.Data.Score),
			AvatarUrl:   response.Data.AvatarURL,
			Description: response.Data.Description,
		},
		Msg: response.Msg,
	}, errors.Wrap(err, "user_app/GetAuthorProfile")
}

func (m *UserManager) GetUserProfile(ctx context.Context, sessionID *app.SessionID) (*app.GetUserResponse, error) {
	response, err := m.usecase.GetUserProfile(ctx, sessionID.SesionID)

	return &app.GetUserResponse{
		Status: uint32(response.Status),
		Data: &app.GetUserData{
			Login:       response.Data.Login,
			FirstName:   response.Data.Name,
			LastName:    response.Data.Surname,
			Score:       int32(response.Data.Score),
			AvatarUrl:   response.Data.AvatarURL,
			Description: response.Data.Description,
		},
		Msg: response.Msg,
	}, errors.Wrap(err, "user_app/GetUserProfile")
}

func (m *UserManager) LoginUser(ctx context.Context, user *app.User) (*app.LoginResponse, error) {
	response, sessionID, err := m.usecase.LoginUser(ctx, convertToUserModel(*user))

	return &app.LoginResponse{
		Status: uint32(response.Status),
		Data: &app.LoginData{
			Login:       response.Data.Login,
			FirstName:   response.Data.Name,
			LastName:    response.Data.Surname,
			Email:       response.Data.Email,
			Score:       int32(response.Data.Score),
			AvatarUrl:   response.Data.AvatarURL,
			Description: response.Data.Description,
		},
		Msg:       response.Msg,
		SessionID: sessionID,
	}, errors.Wrap(err, "user_app/LoginUser")
}

func (m *UserManager) Signup(ctx context.Context, user *app.User) (*app.SignupResponse, error) {
	response, sessionID, err := m.usecase.Signup(ctx, convertToUserModel(*user))

	return &app.SignupResponse{
		Status: uint32(response.Status),
		Data: &app.SignUpData{
			Login:       response.Data.Login,
			FirstName:   response.Data.Name,
			LastName:    response.Data.Surname,
			Email:       response.Data.Email,
			Score:       int32(response.Data.Score),
			AvatarUrl:   response.Data.AvatarURL,
			Description: response.Data.Description,
		},
		Msg:       response.Msg,
		SessionID: sessionID,
	}, err
}

func (m *UserManager) Logout(ctx context.Context, cookieValue *app.CookieValue) (*app.Nothing, error) {
	err := m.usecase.Logout(ctx, cookieValue.CookieValue)

	return &app.Nothing{}, errors.Wrap(err, "user_app/Logout")
}
