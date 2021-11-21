package usecase

import (
	"context"
	"net/http"

	smodels "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/session/models"
	umodels "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/user/models"
	"github.com/pkg/errors"
)

type sessionUsecase struct {
	userRepo    umodels.UserRepository
	sessionRepo smodels.SessionRepository
}

func NewsessionUsecase(ur umodels.UserRepository, sr smodels.SessionRepository) smodels.SessionUsecase {
	return &sessionUsecase{
		userRepo:    ur,
		sessionRepo: sr,
	}
}

func (su *sessionUsecase) IsSession(ctx context.Context, cookie string) (umodels.LoginResponse, error) {
	var response umodels.LoginResponse

	login, err := su.sessionRepo.GetSessionLogin(ctx, cookie)
	if err != nil {
		return response, errors.Wrap(err, "sessionUsecase/GetSessionLogin")
	}

	userInRepo, err := su.userRepo.GetByLogin(ctx, login)
	if err != nil {
		return response, errors.Wrap(err, "sessionUsecase/GetSessionLogin")
	}

	d := umodels.LoginData{
		Login:       userInRepo.Login,
		Name:        userInRepo.Name,
		Surname:     userInRepo.Surname,
		Email:       userInRepo.Email,
		Score:       userInRepo.Score,
		AvatarURL:   userInRepo.AvatarURL,
		Description: userInRepo.Description,
	}
	response = umodels.LoginResponse{
		Status: http.StatusOK,
		Data:   d,
		Msg:    "OK",
	}

	return response, nil
}
