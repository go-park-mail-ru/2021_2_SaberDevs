package middleware

import (
	"net/http"
	"time"

	wrapper "github.com/go-park-mail-ru/2021_2_SaberDevs/internal"
	"github.com/labstack/echo/v4"
)

var layer = "request"

func Logging(r *http.Request, c echo.Context, start time.Time) {
	Id := r.Header.Get(echo.HeaderXRequestID)
	time := time.Since(start)
	wrapper.Duration.WithLabelValues(layer, c.Request().RequestURI).Observe(float64(time.Milliseconds()))
	wrapper.Hits.WithLabelValues(layer, c.Request().RequestURI, r.Method).Inc()
	c.Logger().Info("Id = ", Id, " method = ", r.Method, " address = ", r.RemoteAddr, " RequestUri = ", r.RequestURI, " Request Time = ", time)
}

func AccessLogger(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		start := time.Now()
		err := next(c)
		Logging(c.Request(), c, start)
		return err
	}
}
