package repository

import (
	"context"
	"github.com/SherClockHolmes/webpush-go"
	pnmodels "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/pushNotifications/models"
	"github.com/tarantool/go-tarantool"
)

type pushNotificationTarantoolRepo struct {
	conn *tarantool.Connection
}

func NewPushNotificationRepository(conn *tarantool.Connection) pnmodels.PushNotificationRepository {
	return &pushNotificationTarantoolRepo{conn: conn}
}

func (pnr *pushNotificationTarantoolRepo) StoreSubscription(ctx context.Context, subscription webpush.Subscription) (string, error) {

}

func (pnr *pushNotificationTarantoolRepo) DeleteSubscription(ctx context.Context) error {

}

func (pnr *pushNotificationTarantoolRepo) GetSubscription(ctx context.Context, login string) (string, error) {
}

func (pnr *pushNotificationTarantoolRepo) QueueArticleLike() error    {

}

func (pnr *pushNotificationTarantoolRepo) DequeueArticleLike() error  {

}

func (pnr *pushNotificationTarantoolRepo) QueueArticleComment() error {

}

func (pnr *pushNotificationTarantoolRepo) DequeueArticleComment() error {

}
