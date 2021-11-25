package commentStream

import (
	sbErr "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/syberErrors"
	"github.com/labstack/echo/v4"
)

type Comment struct {
	Type        string `json:"type"`
	Id          int64  `json:"id"  db:"id"`
	Text        string `json:"text" db:"text"`
	ArticleId   int64  `json:"articleId" db:"articleid"`
	ArticleName string `json:"articleName"`
}

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
		send: make(chan []Comment),
	}
	sub.pub.register <- sub

	go sub.writeWS()
	go sub.readWS()
	return nil
}
