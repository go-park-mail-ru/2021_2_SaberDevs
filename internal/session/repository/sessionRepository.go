package repository

import (
	"context"
	smodels "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/session/models"
	"sync"
)

type sessionMemoryRepo struct {
	sessions    sync.Map
}

func NewSessionRepository() smodels.SessionRepository {
	return &sessionMemoryRepo{}
}

func (r *sessionMemoryRepo) CreateSession(ctx context.Context, email string, cookieValue string) error {

}

func (r *sessionMemoryRepo) DeleteSession(ctx context.Context) error {

}

func (r *sessionMemoryRepo) IsSession(ctx context.Context) error {

}
