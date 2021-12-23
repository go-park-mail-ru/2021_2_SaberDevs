package repository

import (
	"context"
	"fmt"

	"github.com/SherClockHolmes/webpush-go"
	wrapper "github.com/go-park-mail-ru/2021_2_SaberDevs/internal"
	pnmodels "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/pushNotifications/models"
	sbErr "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/syberErrors"
	"github.com/tarantool/go-tarantool"
)

type pushNotificationTarantoolRepo struct {
	conn *tarantool.Connection
	lg   *wrapper.MyLogger
}

func NewPushNotificationRepository(conn *tarantool.Connection, lg *wrapper.MyLogger) pnmodels.PushNotificationRepository {
	return &pushNotificationTarantoolRepo{conn: conn, lg: lg}
}

func (pnr *pushNotificationTarantoolRepo) StoreSubscription(ctx context.Context, subscription webpush.Subscription, login string) error {
	path := "StoreSubscription"
	// _, err := pnr.conn.Replace("subscriptions", []interface{}{login, subscription.Endpoint, subscription.Keys.Auth, subscription.Keys.P256dh})
	_, err := pnr.lg.MyReplace(pnr.conn, path, "subscriptions", []interface{}{login, subscription.Endpoint, subscription.Keys.Auth, subscription.Keys.P256dh})
	if err != nil {
		return sbErr.ErrInternal{
			Reason:   err.Error(),
			Function: "pushNotificationRepositiry/StoreSubscription"}
	}

	return nil
}

func (pnr *pushNotificationTarantoolRepo) UpdateSubscription(ctx context.Context, subscription webpush.Subscription, login string) error {
	path := "UpdateSubscription"
	_, err := pnr.lg.MyReplace(pnr.conn, path, "subscriptions", []interface{}{login, subscription.Endpoint, subscription.Keys.Auth, subscription.Keys.P256dh})
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
	path := "GetSubscription"
	var sub []pnmodels.Subscription
	err := pnr.lg.MySelectTyped(pnr.conn, path, "subscriptions", "primary", 0, 1, tarantool.IterEq, []interface{}{login}, &sub)
	if err != nil {
		fmt.Println(err.Error())
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
	path := "QueueArticleLike"
	_, err := pnr.lg.MyCall(pnr.conn, path, "articleLikesPut", []interface{}{like})
	if err != nil {
		return sbErr.ErrInternal{
			Reason:   err.Error(),
			Function: "pushNotificationRepositiry/QueueArticleLike"}
	}

	return nil
}

func (pnr *pushNotificationTarantoolRepo) DequeueArticleLike() (string, error) {
	path := "DequeueArticleLike"
	res, err := pnr.lg.MyCall(pnr.conn, path, "articleLikesTake", []interface{}{})
	if err != nil {
		return "", err
	}
	if len(res.Tuples()) == 0 {
		return "", sbErr.ErrInternal{
			Reason:   err.Error(),
			Function: "pushNotificationRepositiry/DequeueArticleLike"}
	}
	fmt.Println(res.Tuples()[0][2].(string))

	return res.Tuples()[0][2].(string), nil
}

func (pnr *pushNotificationTarantoolRepo) QueueArticleComment(comment []byte) error {
	path := "QueueArticleComment"
	_, err := pnr.lg.MyCall(pnr.conn, path, "articleCommentPut", []interface{}{comment})
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	return nil
}

func (pnr *pushNotificationTarantoolRepo) DequeueArticleComment() (string, error) {
	path := "DequeueArticleComment"
	res, err := pnr.lg.MyCall(pnr.conn, path, "articleCommentTake", []interface{}{})
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
