package usecase

import (
	"context"
	"github.com/SherClockHolmes/webpush-go"
	pnmodels "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/pushNotifications/models"
)

type pushNotificationUsecase struct {
	pnRepo pnmodels.PushNotificationRepository
}

func NewPushNotificationUsecase(pnr pnmodels.PushNotificationRepository) pnmodels.PushNotificationUsecase {
	return &pushNotificationUsecase{
		pnRepo: pnr,
	}
}

func (pnu *pushNotificationUsecase) CreateSubscription(ctx context.Context, subscription webpush.Subscription) error {

}

func (pnu *pushNotificationUsecase) DeleteSubscription(ctx context.Context) error {

}
