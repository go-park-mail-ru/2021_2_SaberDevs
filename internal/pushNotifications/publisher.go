package pushNotifications

import (
	"context"
	"encoding/json"
	"github.com/SherClockHolmes/webpush-go"
	pnmodels "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/pushNotifications/models"
)

func NotificationSevice(r pnmodels.PushNotificationRepository) {
	for {
		currentMsg, err := r.DequeueArticleComment()
		if err == nil {
			commentModel := pnmodels.PushComment{}
			json.Unmarshal([]byte(currentMsg), &commentModel)
			subscription, _ := r.GetSubscription(context.Background(), commentModel.Login)

			notificationModel := pnmodels.PushCommentNotification{
				To:   commentModel.Login,
				Type: 1,
				Data: commentModel,
			}

			byteNotification, _ := json.Marshal(notificationModel)

			_, err := webpush.SendNotification(byteNotification, &subscription, &webpush.Options{
				Subscriber:      "example@example.com",
				VAPIDPublicKey:  "BAm53SFQL61CJdkPZYxN4qcdNTpnRc5yVSrL182-GNHW1RYmgRSeHoF5rYdMUfZMGT93MzVsN64NBe0azXKcplM",
				VAPIDPrivateKey: "fjYQPyyzqYN4Kh_b76obpmfKkMSAz48YnbvKamm9Azw",
				TTL:             30,
			})
			if err != nil {
				continue
			}

			// resp.Body.Close()
		}
	}
}
