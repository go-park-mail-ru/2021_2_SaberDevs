package usecases

import (
	"context"
	umodels "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/user/models"
)

type userUsecase struct {
	userRepo umodels.UserRepository
}

func NewUserUsecase(ur umodels.UserRepository) umodels.UserUsecase {
	return &userUsecase{
		userRepo: ur,
	}
}

func (u *userUsecase) GetByLogin(ctx context.Context, login string) error {
	return nil
}

func (u *userUsecase) Store(ctx context.Context, user umodels.User) error {
	return nil
}
