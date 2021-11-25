package commentStream

import (
	cmodels "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/comment/models"
	sbErr "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/syberErrors"
	"github.com/labstack/echo/v4"
	"strings"
)

const firstComment = 0

type commentStreamHandler struct {
	pub         *Publisher
	commentRepo cmodels.CommentRepository
}

func NewCommentStreamHandler(p *Publisher, cr cmodels.CommentRepository) *commentStreamHandler {
	return &commentStreamHandler{
		pub:         p,
		commentRepo: cr,
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

	var lastComment int64 = 0
	comments, err := api.commentRepo.GetCommentsStream(lastComment)
	if len(comments) != 0 {
		lastComment = comments[0].Id

		for _, comment := range comments {
			articleNameSlice := strings.Split(comment.ArticleName, "")[:25]
			err = conn.WriteJSON(streamComment{
				Type:        "stream-comment",
				Id:          comment.Id,
				Text:        comment.Text,
				ArticleId:   comment.ArticleId,
				ArticleName: strings.Join(articleNameSlice, ""),
				author: author{
					Login:     comment.Login,
					Surname:   comment.Surname,
					Name:      comment.Name,
					AvatarURL: comment.AvatarURL,
				},
			})
		}
	}
	sub := &Subscriber{
		pub:  api.pub,
		conn: conn,
		send: make(chan []cmodels.StreamComment),
	}
	sub.pub.register <- sub

	go sub.writeWS(lastComment)
	go sub.readWS()
	return nil
}
