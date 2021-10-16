package models

import "context"

type Session struct {
	CookieValue string
	UserEmail string
}

type SessionRepository interface {
	CreateSession(ctx context.Context, email string, cookieValue string) error
	DeleteSession(ctx context.Context) error
	IsSession(ctx context.Context) error
}
