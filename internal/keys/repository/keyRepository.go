package repository

import (
	"context"

	kmodels "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/keys/models"

	sbErr "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/syberErrors"
	"github.com/tarantool/go-tarantool"
)

type keyTarantoolRepo struct {
	conn *tarantool.Connection
}

func NewKeyRepository(conn *tarantool.Connection) kmodels.KeyRepository {
	return &keyTarantoolRepo{conn: conn}
}

func (r *keyTarantoolRepo) DeleteSalt(ctx context.Context, login string) error {
	_, err := r.conn.Delete("keys", "primary", []interface{}{login})
	if err != nil {
		return sbErr.ErrInternal{
			Reason:   err.Error(),
			Function: "keyRepositiry/DeleteSalt"}
	}

	return nil
}

func (r *keyTarantoolRepo) StoreSalt(ctx context.Context, key kmodels.Key) error {
	_, err := r.conn.Insert("keys", []interface{}{key.Login, key.Salt})
	if err != nil {
		return sbErr.ErrInternal{
			Reason:   err.Error(),
			Function: "keyRepositiry/StoreSalt"}
	}

	return nil
}

func (r *keyTarantoolRepo) GetSalt(ctx context.Context, email string) (string, error) {
	var key []kmodels.Key

	err := r.conn.SelectTyped("keys", "primary", 0, 1, tarantool.IterEq, []interface{}{email}, &key)
	if err != nil {
		return "", sbErr.ErrNoSession{
			Reason:   err.Error(),
			Function: "keyRepositiry/GetSalt"}
	}
	if len(key) == 0 {
		return "", sbErr.ErrNoSession{
			Reason:   "no key",
			Function: "keyRepositiry/GetSalt"}
	}

	return key[0].Salt, nil
}
