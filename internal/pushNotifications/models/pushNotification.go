package models

import (
	"context"
	"github.com/SherClockHolmes/webpush-go"
)

type Subscription struct {
	_msgpack struct{} `msgpack:",asArray"`
	Endpoint string
	Auth     string
	P256dh   string
}

type PushCommentNotification struct {
	To string `json:"to"`
	Type int `json:"type"`
	Data PushComment `json:"data"`
}

type PushComment struct {
	Login        string `json:"login"`
	ArticleTitle string `json:"articleTitle"`
	Text         string `json:"commentText" db:"text"`
	ArticleId    int64  `json:"articleId" db:"articleid"`
	CommentID    int64  `json:"commentId"`
	FirstName    string `json:"firstName"`
	LastName     string `json:"lastName"`
}

type Response struct {
	Status uint        `json:"status"`
	Data   interface{} `json:"data"`
	Msg    string      `json:"msg"`
}

type PushNotificationUsecase interface {
	CreateSubscription(ctx context.Context, subscription webpush.Subscription, sessionID string) error
	UpdateSubscription(ctx context.Context, subscription webpush.Subscription, sessionID string) error
	DeleteSubscription(ctx context.Context) error
}

type PushNotificationRepository interface {
	StoreSubscription(ctx context.Context, subscription webpush.Subscription, login string) error
	UpdateSubscription(ctx context.Context, subscription webpush.Subscription, login string) error
	DeleteSubscription(ctx context.Context) error
	GetSubscription(ctx context.Context, login string) (webpush.Subscription, error)
	QueueArticleLike(like []byte) error
	DequeueArticleLike() (string, error)
	QueueArticleComment(comment []byte) error
	DequeueArticleComment() (string, error)
}
