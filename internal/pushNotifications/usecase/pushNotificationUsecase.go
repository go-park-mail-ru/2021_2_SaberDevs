package usecase

import (
	"context"
	"github.com/SherClockHolmes/webpush-go"
	pnmodels "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/pushNotifications/models"
	smodels "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/session/models"
	"github.com/pkg/errors"
)

type pushNotificationUsecase struct {
	pnRepo      pnmodels.PushNotificationRepository
	sessionRepo smodels.SessionRepository
}

func NewPushNotificationUsecase(pnr pnmodels.PushNotificationRepository, sr smodels.SessionRepository) pnmodels.PushNotificationUsecase {
	return &pushNotificationUsecase{
		pnRepo:      pnr,
		sessionRepo: sr,
	}
}

func (pnu *pushNotificationUsecase) CreateSubscription(ctx context.Context, subscription webpush.Subscription, sessionID string) error {
	login, err := pnu.sessionRepo.GetSessionLogin(ctx, sessionID)
	if err != nil {
		return errors.Wrap(err, "pushNotificationUsecase/CreateSubscription")
	}

	err = pnu.pnRepo.StoreSubscription(ctx, subscription, login)
	if err != nil {
		return errors.Wrap(err, "pushNotificationUsecase/CreateSubscription")
	}

	return nil
}

func (pnu *pushNotificationUsecase) UpdateSubscription(ctx context.Context, subscription webpush.Subscription, sessionID string) error {
	login, err := pnu.sessionRepo.GetSessionLogin(ctx, sessionID)
	if err != nil {
		return errors.Wrap(err, "pushNotificationUsecase/UpdateSubscription")
	}

	err = pnu.pnRepo.StoreSubscription(ctx, subscription, login)
	if err != nil {
		return errors.Wrap(err, "pushNotificationUsecase/UpdateSubscription")
	}

	return nil
}

func (pnu *pushNotificationUsecase) DeleteSubscription(ctx context.Context) error {
	return nil
}
