package handler

import (
	"github.com/SherClockHolmes/webpush-go"
	"github.com/go-park-mail-ru/2021_2_SaberDevs/internal/pushNotifications/models"
	sbErr "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/syberErrors"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"net/http"
)

type PushNotification struct {
	pushNotificationUsecase models.PushNotificationUsecase
}

func NewPushNotificationHandler(pnu models.PushNotificationUsecase) *PushNotification {
	return &PushNotification{pnu}
}

func (api *PushNotification) CreateSubscription(c echo.Context) error {
	ctx := c.Request().Context()

	cookie, err := c.Cookie("session")
	if err != nil {
		return sbErr.ErrNoSession{
			Reason:   err.Error(),
			Function: "PushNotificationHandler/CreateSubscription",
		}
	}

	subObject := webpush.Subscription{}

	err = c.Bind(&subObject)
	if err != nil {
		return sbErr.ErrUnpackingJSON{
			Reason:   err.Error(),
			Function: "PushNotificationHandler/CreateSubscription",
		}
	}

	err = api.pushNotificationUsecase.CreateSubscription(ctx, subObject, cookie.Value)
	if err != nil {
		return errors.Wrap(err, "PushNotificationHandler/CreateSubscription")
	}

	response := models.Response{
		Status: http.StatusOK,
		Data:   nil,
		Msg:    "OK",
	}

	return c.JSON(http.StatusOK, response)
}

func (api *PushNotification) UpdateSubscription(c echo.Context) error {
	ctx := c.Request().Context()

	cookie, err := c.Cookie("session")
	if err != nil {
		return sbErr.ErrNoSession{
			Reason:   err.Error(),
			Function: "PushNotificationHandler/UpdateSubscription",
		}
	}

	subObject := webpush.Subscription{}

	err = c.Bind(&subObject)
	if err != nil {
		return sbErr.ErrUnpackingJSON{
			Reason:   err.Error(),
			Function: "PushNotificationHandler/UpdateSubscription",
		}
	}

	err = api.pushNotificationUsecase.CreateSubscription(ctx, subObject, cookie.Value)
	if err != nil {
		return errors.Wrap(err, "PushNotificationHandler/CreateSubscription")
	}

	response := models.Response{
		Status: http.StatusOK,
		Data:   nil,
		Msg:    "OK",
	}

	return c.JSON(http.StatusOK, response)
}
