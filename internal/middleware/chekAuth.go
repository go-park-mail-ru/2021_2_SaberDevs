package middleware

import (
	"github.com/labstack/echo/v4"
	sbErr "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/syberErrors"
)

func CheckAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cookie, err := c.Cookie("session")
		if err != nil {
			return sbErr.ErrNoSession{
				Reason:  "no session",
				Function: "middleware/CheckAuth",
			}
		}

		return next(c)
	}
}
