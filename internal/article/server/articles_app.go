package main

import (
	"context"
	"fmt"
	"strconv"
	"sync"

	server "github.com/go-park-mail-ru/2021_2_SaberDevs/cmd/sybernews"
	app "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/article/article_app"
	"github.com/go-park-mail-ru/2021_2_SaberDevs/internal/article/models"
	amodels "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/article/models"
	"github.com/jmoiron/sqlx"
	"github.com/tarantool/go-tarantool"
	"google.golang.org/grpc/metadata"
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

func IdToString(id string) (int, error) {
	if id == "" {
		id = "0"
	}
	idInt, err := strconv.Atoi(id)
	return idInt, err
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
	err := m.handler.Delete(ctx, id.Value, myid)
	return &app.Nothing{Dummy: true}, err
}

func previewConv(a models.Preview) *app.Preview {
	val := new(app.Preview)
	val.Category = a.Category
	val.Comments = int64(a.Comments)
	val.CommentsUrl = a.CommentsUrl
	val.DateTime = a.DateTime
	val.Id = a.Id
	val.Likes = a.Likes
	val.Liked = a.Liked
	val.PreviewUrl = a.PreviewUrl
	val.Tags = a.Tags
	val.Text = a.Text
	val.Title = a.Title
	val.Author = auConv(a.Author)
	return val
}

func fullConv(a models.FullArticle) *app.FullArticle {
	val := new(app.FullArticle)
	val.Category = a.Category
	val.Comments = int64(a.Comments)
	val.CommentsUrl = a.CommentsUrl
	val.DateTime = a.DateTime
	val.Id = a.Id
	val.Likes = int64(a.Likes)
	val.Liked = a.Liked
	val.PreviewUrl = a.PreviewUrl
	val.Tags = a.Tags
	val.Text = a.Text
	val.Title = a.Title
	val.Author = auConv(a.Author)
	return val
}

func auConv(a models.Author) *app.Author {
	thor := new(app.Author)
	thor.Id = int64(a.Id)
	thor.AvatarUrl = a.AvatarUrl
	thor.Description = a.Description
	thor.Email = a.Email
	thor.Login = a.Login
	thor.Name = a.Name
	thor.Password = a.Password
	thor.Score = int64(a.Score)
	thor.Surname = thor.Surname
	return thor
}

func arConv(a *app.ArticleCreate) *models.ArticleCreate {
	ar := new(models.ArticleCreate)
	ar.Category = a.Category
	ar.Img = a.Img
	ar.Tags = a.Tags
	ar.Text = a.Text
	ar.Title = a.Title
	return ar
}

func upConv(a *app.ArticleUpdate) *models.ArticleUpdate {
	ar := new(models.ArticleUpdate)
	ar.Category = a.Category
	ar.Img = a.Img
	ar.Tags = a.Tags
	ar.Text = a.Text
	ar.Title = a.Title
	ar.Id = a.Id
	return ar
}

func (m *ArticleManager) Fetch(ctx context.Context, chunk *app.Chunk) (*app.Repview, error) {
	ch := int(chunk.ChunkSize)
	id := chunk.IdLastLoaded
	m.mu.Lock()
	defer m.mu.Unlock()
	res, err := m.handler.Fetch(ctx, chunk.Value, id, ch)
	md, ok := metadata.FromIncomingContext(ctx)
	value := md["x-request-id"]
	if ok {
		fmt.Println(value)
	}
	retval := app.Repview{}
	for _, a := range res {
		val := previewConv(a)
		fmt.Println(val.Liked)
		retval.Preview = append(retval.Preview, val)
	}
	return &retval, err
}

func (m *ArticleManager) FindArticles(ctx context.Context, q *app.Queries) (*app.Repview, error) {
	ch := int(q.Chunk.ChunkSize)
	id := q.Chunk.IdLastLoaded
	query := q.Query
	m.mu.Lock()
	defer m.mu.Unlock()
	res, err := m.handler.FindArticles(ctx, q.Chunk.Value, query, id, ch)
	md, ok := metadata.FromIncomingContext(ctx)
	value := md["x-request-id"]
	if ok {
		fmt.Println(value)
	}
	retval := app.Repview{}
	for _, a := range res {
		val := previewConv(a)
		retval.Preview = append(retval.Preview, val)
	}
	return &retval, err
}

func (m *ArticleManager) FindByTag(ctx context.Context, q *app.Queries) (*app.Repview, error) {
	ch := int(q.Chunk.ChunkSize)
	id := q.Chunk.IdLastLoaded
	query := q.Query
	m.mu.Lock()
	defer m.mu.Unlock()
	res, err := m.handler.FindByTag(ctx, q.Chunk.Value, query, id, ch)
	md, ok := metadata.FromIncomingContext(ctx)
	value := md["x-request-id"]
	if ok {
		fmt.Println(value)
	}
	retval := app.Repview{}
	for _, a := range res {
		val := previewConv(a)
		retval.Preview = append(retval.Preview, val)
	}
	return &retval, err
}

func (m *ArticleManager) GetByAuthor(ctx context.Context, au *app.Authors) (*app.Repview, error) {
	ch := int(au.Chunk.ChunkSize)
	id := au.Chunk.IdLastLoaded
	author := au.Author
	m.mu.Lock()
	defer m.mu.Unlock()
	res, err := m.handler.GetByAuthor(ctx, au.Chunk.Value, author, id, ch)
	md, ok := metadata.FromIncomingContext(ctx)
	value := md["x-request-id"]
	if ok {
		fmt.Println(value)
	}
	retval := app.Repview{}
	for _, a := range res {
		val := previewConv(a)
		retval.Preview = append(retval.Preview, val)
	}
	return &retval, err
}

func (m *ArticleManager) GetByCategory(ctx context.Context, cat *app.Categories) (*app.Repview, error) {
	ch := int(cat.Chunk.ChunkSize)
	id := cat.Chunk.IdLastLoaded
	category := cat.Category
	m.mu.Lock()
	defer m.mu.Unlock()
	res, err := m.handler.GetByCategory(ctx, cat.Chunk.Value, category, id, ch)
	md, ok := metadata.FromIncomingContext(ctx)
	value := md["x-request-id"]
	if ok {
		fmt.Println(value)
	}
	retval := app.Repview{}
	for _, a := range res {
		val := previewConv(a)
		retval.Preview = append(retval.Preview, val)
	}
	return &retval, err
}

func (m *ArticleManager) GetByTag(ctx context.Context, cat *app.Tags) (*app.Repview, error) {
	ch := int(cat.Chunk.ChunkSize)
	id := cat.Chunk.IdLastLoaded
	tag := cat.Tag
	m.mu.Lock()
	defer m.mu.Unlock()
	res, err := m.handler.GetByTag(ctx, cat.Chunk.Value, tag, id, ch)
	md, ok := metadata.FromIncomingContext(ctx)
	value := md["x-request-id"]
	if ok {
		fmt.Println(value)
	}
	retval := app.Repview{}
	for _, a := range res {
		val := previewConv(a)
		retval.Preview = append(retval.Preview, val)
	}
	return &retval, err
}

func (m *ArticleManager) GetByID(ctx context.Context, id *app.Id) (*app.FullArticle, error) {
	nId, err := IdToString(id.Id)
	retval := app.FullArticle{}
	if err != nil {
		return &retval, err
	}
	m.mu.Lock()
	defer m.mu.Unlock()
	res, err := m.handler.GetByID(ctx, id.Value, int64(nId))
	md, ok := metadata.FromIncomingContext(ctx)
	value := md["x-request-id"]
	if ok {
		fmt.Println(value)
	}
	retval = *fullConv(res)
	return &retval, err
}

func (m *ArticleManager) Store(ctx context.Context, a *app.Create) (*app.Created, error) {
	ar := arConv(a.Art)
	m.mu.Lock()
	defer m.mu.Unlock()
	res, err := m.handler.Store(ctx, a.Value, ar)
	md, ok := metadata.FromIncomingContext(ctx)
	value := md["x-request-id"]
	if ok {
		fmt.Println(value)
	}

	return &app.Created{Id: int64(res)}, err
}

func (m *ArticleManager) Update(ctx context.Context, a *app.ArticleUpdate) (*app.Nothing, error) {
	ar := upConv(a)
	m.mu.Lock()
	defer m.mu.Unlock()
	err := m.handler.Update(ctx, a.Value, ar)
	md, ok := metadata.FromIncomingContext(ctx)
	value := md["x-request-id"]
	if ok {
		fmt.Println(value)
	}
	return &app.Nothing{Dummy: true}, err
}
func (m *ArticleManager) FindAuthors(ctx context.Context, q *app.Queries) (*app.AView, error) {
	ch := int(q.Chunk.ChunkSize)
	id := q.Chunk.IdLastLoaded
	query := q.Query
	m.mu.Lock()
	defer m.mu.Unlock()
	res, err := m.handler.FindAuthors(ctx, query, id, ch)
	md, ok := metadata.FromIncomingContext(ctx)
	value := md["x-request-id"]
	if ok {
		fmt.Println(value)
	}
	retval := app.AView{}
	for _, a := range res {
		val := auConv(a)
		retval.Author = append(retval.Author, val)
	}
	return &retval, err
}
