package repository

import (
	"context"
	"google.golang.org/grpc/status"

	wrapper "github.com/go-park-mail-ru/2021_2_SaberDevs/internal"
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
	path := "GetByName"
	user := umodels.User{}
	schema := `SELECT Login, Name, Surname, Email, Password, Score, AvatarUrl, Description FROM author WHERE Name = $1`
	err := wrapper.MyGet(r.Db, path, schema, &user, name)
	if err != nil {
		return umodels.User{}, sbErr.ErrUserDoesntExist{
			Reason:   err.Error(),
			Function: "userRepositiry/GetByName",
		}
	}

	return user, nil
}

func (r *userPsqlRepo) UpdateUser(ctx context.Context, user *umodels.User) (umodels.User, error) {
	path := "UpdateUser"
	tx, err := wrapper.MyBegin(r.Db, path)
	if err != nil {
		return umodels.User{}, sbErr.ErrInternal{
			Reason:   err.Error(),
			Function: "userRepository/UpdateUser",
		}
	}

	if user.Description != "" {
		schema := `UPDATE author SET Description = $1 WHERE Login = $2`
		_, err := wrapper.MyTxExec(tx, path, schema, user.Description, user.Login)
		if err != nil {
			err := wrapper.MyRollBack(tx, path)
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
			err := wrapper.MyRollBack(tx, path)
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
			err := wrapper.MyRollBack(tx, path)
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
		schema := `UPDATE author SET PASSWORD = $1 WHERE Login = $2`
		_, err := wrapper.MyTxExec(tx, path, schema, user.Password, user.Login)
		if err != nil {
			err := wrapper.MyRollBack(tx, path)
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
		schema := `UPDATE author SET AvatarUrl = $1 WHERE Login = $2`
		_, err := wrapper.MyTxExec(tx, path, schema, user.AvatarURL, user.Login)
		if err != nil {
			err := wrapper.MyRollBack(tx, path)
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

	err = wrapper.MyCommit(tx, path)
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
	path := "GetByLogin"
	user := umodels.User{}
	schema := `SELECT Login, Name, Surname, Email, Password, Score, AvatarUrl, Description FROM author WHERE Login = $1`
	err := wrapper.MyGet(r.Db, path, schema, &user, login)
	if err != nil {
		return umodels.User{}, sbErr.ErrUserDoesntExist{
			Reason:   err.Error(),
			Function: "userRepositiry/GetByLogin",
		}
	}

	return user, nil
}

func (r *userPsqlRepo) Store(ctx context.Context, user *umodels.User) (umodels.User, error) {
	path := "Store"
	var login string
	schema := "SELECT login FROM author WHERE login = $1"
	err := wrapper.MyGet(r.Db, path, schema, &login, user.Login)
	if login != "" {
		// return umodels.User{}, sbErr.ErrUserExists{
		// 	Reason:   "Логин уже занят",
		// 	Function: "userRepository/Store",
		// }
		return umodels.User{}, status.Error(17, "Логин уже занят")
	}
	schema = `INSERT INTO author (Login, Name, Surname, Email, Password, Score, AvatarUrl, Description) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`
	_, err = wrapper.MyExec(r.Db, path, schema, user.Login, user.Name, user.Surname, user.Email, user.Password, 0, "", "")
	if err != nil {
		return umodels.User{}, sbErr.ErrInternal{
			Reason:   err.Error(),
			Function: "userRepository/Store",
		}
	}

	return *user, nil
}
