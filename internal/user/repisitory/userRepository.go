package repisitory

import (
	"context"
	"errors"
	"github.com/go-park-mail-ru/2021_2_SaberDevs/internal/data"
	umodels "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/user/models"
	"sync"
)

type userMemoryRepo struct {
	users    sync.Map
}

func NewUserRepository() umodels.UserRepository {
	var userRepo userMemoryRepo
	for _, user := range data.TestUsers {
		userRepo.users.Store(user.Login, user)
	}
	return &userRepo
}

// TODO error handling
func (r *userMemoryRepo) GetByEmail(ctx context.Context, email string) (umodels.User, error) {
	u, ok := r.users.Load(email)
	if !ok {
		var err = errors.New("wrong password")
		return u.(umodels.User), err
	}
	return u.(umodels.User), nil
}

func (r *userMemoryRepo) Store(ctx context.Context, user *umodels.User) error {
	return nil
}
