package repository

import (
	"context"
	"github.com/SherClockHolmes/webpush-go"
	pnmodels "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/pushNotifications/models"
	sbErr "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/syberErrors"
	"github.com/tarantool/go-tarantool"
)

type pushNotificationTarantoolRepo struct {
	conn *tarantool.Connection
}

func NewPushNotificationRepository(conn *tarantool.Connection) pnmodels.PushNotificationRepository {
	return &pushNotificationTarantoolRepo{conn: conn}
}

func (pnr *pushNotificationTarantoolRepo) StoreSubscription(ctx context.Context, subscription webpush.Subscription, login string) error {
	_, err := pnr.conn.Insert("subscription", []interface{}{login, subscription.Endpoint, subscription.Keys.Auth, subscription.Keys.P256dh})
	if err != nil {
		return sbErr.ErrInternal{
			Reason:   err.Error(),
			Function: "pushNotificationRepositiry/StoreSubscription"}
	}

	return nil
}

func (pnr *pushNotificationTarantoolRepo) UpdateSubscription(ctx context.Context, subscription webpush.Subscription, login string) error {
	_, err := pnr.conn.Replace("subscription", []interface{}{login, subscription.Endpoint, subscription.Keys.Auth, subscription.Keys.P256dh})
	if err != nil {
		return sbErr.ErrInternal{
			Reason:   err.Error(),
			Function: "pushNotificationRepositiry/UpdateSubscription"}
	}

	return nil
}

func (pnr *pushNotificationTarantoolRepo) DeleteSubscription(ctx context.Context) error {
	return nil
}

func (pnr *pushNotificationTarantoolRepo) GetSubscription(ctx context.Context, login string) (webpush.Subscription, error) {
	var sub []pnmodels.Subscription

	err := pnr.conn.SelectTyped("subscription", "primary", 0, 1, tarantool.IterEq, []interface{}{login}, &sub)
	if err != nil {
		return webpush.Subscription{}, sbErr.ErrInternal{
			Reason:   err.Error(),
			Function: "pushNotificationRepositiry/GetSubscription"}
	}
	if len(sub) == 0 {
		return webpush.Subscription{}, sbErr.ErrNoSession{
			Reason:   "no Subscription",
			Function: "pushNotificationRepositiry/GetSubscription"}
	}

	return webpush.Subscription{
		Endpoint: sub[0].Endpoint,
		Keys: webpush.Keys{
			Auth:   sub[0].Auth,
			P256dh: sub[0].P256dh,
		},
	}, nil
}

func (pnr *pushNotificationTarantoolRepo) QueueArticleLike(like []byte) error {
	_, err := pnr.conn.Call("articleLikesPut", []interface{}{like})
	if err != nil {
		return sbErr.ErrInternal{
			Reason:   err.Error(),
			Function: "pushNotificationRepositiry/QueueArticleLike"}
	}

	return nil
}

func (pnr *pushNotificationTarantoolRepo) DequeueArticleLike() (string, error) {
	res, err := pnr.conn.Call("articleLikesTake", []interface{}{})
	if err != nil {
		return "", err
	}
	if len(res.Tuples()) == 0 {
		return "", sbErr.ErrInternal{
			Reason:   err.Error(),
			Function: "pushNotificationRepositiry/DequeueArticleLike"}
	}

	return res.Tuples()[0][2].(string), nil
}

func (pnr *pushNotificationTarantoolRepo) QueueArticleComment(comment []byte) error {
	_, err := pnr.conn.Call("articleCommentPut", []interface{}{comment})
	if err != nil {
		return err
	}

	return nil
}

func (pnr *pushNotificationTarantoolRepo) DequeueArticleComment() (string, error) {
	res, err := pnr.conn.Call("articleCommentTake", []interface{}{})
	if err != nil {
		return "", err
	}
	if len(res.Tuples()) == 0 {
		return "", sbErr.ErrInternal{
			Reason:   err.Error(),
			Function: "pushNotificationRepositiry/DequeueArticleComment"}
	}

	return res.Tuples()[0][2].(string), nil
}
