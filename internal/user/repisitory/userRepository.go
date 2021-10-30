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

	err := r.Db.Get(&user, "SELECT Name, Surname, Email, Password, Score FROM author WHERE Email = $1", email)
	if err != nil {
		return user, sbErr.ErrUserDoesntExist{
			Reason: err.Error(),
			Function: "userRepositiry/GetByEmail",
		}
	}

	return user, nil
}

func (r *userPsqlRepo) Store(ctx context.Context, user *umodels.User) (umodels.User, error) {
	var email string

	err := r.Db.Get(&email, "SELECT Email FROM author WHERE Email = $1", user.Email)
	if err != nil {
		return *user, sbErr.ErrInternal{
			Reason:  err.Error(),
			Function: "userRepository/Store",
		}
	}
	if email != "" {
		return *user, sbErr.ErrUserExists{
			Reason:  "email already in use",
			Function: "userRepository/Store",
		}
	}

	schema := `INSERT INTO authors (Name, Surname, Email, Password, Score) VALUES ($1, $2, $3, $4, $5)`
	_, err = r.Db.Exec(schema, user.Name, user.Surname, user.Email, user.Password, 0)
	if err != nil {
		return *user, sbErr.ErrInternal{
			Reason:  err.Error(),
			Function: "userRepository/Store",
		}
	}
	// r.users.Store(user.Email, user)
	return *user, nil
}
