package middleware

import (
	"github.com/labstack/echo/v4"
)

func ValidateRequestBody(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		switch c.Path() {
		case "/login":
		case "/signup":
			// user := new(models.User)
			// err := c.Bind(user)
			// ok, err := valid.ValidateStruct(user)
			// if err != nil {
			//
			// }
			// if !ok {
			// 	return c.String(http.StatusBadRequest, "ok")
			// }
		}
		return next(c)
	}
}
