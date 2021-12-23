package repository

import (
	"context"

	wrapper "github.com/go-park-mail-ru/2021_2_SaberDevs/internal"
	sbErr "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/syberErrors"
	"github.com/tarantool/go-tarantool"

	smodels "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/session/models"
	uuid "github.com/satori/go.uuid"
)

type sessionTarantoolRepo struct {
	conn *tarantool.Connection
	lg   *wrapper.MyLogger
}

func NewSessionRepository(conn *tarantool.Connection, lg *wrapper.MyLogger) smodels.SessionRepository {
	return &sessionTarantoolRepo{conn: conn, lg: lg}
}

func (r *sessionTarantoolRepo) CreateSession(ctx context.Context, login string) (string, error) {
	path := "CreateSession"
	sessionID := uuid.NewV4().String()
	_, err := r.lg.MyInsert(r.conn, path, "sessions", []interface{}{sessionID, login})
	if err != nil {
		return "", sbErr.ErrInternal{
			Reason:   err.Error(),
			Function: "sessionRepositiry/CreateSession"}
	}
	_, err = r.lg.MyInsert(r.conn, "path", "keys", []interface{}{sessionID, login})
	// if err != nil {

	// }

	return sessionID, nil
}

func (r *sessionTarantoolRepo) DeleteSession(ctx context.Context, sessionID string) error {
	path := "DeleteSession"
	_, err := r.lg.MyDelete(r.conn, path, "sessions", "primary", []interface{}{sessionID})
	if err != nil {
		return sbErr.ErrInternal{
			Reason:   err.Error(),
			Function: "sessionRepositiry/GetSessionLogin"}
	}

	return nil
}

func (r *sessionTarantoolRepo) GetSessionLogin(ctx context.Context, sessionID string) (string, error) {
	path := "GetSessionLogin"
	var user []smodels.Session
	err := r.lg.MySelectTyped(r.conn, path, "sessions", "primary", 0, 1, tarantool.IterEq, []interface{}{sessionID}, &user)
	if err != nil {
		return "", sbErr.ErrInternal{
			Reason:   err.Error(),
			Function: "sessionRepositiry/GetSessionLogin"}
	}
	if len(user) == 0 {
		return "", sbErr.ErrNoSession{
			Reason:   "no session",
			Function: "sessionRepositiry/GetSessionLogin"}
	}

	return user[0].UserLogin, nil
}
