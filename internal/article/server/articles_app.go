package main

import (
	"context"
	"sync"

	server "github.com/go-park-mail-ru/2021_2_SaberDevs/cmd/sybernews"
	app "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/article/article_app"
	amodels "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/article/models"
	"github.com/jmoiron/sqlx"
	"github.com/tarantool/go-tarantool"
)

func TarantoolConnect() (*tarantool.Connection, error) {
	user, pass, addr, err := server.TarantoolConfig()
	if err != nil {
		return nil, err
	}

	opts := tarantool.Opts{User: user, Pass: pass}
	conn, err := tarantool.Connect(addr, opts)
	if err != nil {
		return nil, err
	}

	_, err = conn.Ping()
	if err != nil {
		return nil, err
	}
	return conn, nil
}

func DbConnect() (*sqlx.DB, error) {
	connStr, err := server.DbConfig()
	if err != nil {
		return nil, err
	}
	db, err := sqlx.Open("postgres", connStr)
	if err != nil {
		return db, err
	}
	err = db.Ping()
	if err != nil {
		return db, err
	}
	return db, err
}

func DbClose(db *sqlx.DB) error {
	err := db.Close()
	if err != nil {
		return err
	}
	return err
}

type ArticleManager struct {
	mu      sync.RWMutex
	handler amodels.ArticleUsecase
}

func NewArticleManager(handler amodels.ArticleUsecase) *ArticleManager {
	return &ArticleManager{
		mu:      sync.RWMutex{},
		handler: handler,
	}
}

func (m *ArticleManager) Delete(ctx context.Context, id *app.Id) (*app.Nothing, error) {
	myid := id.GetId()
	m.mu.Lock()
	defer m.mu.Unlock()
	err := m.handler.Delete(ctx, myid)
	return &app.Nothing{Dummy: true}, err
}
