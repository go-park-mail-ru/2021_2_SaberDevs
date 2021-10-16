package usecases

import (
	"context"
	"errors"
	smodels "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/session/models"
	umodels "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/user/models"
	"net/http"
)

type userUsecase struct {
	userRepo umodels.UserRepository
	sessionRepo smodels.SessionRepository
}

func NewUserUsecase(ur umodels.UserRepository, sr smodels.SessionRepository) umodels.UserUsecase {
	return &userUsecase{
		userRepo: ur,
		sessionRepo: sr,
	}
}

// TODO error handling
func (uu *userUsecase) LoginUser(ctx context.Context, user *umodels.User) (umodels.LoginResponse, string, error) {
	var response umodels.LoginResponse
	userInRepo, err := uu.userRepo.GetByEmail(ctx, user.Email)
	if err != nil {
		// TODO user doesnt exist err
		return response, "", err
	}

	if userInRepo.Password != user.Password {
		var err = errors.New("wrong password")
		return response, "", err
	}

	cookieValue, err := uu.sessionRepo.CreateSession(ctx, user.Email)

	d := umodels.LoginData{
		Login:   userInRepo.Login,
		Name:    userInRepo.Name,
		Surname: userInRepo.Surname,
		Email:   userInRepo.Email,
		Score:   userInRepo.Score,
	}
	response = umodels.LoginResponse{
		Status: http.StatusOK,
		Data:   d,
		Msg:    "OK",
	}

	return response, cookieValue, nil
}

func (uu *userUsecase) Signup(ctx context.Context, user *umodels.User) (umodels.SignupResponse, error) {
	var response umodels.SignupResponse
	return response, nil
}

func (uu *userUsecase) Logout(ctx context.Context, cookieValue string) error {
	err := uu.sessionRepo.DeleteSession(ctx, cookieValue)
	return err
}
