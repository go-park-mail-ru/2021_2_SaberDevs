package models

import (
	"context"
	"github.com/SherClockHolmes/webpush-go"
)

type Subscription struct {
	_msgpack struct{} `msgpack:",asArray"`
}

type PushNotificationUsecase interface {
	CreateSubscription(ctx context.Context, subscription webpush.Subscription) error
	DeleteSubscription(ctx context.Context) error
}

type PushNotificationRepository interface {
	StoreSubscription(ctx context.Context, subscription webpush.Subscription) (string, error)
	DeleteSubscription(ctx context.Context) error
	GetSubscription(ctx context.Context, login string) (string, error)
	QueueArticleLike() error
	DequeueArticleLike() error
	QueueArticleComment() error
	DequeueArticleComment() error
}
