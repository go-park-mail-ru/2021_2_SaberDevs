package models

import "context"

type Session struct {
	CookieValue string
	UserEmail string
}

type SessionRepository interface {
	CreateSession(ctx context.Context, email string) (string, error)
	DeleteSession(ctx context.Context, cookieValue string) error
	IsSession(ctx context.Context) error
}
