package handler

import (
	errResp "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/errResponses"
	"github.com/go-park-mail-ru/2021_2_SaberDevs/internal/session/models"
	"github.com/labstack/echo/v4"
	"net/http"
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
		return c.JSON(http.StatusNotAcceptable, errResp.ErrNoSession)
	}

	response, err := api.SessionUsecase.IsSession(ctx, cookie.Value)
	if err != nil {
		return c.JSON(http.StatusNotAcceptable, errResp.ErrNoSession)
	}
	return c.JSON(http.StatusOK, response)
}