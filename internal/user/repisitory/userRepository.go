package repisitory

import (
	"context"

	sbErr "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/syberErrors"
	umodels "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/user/models"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type userPsqlRepo struct {
	Db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) umodels.UserRepository {
	return &userPsqlRepo{db}
}

// TODO error handling
func (r *userPsqlRepo) GetByEmail(ctx context.Context, email string) (umodels.User, error) {
	user := umodels.User{}
	err := r.Db.Get(&user, "SELECT Id, Name, Surname, Email, Password, Score FROM author WHERE Email = $1", email)
	if err != nil {
		return user, sbErr.ErrUserDoesntExist{
			Reason: err.Error(),
			Function: "userRepositiry/GetByEmail",
		}
	}
	// u, ok := r.users.Load(email)
	// if !ok {
	// 	return umodels.User{}, sbErr.ErrUserDoesntExist{
	// 		Reason:   "no user in memory repo",
	// 		Function: "userRepository/GetByEmail",
	// 	}
	// }
	return user, nil
}

func (r *userPsqlRepo) Store(ctx context.Context, user *umodels.User) (umodels.User, error) {
	// r.users.Store(user.Email, user)
	return *user, nil
}
