package middleware

import (
	errResp "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/errResponses"
	sbErr "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/syberErrors"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"net/http"
)

func ErrorHandler(err error, c echo.Context) {
	if errors.As(err, &sbErr.ErrUserDoesntExist{}) {
		c.JSON(http.StatusUnprocessableEntity, errResp.ErrUserDoesntExist)
		return
	}
}
