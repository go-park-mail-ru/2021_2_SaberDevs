package handler

import (
	"net/http"

	"github.com/go-park-mail-ru/2021_2_SaberDevs/internal/session/models"
	sbErr "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/syberErrors"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
)

type SessionHandler struct {
	SessionUsecase models.SessionUsecase
}

func NewSessionHandler(su models.SessionUsecase) *SessionHandler {
	return &SessionHandler{su}
}

func (api *SessionHandler) CheckSession(c echo.Context) error {
	ctx := c.Request().Context()

	cookie, err := c.Cookie("session")
	if err != nil {

		return sbErr.ErrNoSession{
			Reason:   err.Error(),
			Function: "sessionHandler/CheckSession",
		}
	}

	response, err := api.SessionUsecase.IsSession(ctx, cookie.Value)
	if err != nil {
		return errors.Wrap(err, "sessionHandler/CheckSession")
	}

	return c.JSON(http.StatusOK, response)
}
