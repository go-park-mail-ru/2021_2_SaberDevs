package models

import (
	"context"
)

//easyjson:json
type Key struct {
	_msgpack struct{} `msgpack:",asArray"`
	Salt     string
	Login    string
}

type KeyRepository interface {
	StoreSalt(ctx context.Context, key Key) error
	GetSalt(ctx context.Context, login string) (string, error)
}
