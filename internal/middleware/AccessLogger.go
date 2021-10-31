package middleware

import (
	"time"

	"github.com/labstack/echo/v4"
)

func AccessLogger(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		start := time.Now()
		err := next(c)
		// if err != nil {
		// 	c.Logger().Error(err)
		// }
		Id := c.Request().Header.Get(echo.HeaderXRequestID)
		c.Logger().Info("Id = ", Id, " method = ", c.Request().Method, " address = ", c.Request().RemoteAddr, " RequestUri = ", c.Request().RequestURI, " Request Time = ", time.Since(start))
		return err
	}
}