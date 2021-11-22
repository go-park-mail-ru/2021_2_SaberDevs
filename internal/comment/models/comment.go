package models

import "context"

type Comment struct {
	Id          int    `json:"Id"  db:"id"`
	DateTime    string `json:"datetime" db:"datetime"`
	Text        string `json:"text" db:"text"`
	AuthorLogin string `json:"authorLogin" db:"authorlogin"`
	ArticleId   string `json:"articleIdd" db:"articleid"`
	ParentId    string `json:"parentId" db:"parentid"`
	IsEdited    bool   `json:"isEdited" db:"isedited"`
}

type PreparedComment struct {
	Id        int    `json:"Id"  db:"id"`
	DateTime  string `json:"datetime" db:"datetime"`
	Text      string `json:"text" db:"text"`
	ArticleId string `json:"articleIdd" db:"articleid"`
	ParentId  string `json:"parentId" db:"parentid"`
	IsEdited  bool   `json:"isEdited" db:"isedited"`
	Author    Author `json:"author"`
}

type Author struct {
	Login     string `json:"login"`
	Surname   string `json:"lastName"`
	Name      string `json:"firstName"`
	Score     int    `json:"score"`
	AvatarURL string `json:"avatarUrl"`
}

type Response struct {
	Status uint        `json:"status"`
	Data   interface{} `json:"data"`
	Msg    string      `json:"msg"`
}

// -----------------------------------------------

type CommentUsecase interface {
	CreateComment(ctx context.Context, comment *Comment, sessionID string) (Response, error)
	UpdateComment(ctx context.Context, comment *Comment, sessionID string) (Response, error)
	GetCommentsByArticleID(ctx context.Context, articleID string) (Response, error)
}

type CommentRepository interface {
	StoreComment(ctx context.Context, comment *Comment) (Comment, error)
	UpdateComment(ctx context.Context, comment *Comment) (Comment, error)
	GetCommentsByArticleID(ctx context.Context, articleID string, lastCommentID string) ([]PreparedComment, error)
	GetCommentByID(ctx context.Context, commentID string) (Comment, error)
}
