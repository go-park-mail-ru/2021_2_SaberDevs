package models

import "context"

type Comment struct {
	Id          int64  `json:"Id"  db:"id"`
	DateTime    string `json:"datetime" db:"datetime"`
	Text        string `json:"text" db:"text"`
	AuthorLogin string `json:"authorLogin" db:"authorlogin"`
	ArticleId   int64  `json:"articleId" db:"articleid"`
	ParentId    int64  `json:"parentId" db:"parentid"`
	IsEdited    bool   `json:"isEdited" db:"isedited"`
}

type PreparedComment struct {
	Id        int64  `json:"Id"  db:"id"`
	DateTime  string `json:"datetime" db:"datetime"`
	Text      string `json:"text" db:"text"`
	ArticleId int64  `json:"articleIdd" db:"articleid"`
	ParentId  int64  `json:"parentId" db:"parentid"`
	IsEdited  bool   `json:"isEdited" db:"isedited"`
	Author    Author `json:"author"`
}

type Author struct {
	Login     string `json:"login" db:"login"`
	Surname   string `json:"lastName" db:"surname"`
	Name      string `json:"firstName" db:"name"`
	Score     int    `json:"score" db:"score"`
	AvatarURL string `json:"avatarUrl" db:"avatarurl"`
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
	GetCommentsByArticleID(ctx context.Context, articleID int64) (Response, error)
}

type CommentRepository interface {
	StoreComment(ctx context.Context, comment *Comment) (Comment, error)
	UpdateComment(ctx context.Context, comment *Comment) (Comment, error)
	GetCommentsByArticleID(ctx context.Context, articleID int64, lastCommentID int64) ([]PreparedComment, error)
	GetCommentByID(ctx context.Context, commentID int64) (Comment, error)
}
