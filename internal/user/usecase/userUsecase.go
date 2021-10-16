package usecases

import (
	"context"
	"errors"
	umodels "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/user/models"
	"net/http"
)

type userUsecase struct {
	userRepo umodels.UserRepository
}

func NewUserUsecase(ur umodels.UserRepository) umodels.UserUsecase {
	return &userUsecase{
		userRepo: ur,
	}
}

// TODO error handling
func (ur *userUsecase) LoginUser(ctx context.Context, user *umodels.User) (umodels.LoginResponse, error) {
	var response umodels.LoginResponse
	userInRepo, err := ur.userRepo.GetByEmail(ctx, user.Email)
	if err != nil {
		// TODO user doesnt exist err
		return response, err
	}

	if userInRepo.Password != user.Password {
		var err = errors.New("wrong password")
		return response, err
	}

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

	return response, nil
}

func (ur *userUsecase) Signup(ctx context.Context, user *umodels.User) (umodels.SignupResponse, error) {
	var response umodels.SignupResponse
	return response, nil
}
