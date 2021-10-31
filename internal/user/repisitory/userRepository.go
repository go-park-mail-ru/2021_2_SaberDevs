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
func (r *userPsqlRepo) GetByLogin(ctx context.Context, login string) (umodels.User, error) {
	user := umodels.User{}

	err := r.Db.Get(&user, "SELECT Login, Name, Surname, Email, Password, Score FROM author WHERE Login = $1", login)
	if err != nil {
		return user, sbErr.ErrUserDoesntExist{
			Reason:   err.Error(),
			Function: "userRepositiry/GetByEmail",
		}
	}

	return user, nil
}

func (r *userPsqlRepo) Store(ctx context.Context, user *umodels.User) (umodels.User, error) {
	var login string

	err := r.Db.Get(&login, "SELECT Email FROM author WHERE Email = $1", user.Login)
	if login != "" {
		return *user, sbErr.ErrUserExists{
			Reason:   "login already in use",
			Function: "userRepository/Store",
		}
	}

	schema := `INSERT INTO author (Login, Name, Surname, Email, Password, Score) VALUES ($1, $2, $3, $4, $5, $6)`
	_, err = r.Db.Exec(schema, user.Login, user.Name, user.Surname, user.Email, user.Password, 0)
	if err != nil {
		return *user, sbErr.ErrInternal{
			Reason:   err.Error(),
			Function: "userRepository/Store",
		}
	}

	return *user, nil
}
