package commentStream

import (
	cmodels "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/comment/models"
	"github.com/gorilla/websocket"
	"net/http"
	"strings"
	"time"
)

const (
	writeWait      = 10 * time.Second
	pongWait       = 11 * time.Second
	pingPeriod     = (pongWait * 8) / 10
	maxMessageSize = 512
)

type streamComment struct {
	Type string `json:"type"`
	Id          int64  `json:"Id"  db:"id"`
	Text        string `json:"text" db:"text"`
	ArticleId   int64  `json:"articleId" db:"articleid"`
	ArticleName string `json:"articleName" db:"title"`
	author      `json:"author"`
}

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
	send chan []cmodels.StreamComment // todo поменять на коменты
}

func (sub *Subscriber) readWS() {
	defer func() {
		sub.pub.unregister <- sub
		sub.conn.Close()
	}()

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
			
			for _, comment := range message {
				articleNameSlice := strings.Split(comment.ArticleName, "")[:25]
				err = sub.conn.WriteJSON(streamComment{
					Type:        "stream-comment",
					Id:          comment.Id,
					Text:        comment.Text,
					ArticleId:   comment.ArticleId,
					ArticleName: strings.Join(articleNameSlice, ""),
					author:      author{
						Login:     comment.Login,
						Surname:   comment.Surname,
						Name:      comment.Name,
						AvatarURL: comment.AvatarURL,
					},
				})
				if err != nil {
					return
				}
			}

			// err = sub.conn.WriteJSON(message)
			// if err != nil {
			// 	return
			// }

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
