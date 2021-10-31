package models

import (
	"context"
)

type Key struct {
	_msgpack struct{} `msgpack:",asArray"`
	Salt     string
	Email    string
}

type KeyRepository interface {
	StoreSalt(ctx context.Context, key Key) error
	GetSalt(ctx context.Context, email string) (string, error)
}
