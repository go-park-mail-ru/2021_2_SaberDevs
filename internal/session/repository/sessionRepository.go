package repository

import (
	"context"
	smodels "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/session/models"
	uuid "github.com/satori/go.uuid"
	"sync"
)

type sessionMemoryRepo struct {
	sessions sync.Map
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

func (r *sessionMemoryRepo) DeleteSession(ctx context.Context, cookieValue string) error {
	r.sessions.Delete(cookieValue)
	return nil
}

func (r *sessionMemoryRepo) IsSession(ctx context.Context, cookie string) (string, error) {
	email, ok := r.sessions.Load(cookie)
	if !ok {
		// TODO return good error
		return "", nil
	}
	return email.(string), nil
}
