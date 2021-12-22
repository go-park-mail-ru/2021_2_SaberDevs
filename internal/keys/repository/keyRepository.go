package repository

import (
	"context"

	wrapper "github.com/go-park-mail-ru/2021_2_SaberDevs/internal"
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

func (r *keyTarantoolRepo) StoreSalt(ctx context.Context, key kmodels.Key) error {
	path := "StoreSalt"
	_, err := wrapper.MyInsert(r.conn, path, "keys", []interface{}{key.Login, key.Salt})
	if err != nil {
		return sbErr.ErrInternal{
			Reason:   err.Error(),
			Function: "keyRepositiry/StoreSalt"}
	}

	return nil
}

func (r *keyTarantoolRepo) GetSalt(ctx context.Context, email string) (string, error) {
	path := "GetSalt"
	var key []kmodels.Key
	err := wrapper.MySelectTyped(r.conn, path, "keys", "primary", 0, 1, tarantool.IterEq, []interface{}{email}, &key)
	if err != nil {
		return "", sbErr.ErrNoSession{
			Reason:   err.Error(),
			Function: "keyRepositiry/GetSalt"}
	}

	return key[0].Salt, nil
}
