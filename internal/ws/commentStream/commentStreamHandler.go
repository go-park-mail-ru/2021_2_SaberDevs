package commentStream

import (
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
)

type commentStreamHandler struct {
	pub *Publisher
}

// func serveWs(pub *Publisher, w http.ResponseWriter, r *http.Request) error {
// 	conn, err := upgrader.Upgrade(w, r, nil)
// 	if err != nil {
// 		log.Println(err)
// 		return err
// 	}
// 	client := &Subscriber{
// 		pub: pub,
// 		conn: conn,
// 		send: make(chan []string),
// 	}
// 	client.hub.register <- client
//
// 	// Allow collection of memory referenced by the caller by doing all work in
// 	// new goroutines.
// 	go client.writePump()
// 	go client.readPump()
// }

func (api *commentStreamHandler) HandleWS(c echo.Context) error {
	conn, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	// коннект закрывается в горутинах
	if err != nil {
		return err
	}

	sub := &Subscriber{
		pub: api.pub,
		conn: conn,
		send: make(chan []string),
	}
	sub.pub.register <- sub


	go sub.writeWS()
	go sub.readWS()
}
