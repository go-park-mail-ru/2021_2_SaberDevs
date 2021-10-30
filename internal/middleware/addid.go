package middleware

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func AddId(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Request().Header.Set(echo.HeaderXRequestID, uuid.New().String())
		return next(c)
	}
}
