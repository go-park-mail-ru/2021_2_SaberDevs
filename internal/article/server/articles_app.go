package main

import (
	"context"
	"sync"

	server "github.com/go-park-mail-ru/2021_2_SaberDevs/cmd/sybernews"
	app "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/article/article_app"
	"github.com/go-park-mail-ru/2021_2_SaberDevs/internal/article/models"
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

func previewConv(a models.Preview) *app.Preview {
	val := new(app.Preview)
	val.Category = a.Category
	val.Comments = int64(a.Comments)
	val.CommentsUrl = a.CommentsUrl
	val.DateTime = a.DateTime
	val.Id = a.Id
	val.Likes = int64(a.Likes)
	val.PreviewUrl = a.PreviewUrl
	val.Tags = a.Tags
	val.Text = a.Text
	val.Title = a.Title
	val.Author.Id = int64(a.Author.Id)
	val.Author.AvatarUrl = a.Author.AvatarUrl
	val.Author.Description = a.Author.Description
	val.Author.Email = a.Author.Email
	val.Author.Login = a.Author.Login
	val.Author.Name = a.Author.Name
	val.Author.Password = a.Author.Password
	val.Author.Score = int64(a.Author.Score)
	val.Author.Surname = a.Author.Surname
	return val
}

func (m *ArticleManager) Fetch(ctx context.Context, chunk *app.Chunk) (*app.Repview, error) {
	ch := int(chunk.ChunkSize)
	id := chunk.IdLastLoaded
	m.mu.Lock()
	defer m.mu.Unlock()
	res, err := m.handler.Fetch(ctx, id, ch)
	retval := app.Repview{}
	for _, a := range res {
		val := previewConv(a)
		retval.Preview = append(retval.Preview, val)
	}
	return &retval, err
}
