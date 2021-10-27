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
	Id           int
	StringId     string `json:"id"`
	PreviewUrl   string `json:"previewUrl"`
	Title        string `json:"title"`
	Text         string `json:"text"`
	AuthorUrl    string `json:"authorUrl"`
	AuthorName   string `json:"authorName"`
	AuthorAvatar string `json:"authorAvatar"`
	CommentsUrl  string `json:"commentsUrl"`
	Comments     uint   `json:"comments"`
	Likes        uint   `json:"likes"`
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
type ArticleUseCase interface {
	Fetch(ctx context.Context, idLastLoaded string, chunkSize int) ([]Article, error)
	// GetByID(ctx context.Context, id int64) (Article, error)
	// Update(ctx context.Context, ar *Article) error
	// GetByTitle(ctx context.Context, title string) (Article, error)
	// Store(context.Context, *Article) error
	// Delete(ctx context.Context, id int64) error
}

// ArticleRepository represent the article's repository contract
type ArticleRepository interface {
	Fetch(ctx context.Context, from, chunkSize int) ([]Article, error)
	// GetByID(ctx context.Context, id int64) (Article, error)
	// GetByTitle(ctx context.Context, title string) (Article, error)
	// Update(ctx context.Context, ar *Article) error
	// Store(ctx context.Context, a *Article) error
	// Delete(ctx context.Context, id int64) error
}
