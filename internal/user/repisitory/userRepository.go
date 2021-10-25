package repisitory

import (
	"context"
	"github.com/go-park-mail-ru/2021_2_SaberDevs/internal/data"
	sbErr "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/syberErrors"
	umodels "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/user/models"
	"sync"
)

type userMemoryRepo struct {
	users sync.Map
}

func NewUserRepository() umodels.UserRepository {
	var userRepo userMemoryRepo
	for _, user := range data.TestUsers {
		userRepo.users.Store(user.Email, user)
	}
	return &userRepo
}

// TODO error handling
func (r *userMemoryRepo) GetByEmail(ctx context.Context, email string) (umodels.User, error) {
	u, ok := r.users.Load(email)
	if !ok {
		return umodels.User{}, sbErr.ErrUserDoesntExist{
			Reason:   "no user in memory repo",
			Function: "userRepository/GetByEmail",
		}
	}
	return u.(umodels.User), nil
}

func (r *userMemoryRepo) Store(ctx context.Context, user *umodels.User) (umodels.User, error) {
	r.users.Store(user.Email, user)
	return *user, nil
}
