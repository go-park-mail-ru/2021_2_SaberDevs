package commentStream

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

const (
	writeWait      = 10 * time.Second
	pongWait       = 60 * time.Second
	pingPeriod     = (pongWait * 9) / 10
	maxMessageSize = 512
)

var (
	newline = []byte{'\n'}
	space   = []byte{' '}
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type Subscriber struct {
	pub  *Publisher
	conn *websocket.Conn
	// канал для получения последних коментариев от publisher
	send chan []string // todo поменять на коменты
}

func (sub *Subscriber) readWS() {
	defer func() {
		sub.pub.unregister <- sub
		sub.conn.Close()
	}()

	sub.conn.SetReadLimit(maxMessageSize)
	sub.conn.SetReadDeadline(time.Now().Add(pongWait))
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


func serveWs(pub *Publisher, w http.ResponseWriter, r *http.Request) error {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return err
	}
	client := &Subscriber{
		pub: pub,
		conn: conn,
		send: make(chan []string),
	}
	client.hub.register <- client

	// Allow collection of memory referenced by the caller by doing all work in
	// new goroutines.
	go client.writePump()
	go client.readPump()
}
