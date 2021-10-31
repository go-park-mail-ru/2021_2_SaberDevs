package repository

import (
	"context"

	sbErr "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/syberErrors"
	"github.com/tarantool/go-tarantool"

	smodels "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/session/models"
	uuid "github.com/satori/go.uuid"
)

type sessionTarantoolRepo struct {
	conn *tarantool.Connection
}

func NewSessionRepository(conn *tarantool.Connection) smodels.SessionRepository {
	return &sessionTarantoolRepo{conn: conn}
}

func (r *sessionTarantoolRepo) CreateSession(ctx context.Context, login string) (string, error) {
	sessionID := uuid.NewV4().String()

	_, err := r.conn.Insert("sessions", []interface{}{sessionID, login})
	if err != nil {
		return "", sbErr.ErrInternal{
			Reason:   err.Error(),
			Function: "sessionRepositiry/CreateSession"}
	}

	return sessionID, nil
}

func (r *sessionTarantoolRepo) DeleteSession(ctx context.Context, sessionID string) error {
	_, err := r.conn.Delete("sessions", "primary", []interface{}{sessionID})
	if err != nil {
		return sbErr.ErrInternal{
			Reason:   err.Error(),
			Function: "sessionRepositiry/IsSession"}
	}

	return nil
}

func (r *sessionTarantoolRepo) IsSession(ctx context.Context, sessionID string) (string, error) {
	var user []smodels.Session

	err := r.conn.SelectTyped("sessions", "primary", 0, 1, tarantool.IterEq, []interface{}{sessionID}, &user)
	if err != nil {
		return "", sbErr.ErrInternal{
			Reason:   err.Error(),
			Function: "sessionRepositiry/IsSession"}
	}
	if len(user) == 0 {
		return "", sbErr.ErrNoSession{
			Reason:   "no session",
			Function: "sessionRepositiry/IsSession"}
	}

	return user[0].UserLogin, nil
}
