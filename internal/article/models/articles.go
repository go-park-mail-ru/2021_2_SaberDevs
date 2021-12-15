package models

import (
	"context"
)

//Представление записи
type ArticleData struct {
	PreviewUrl   string   `json:"previewUrl"`
	Tags         []string `json:"tags"`
	Title        string   `json:"title"`
	Text         string   `json:"text"`
	AuthorUrl    string   `json:"authorUrl"`
	AuthorName   string   `json:"authorName"`
	AuthorAvatar string   `json:"authorAvatar"`
	CommentsUrl  string   `json:"commentsUrl"`
	Comments     uint     `json:"comments"`
	Likes        int64    `json:"likes"`
}

type Article struct {
	Id           string   `json:"id"`
	DateTime     string   `json:"datetime" db:"datetime"`
	PreviewUrl   string   `json:"previewUrl"`
	Tags         []string `json:"tags"`
	Title        string   `json:"title"`
	Category     string   `json:"category"`
	Text         string   `json:"text"`
	AuthorUrl    string   `json:"authorUrl"`
	AuthorName   string   `json:"authorName"`
	AuthorAvatar string   `json:"authorAvatar"`
	CommentsUrl  string   `json:"commentsUrl"`
	Comments     uint     `json:"comments"`
	Likes        int64    `json:"likes"`
	Liked        int64    `json:"liked"`
}

//easyjson:json
type FullArticle struct {
	Id          string   `json:"id"`
	DateTime    string   `json:"datetime" db:"datetime"`
	PreviewUrl  string   `json:"previewUrl"`
	Tags        []string `json:"tags"`
	Title       string   `json:"title"`
	Category    string   `json:"category"`
	Text        string   `json:"text"`
	Author      Author   `json:"author"`
	CommentsUrl string   `json:"commentsUrl"`
	Comments    uint     `json:"comments"`
	Likes       int64    `json:"likes"`
	Liked       int64    `json:"liked"`
}

//easyjson:json
type Preview struct {
	Id          string   `json:"id"`
	DateTime    string   `json:"datetime" db:"datetime"`
	PreviewUrl  string   `json:"previewUrl"`
	Tags        []string `json:"tags"`
	Title       string   `json:"title"`
	Category    string   `json:"category"`
	Text        string   `json:"text"`
	Author      Author   `json:"author"`
	CommentsUrl string   `json:"commentsUrl"`
	Comments    uint     `json:"comments"`
	Likes       int64    `json:"likes"`
	Liked       int64    `json:"liked"`
}

type DbArticle struct {
	Id          int    `json:"Id"  db:"id"`
	DateTime    string `json:"datetime" db:"datetime"`
	PreviewUrl  string `json:"PreviewUrl" db:"previewurl"`
	Title       string `json:"title" db:"title"`
	Category    string `json:"category" db:"category"`
	Text        string `json:"text" db:"text"`
	AuthorName  string `json:"authorName" db:"authorname"`
	CommentsUrl string `json:"commentsUrl" db:"commentsurl"`
	Comments    uint   `json:"comments" db:"comments"`
	Likes       int64  `json:"likes" db:"likes"`
}

//easyjson:json
type ChunkResponse struct {
	Status    uint      `json:"status"`
	ChunkData []Preview `json:"data"`
}

//easyjson:json
type AuthorsChunks struct {
	Status    uint     `json:"status"`
	ChunkData []Author `json:"data"`
}

//easyjson:json
type ArticleResponse struct {
	Status uint        `json:"status"`
	Data   FullArticle `json:"data"`
}

//easyjson:json
type GenericResponse struct {
	Status uint   `json:"status"`
	Data   string `json:"data"`
}

//easyjson:json
type ArticleCreate struct {
	Title    string   `json:"title" db:"title"`
	Text     string   `json:"text" db:"text"`
	Category string   `json:"category" db:"category"`
	Img      string   `json:"img" db:"img"`
	Tags     []string `json:"tags"`
}

//easyjson:json
type ArticleUpdate struct {
	Id       string   `json:"id"  db:"id"`
	Title    string   `json:"title" db:"title"`
	Text     string   `json:"text" db:"text"`
	Category string   `json:"category" db:"category"`
	Img      string   `json:"img" db:"img"`
	Tags     []string `json:"tags"`
}

type Author struct {
	Id          int    `json:"-"`
	Login       string `json:"login"`
	Name        string `json:"firstName"`
	Surname     string `json:"lastName"`
	AvatarUrl   string `json:"avatarUrl" db:"avatarurl"`
	Description string `json:"description" db:"description"`
	Email       string `json:"-" valid:"email,optional"`
	Password    string `json:"-"`
	Score       int    `json:"score"`
}

// ArticleUsecase represent the article's usecases
type ArticleUsecase interface {
	Fetch(ctx context.Context, c string, idLastLoaded string, chunkSize int) ([]Preview, error)
	GetByID(ctx context.Context, c string, id int64) (FullArticle, error)
	GetByTag(ctx context.Context, c string, tag string, idLastLoaded string, chunkSize int) ([]Preview, error)
	GetByAuthor(ctx context.Context, c string, author string, idLastLoaded string, chunkSize int) ([]Preview, error)
	GetByCategory(ctx context.Context, c string, category string, idLastLoaded string, chunkSize int) ([]Preview, error)
	FindByTag(ctx context.Context, c string, query string, idLastLoaded string, chunkSize int) ([]Preview, error)
	FindAuthors(ctx context.Context, query string, idLastLoaded string, chunkSize int) ([]Author, error)
	FindArticles(ctx context.Context, c string, query string, idLastLoaded string, chunkSize int) ([]Preview, error)
	Update(ctx context.Context, c string, a *ArticleUpdate) error
	Store(ctx context.Context, c string, a *ArticleCreate) (int, error)
	Delete(ctx context.Context, c string, id string) error
}

// ArticleRepository represent the article's repository contract
type ArticleRepository interface {
	Fetch(ctx context.Context, login string, from, chunkSize int) ([]Preview, error)
	GetByID(ctx context.Context, login string, id int64) (FullArticle, error)
	GetByTag(ctx context.Context, login string, tag string, from, chunkSize int) ([]Preview, error)
	GetByAuthor(ctx context.Context, login string, author string, from, chunkSize int) ([]Preview, error)
	GetByCategory(ctx context.Context, login string, category string, from, chunkSize int) ([]Preview, error)
	FindByTag(ctx context.Context, login string, query string, from, chunkSize int) ([]Preview, error)
	FindAuthors(ctx context.Context, query string, from, chunkSize int) ([]Author, error)
	FindArticles(ctx context.Context, login string, query string, from, chunkSize int) ([]Preview, error)
	Update(ctx context.Context, a *Article) error
	Store(ctx context.Context, a *Article) (int, error)
	Delete(ctx context.Context, author string, id int64) error
}
