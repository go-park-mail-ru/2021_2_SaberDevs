package repository

import (
	"context"
	smodels "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/session/models"
	uuid "github.com/satori/go.uuid"
	"sync"
)

type sessionMemoryRepo struct {
	sessions    sync.Map
}

func NewSessionRepository() smodels.SessionRepository {
	return &sessionMemoryRepo{}
}

// TODO чья обязанность создавать значение куки?
func (r *sessionMemoryRepo) CreateSession(ctx context.Context, email string) (string, error) {
	cookieValue := uuid.NewV4().String()
	r.sessions.Store(cookieValue, email)
	return cookieValue, nil
}

func (r *sessionMemoryRepo) DeleteSession(ctx context.Context) error {
	return nil
}

func (r *sessionMemoryRepo) IsSession(ctx context.Context) error {
	return nil
}
