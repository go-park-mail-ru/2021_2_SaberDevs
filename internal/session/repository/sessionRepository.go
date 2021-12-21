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

func myInsert(tr *tarantool.Connection, path string, space interface{}, tuple interface{}) (resp *tarantool.Response, err error) {
	//TODO Metrics
	result, err := tr.Insert(space, tuple)
	return result, err
}

func (r *sessionTarantoolRepo) CreateSession(ctx context.Context, login string) (string, error) {
	sessionID := uuid.NewV4().String()

	// _, err := r.conn.Insert("sessions", []interface{}{sessionID, login})
	_, err := myInsert(r.conn, "path", "sessions", []interface{}{sessionID, login})
	if err != nil {
		return "", sbErr.ErrInternal{
			Reason:   err.Error(),
			Function: "sessionRepositiry/CreateSession"}
	}
	//_, err = r.conn.Insert("keys", []interface{}{sessionID, login})
	_, err = myInsert(r.conn, "path", "keys", []interface{}{sessionID, login})
	// if err != nil {

	// }

	return sessionID, nil
}

func (r *sessionTarantoolRepo) DeleteSession(ctx context.Context, sessionID string) error {
	_, err := r.conn.Delete("sessions", "primary", []interface{}{sessionID})
	if err != nil {
		return sbErr.ErrInternal{
			Reason:   err.Error(),
			Function: "sessionRepositiry/GetSessionLogin"}
	}

	return nil
}

func (r *sessionTarantoolRepo) GetSessionLogin(ctx context.Context, sessionID string) (string, error) {
	var user []smodels.Session

	err := r.conn.SelectTyped("sessions", "primary", 0, 1, tarantool.IterEq, []interface{}{sessionID}, &user)
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
