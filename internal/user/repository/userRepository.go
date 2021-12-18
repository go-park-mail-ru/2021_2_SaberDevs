package repository

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

func (r *userPsqlRepo) GetByName(ctx context.Context, name string) (umodels.User, error) {
	user := umodels.User{}

	err := r.Db.Get(&user, `SELECT Login, Name, Surname, Email, Password, Score, AvatarUrl, Description FROM author WHERE Name = $1`, name)
	if err != nil {
		return umodels.User{}, sbErr.ErrUserDoesntExist{
			Reason:   err.Error(),
			Function: "userRepositiry/GetByName",
		}
	}

	return user, nil
}

func (r *userPsqlRepo) UpdateUser(ctx context.Context, user *umodels.User) (umodels.User, error) {
	tx, err := r.Db.Beginx()
	if err != nil {
		return umodels.User{}, sbErr.ErrInternal{
			Reason:   err.Error(),
			Function: "userRepository/UpdateUser",
		}
	}

	if user.Description != "" {
		_, err := tx.Exec(`UPDATE author SET Description = $1 WHERE Login = $2`, user.Description, user.Login)
		if err != nil {
			err := tx.Rollback()
			if err != nil {
				return umodels.User{}, sbErr.ErrInternal{
					Reason:   err.Error(),
					Function: "userRepository/UpdateUser",
				}
			}
			return umodels.User{}, sbErr.ErrInternal{
				Reason:   err.Error(),
				Function: "userRepository/UpdateUser",
			}
		}
	}

	if user.Name != "" {
		_, err := tx.Exec(`UPDATE author SET NAME = $1 WHERE Login = $2`, user.Name, user.Login)
		if err != nil {
			err := tx.Rollback()
			if err != nil {
				return umodels.User{}, sbErr.ErrInternal{
					Reason:   err.Error(),
					Function: "userRepository/UpdateUser",
				}
			}
			return umodels.User{}, sbErr.ErrInternal{
				Reason:   err.Error(),
				Function: "userRepository/UpdateUser",
			}
		}
	}

	if user.Surname != "" {
		_, err := tx.Exec(`UPDATE author SET SURNAME = $1 WHERE Login = $2`, user.Surname, user.Login)
		if err != nil {
			err := tx.Rollback()
			if err != nil {
				return umodels.User{}, sbErr.ErrInternal{
					Reason:   err.Error(),
					Function: "userRepository/UpdateUser",
				}
			}
			return umodels.User{}, sbErr.ErrInternal{
				Reason:   err.Error(),
				Function: "userRepository/UpdateUser",
			}
		}
	}

	if user.Password != "" {
		_, err := tx.Exec(`UPDATE author SET PASSWORD = $1 WHERE Login = $2`, user.Password, user.Login)
		if err != nil {
			err := tx.Rollback()
			if err != nil {
				return umodels.User{}, sbErr.ErrInternal{
					Reason:   err.Error(),
					Function: "userRepository/UpdateUser",
				}
			}
			return umodels.User{}, sbErr.ErrInternal{
				Reason:   err.Error(),
				Function: "userRepository/UpdateUser",
			}
		}
	}

	if user.AvatarURL != "" {
		_, err := tx.Exec(`UPDATE author SET AvatarUrl = $1 WHERE Login = $2`, user.AvatarURL, user.Login)
		if err != nil {
			err := tx.Rollback()
			if err != nil {
				return umodels.User{}, sbErr.ErrInternal{
					Reason:   err.Error(),
					Function: "userRepository/UpdateUser",
				}
			}
			return umodels.User{}, sbErr.ErrInternal{
				Reason:   err.Error(),
				Function: "userRepository/UpdateUser",
			}
		}
	}

	err = tx.Commit()
	if err != nil {
		// добавить Rollback?
		return umodels.User{}, sbErr.ErrInternal{
			Reason:   err.Error(),
			Function: "userRepository/UpdateUser",
		}
	}

	updatedUser := umodels.User{
		Name:        user.Name,
		Surname:     user.Surname,
		Description: user.Description,
		AvatarURL:   user.AvatarURL,
	}

	return updatedUser, nil
}

func (r *userPsqlRepo) GetByLogin(ctx context.Context, login string) (umodels.User, error) {
	user := umodels.User{}

	err := r.Db.Get(&user, `SELECT Login, Name, Surname, Email, Password, Score, AvatarUrl, Description FROM author WHERE Login = $1`, login)
	if err != nil {
		return umodels.User{}, sbErr.ErrUserDoesntExist{
			Reason:   err.Error(),
			Function: "userRepositiry/GetByLogin",
		}
	}

	return user, nil
}

func (r *userPsqlRepo) Store(ctx context.Context, user *umodels.User) (umodels.User, error) {
	var login string

	err := r.Db.Get(&login, "SELECT login FROM author WHERE login = $1", user.Login)
	if login != "" {
		return umodels.User{}, sbErr.ErrUserExists{
			Reason:   "Логин уже занят",
			Function: "userRepository/Store",
		}
	}

	schema := `INSERT INTO author (Login, Name, Surname, Email, Password, Score, AvatarUrl, Description) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`
	_, err = r.Db.Exec(schema, user.Login, user.Name, user.Surname, user.Email, user.Password, 0, "", "")
	if err != nil {
		return umodels.User{}, sbErr.ErrInternal{
			Reason:   err.Error(),
			Function: "userRepository/Store",
		}
	}

	return *user, nil
}
