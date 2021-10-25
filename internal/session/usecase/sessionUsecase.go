package usecase

import (
	"context"

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

func (su *sessionUsecase) IsSession(ctx context.Context, cookie string) (umodels.User, error) {
	var user umodels.User

	email, err := su.sessionRepo.IsSession(ctx, cookie)
	if err != nil {
		return user, errors.Wrap(err, "sessionUsecase/IsSession")
	}

	user, err = su.userRepo.GetByEmail(ctx, email)
	if err != nil {
		return user, errors.Wrap(err, "sessionUsecase/IsSession")
	}

	return user, nil
}
