package middleware

import (
	"context"

	smodels "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/session/models"
	sbErr "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/syberErrors"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
)

type AuthMiddleware struct {
	sessionRepo smodels.SessionRepository
}

func NewAuthMiddleware(r smodels.SessionRepository) AuthMiddleware {
	return AuthMiddleware{r}
}

func (m *AuthMiddleware) CheckAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cookie, err := c.Cookie("session")
		if err != nil {
			return sbErr.ErrNoSession{
				Reason:   "no session",
				Function: "middleware/CheckAuth",
			}
		}

		_, err = m.sessionRepo.GetSessionLogin(context.Background(), cookie.Value)
		if err != nil {
			return errors.Wrap(err, "middleware/CheckAuth")
		}

		return next(c)
	}
}
