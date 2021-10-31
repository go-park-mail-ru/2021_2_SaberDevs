package models

import "context"

//Представление записи
type Article struct {
	Id           string   `json:"id"`
	PreviewUrl   string   `json:"previewUrl"`
	Tags         []string `json:"tags"`
	Title        string   `json:"title"`
	Text         string   `json:"text"`
	AuthorUrl    string   `json:"authorUrl"`
	AuthorName   string   `json:"authorName"`
	AuthorAvatar string   `json:"authorAvatar"`
	CommentsUrl  string   `json:"commentsUrl"`
	Comments     uint     `json:"comments"`
	Likes        uint     `json:"likes"`
}

type DbArticle struct {
	Id           int    `json:"Id"  db:"id"`
	StringId     string `json:"StringId"  db:"stringid"`
	PreviewUrl   string `json:"PreviewUrl" db:"previewurl"`
	Title        string `json:"title" db:"title"`
	Text         string `json:"text" db:"text"`
	AuthorUrl    string `json:"authorUrl" db:"authorurl"`
	AuthorName   string `json:"authorName" db:"authorname"`
	AuthorAvatar string `json:"authorAvatar" db:"authoravatar"`
	CommentsUrl  string `json:"commentsUrl" db:"commentsurl"`
	Comments     uint   `json:"comments" db:"comments"`
	Likes        uint   `json:"likes" db:"likes"`
}

//Тело ответа на API-call /getfeed

// type RequestChunk struct {
// 	idLastLoaded string
// 	login        string
// }

type ChunkResponse struct {
	Status    uint      `json:"status"`
	ChunkData []Article `json:"data"`
}

// ArticleUsecase represent the article's usecases
type ArticleUsecase interface {
	Fetch(ctx context.Context, idLastLoaded string, chunkSize int) ([]Article, error)
	GetByID(ctx context.Context, id int64) (Article, error)
	GetByTag(ctx context.Context, tag string) ([]Article, error)
	GetByAuthor(ctx context.Context, author string) ([]Article, error)
	Update(ctx context.Context, a *Article) error
	Store(ctx context.Context, a *Article) error
	Delete(ctx context.Context, id string) error
}

// ArticleRepository represent the article's repository contract
type ArticleRepository interface {
	Fetch(ctx context.Context, from, chunkSize int) ([]Article, error)
	GetByID(ctx context.Context, id int64) (Article, error)
	GetByTag(ctx context.Context, tag string) ([]Article, error)
	GetByAuthor(ctx context.Context, author string) ([]Article, error)
	Update(ctx context.Context, a *Article) error
	Store(ctx context.Context, a *Article) error
	Delete(ctx context.Context, id int64) error
}
