package commentStream

import (
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

const (
	writeWait      = 10 * time.Second
	pongWait       = 11 * time.Second
	pingPeriod     = (pongWait * 8) / 10
	maxMessageSize = 512
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type Subscriber struct {
	pub  *Publisher
	conn *websocket.Conn
	// канал для получения последних коментариев от publisher
	send chan []Comment // todo поменять на коменты
}

func (sub *Subscriber) readWS() {
	defer func() {
		sub.pub.unregister <- sub
		sub.conn.Close()
	}()
	// var c []Comment
	// c = append(c, Comment{
	// 	Type:        "stream-comment",
	// 	Id:          0,
	// 	Text:        "",
	// 	ArticleId:   0,
	// 	ArticleName: "",
	// })
	// sub.pub.broadcast <- c

	sub.conn.SetReadLimit(maxMessageSize)
	err := sub.conn.SetReadDeadline(time.Now().Add(pongWait))
	if err != nil {
		return
	}
	sub.conn.SetPongHandler(func(string) error {
		sub.conn.SetReadDeadline(time.Now().Add(pongWait))
		return nil
	})

	for {
		_, _, err := sub.conn.ReadMessage()
		if err != nil {
			break
		}
	}
}

func (sub *Subscriber) writeWS() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		sub.conn.Close()
	}()

	for {
		select {
		case message, ok := <-sub.send:
			err := sub.conn.SetWriteDeadline(time.Now().Add(writeWait)) // всегда err = nil
			if err != nil {
				return
			}

			if !ok {
				sub.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			err = sub.conn.WriteJSON(message)
			if err != nil {
				return
			}

		case <-ticker.C:
			err := sub.conn.SetWriteDeadline(time.Now().Add(writeWait)) // всегда err = nil
			if err != nil {
				return
			}

			if err := sub.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}

}
