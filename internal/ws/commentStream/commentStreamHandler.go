package commentStream

import (
	cmodels "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/comment/models"
	sbErr "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/syberErrors"
	"github.com/labstack/echo/v4"
)

type commentStreamHandler struct {
	pub *Publisher
}

func NewCommentStreamHandler(p *Publisher) *commentStreamHandler {
	return &commentStreamHandler{
		pub: p,
	}
}

func (api *commentStreamHandler) HandleWS(c echo.Context) error {
	conn, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	// коннект закрывается в горутинах
	if err != nil {
		return sbErr.ErrInternal{
			Reason:   err.Error(),
			Function: "ws/commentStreamHandler/HandleWS",
		}
	}

	sub := &Subscriber{
		pub:  api.pub,
		conn: conn,
		send: make(chan []cmodels.StreamComment),
	}
	sub.pub.register <- sub

	go sub.writeWS()
	go sub.readWS()
	return nil
}
