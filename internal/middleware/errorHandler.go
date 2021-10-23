package middleware

import (
	errResp "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/errResponses"
	sbErr "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/syberErrors"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"net/http"
)

func ErrorHandler(err error, c echo.Context) {
	switch {
	case errors.As(err, &sbErr.ErrUserDoesntExist{}):
		err := c.JSON(http.StatusUnprocessableEntity, errResp.ErrUserDoesntExist)
		if err != nil {
			return
		}

	case errors.As(err, &sbErr.ErrWrongPassword{}):
		err := c.JSON(http.StatusUnprocessableEntity, errResp.ErrWrongPassword)
		if err != nil {
			return
		}

	case errors.As(err, &sbErr.ErrUnpackingJSON{}):
		err := c.JSON(http.StatusUnprocessableEntity, errResp.ErrUnpackingJSON)
		if err != nil {
			return
		}

	case errors.As(err, &sbErr.ErrUserExists{}):
		err := c.JSON(http.StatusFailedDependency, errResp.ErrUserExists)
		if err != nil {
			return
		}

	case errors.As(err, &sbErr.ErrAuthorised{}):
		err := c.JSON(http.StatusFailedDependency, errResp.ErrAuthorised)
		if err != nil {
			return
		}

	case errors.As(err, &sbErr.ErrNotLoggedin{}):
		err := c.JSON(http.StatusFailedDependency, errResp.ErrNotLoggedin)
		if err != nil {
			return
		}

	case errors.As(err, &sbErr.ErrNotFeedNumber{}):
		err := c.JSON(http.StatusNotFound, errResp.ErrNotFeedNumber)
		if err != nil {
			return
		}

	case errors.As(err, &sbErr.ErrInvalidEmail{}):
		err := c.JSON(http.StatusFailedDependency, errResp.ErrInvalidEmail)
		if err != nil {
			return
		}

	case errors.As(err, &sbErr.ErrInvalidPassword{}):
		err := c.JSON(http.StatusFailedDependency, errResp.ErrInvalidPassword)
		if err != nil {
			return
		}

	case errors.As(err, &sbErr.ErrInvalidLogin{}):
		err := c.JSON(http.StatusFailedDependency, errResp.ErrInvalidLogin)
		if err != nil {
			return
		}

	case errors.As(err, &sbErr.ErrNoSession{}):
		err := c.JSON(http.StatusFailedDependency, errResp.ErrInvalidLogin)
		if err != nil {
			return
		}

	default:
		err := c.JSON(http.StatusInternalServerError, errResp.ErrInternal)
		if err != nil {
			return
		}
	}
}
