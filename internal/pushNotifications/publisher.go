package pushNotifications

import (
	"context"
	"encoding/json"
	"github.com/SherClockHolmes/webpush-go"
	pnmodels "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/pushNotifications/models"
)

func NotificationSevice(r pnmodels.PushNotificationRepository) {
	select {
	default:
		currentMsg, err := r.DequeueArticleComment()
		if err == nil {
			commentModel := pnmodels.PushComment{}
			json.Unmarshal([]byte(currentMsg), &commentModel)
			subscription, _ := r.GetSubscription(context.Background(), commentModel.Login)

			resp, err := webpush.SendNotification([]byte(currentMsg), &subscription, &webpush.Options{
				Subscriber:      "example@example.com",
				VAPIDPublicKey:  "<YOUR_VAPID_PUBLIC_KEY>",
				VAPIDPrivateKey: "<YOUR_VAPID_PRIVATE_KEY>",
				TTL:             30,
			})
			if err != nil {
				// TODO: Handle error
			}

			defer resp.Body.Close()
		}
	}
}
