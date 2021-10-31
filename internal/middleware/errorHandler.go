package middleware

import (
	"net/http"

	errResp "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/errResponses"
	sbErr "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/syberErrors"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
)

func ErrorHandler(err error, c echo.Context) {
	var responseCode int
	var responseBody errResp.ErrorResponse

	switch {
	case errors.As(err, &sbErr.ErrUserDoesntExist{}):
		responseCode = http.StatusUnprocessableEntity
		responseBody = errResp.ErrUserDoesntExist

	case errors.As(err, &sbErr.ErrWrongPassword{}):
		responseCode = http.StatusUnprocessableEntity
		responseBody = errResp.ErrWrongPassword

	case errors.As(err, &sbErr.ErrUnpackingJSON{}):
		responseCode = http.StatusUnprocessableEntity
		responseBody = errResp.ErrUnpackingJSON

	case errors.As(err, &sbErr.ErrUserExists{}):
		responseCode = http.StatusUnprocessableEntity
		responseBody = errResp.ErrUserExists

	case errors.As(err, &sbErr.ErrAuthorised{}):
		responseCode = http.StatusFailedDependency
		responseBody = errResp.ErrAuthorised

	case errors.As(err, &sbErr.ErrNotLoggedin{}):
		responseCode = http.StatusFailedDependency
		responseBody = errResp.ErrNotLoggedin

	case errors.As(err, &sbErr.ErrNotFeedNumber{}):
		responseCode = http.StatusNotFound
		responseBody = errResp.ErrNotFeedNumber

	case errors.As(err, &sbErr.ErrInvalidEmail{}):
		responseCode = http.StatusFailedDependency
		responseBody = errResp.ErrInvalidEmail

	case errors.As(err, &sbErr.ErrInvalidPassword{}):
		responseCode = http.StatusFailedDependency
		responseBody = errResp.ErrInvalidPassword

	case errors.As(err, &sbErr.ErrInvalidLogin{}):
		responseCode = http.StatusFailedDependency
		responseBody = errResp.ErrInvalidLogin

	case errors.As(err, &sbErr.ErrNoSession{}):
		responseCode = http.StatusFailedDependency
		responseBody = errResp.ErrNoSession

	case errors.As(err, &sbErr.ErrDbError{}):
		responseCode = http.StatusFailedDependency
		responseBody = errResp.ErrDbFailure

	default:
		responseCode = http.StatusInternalServerError
		responseBody = errResp.ErrInternal
	}

	_err := c.JSON(responseCode, responseBody)
	if _err != nil {
		return
	}
}
