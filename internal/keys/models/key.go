package models

import (
	"context"
)

type Key struct {
	_msgpack struct{} `msgpack:",asArray"`
	Salt     string
	Login    string
}

type KeyRepository interface {
	StoreSalt(ctx context.Context, key Key) error
	DeleteSalt(ctx context.Context, login string) error
	GetSalt(ctx context.Context, login string) (string, error)
}
