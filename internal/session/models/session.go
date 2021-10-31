package models

import (
	"context"
	umodels "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/user/models"
)

type Session struct {
	CookieValue string
	UserEmail string
}

type SessionUsecase interface {
	IsSession(ctx context.Context, cookie string) (umodels.User ,error)
}

type SessionRepository interface {
	CreateSession(ctx context.Context, email string) (string, error)
	DeleteSession(ctx context.Context, cookieValue string) error
	IsSession(ctx context.Context, cookie string) (string, error)
}
