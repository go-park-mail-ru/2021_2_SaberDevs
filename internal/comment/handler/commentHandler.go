package handler

import (
	app "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/comment/comment_app"
	cmodels "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/comment/models"
	sbErr "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/syberErrors"
	"github.com/labstack/echo/v4"
	"github.com/microcosm-cc/bluemonday"
	"github.com/pkg/errors"
	"net/http"
	"strconv"
)

type CommentHandler struct {
	ComentUsecase app.CommentDeliveryClient
}

func NewCommentHandler(cu app.CommentDeliveryClient) *CommentHandler {
	return &CommentHandler{cu}
}

func SanitizeComment(c app.Comment) app.Comment {
	s := bluemonday.StrictPolicy()
	c.Text = s.Sanitize(c.Text)
	return c
}

func (api *CommentHandler) CreateComment(c echo.Context) error {
	requestComment := app.Comment{}
	err := c.Bind(&requestComment)
	if err != nil {
		return sbErr.ErrUnpackingJSON{
			Reason:   err.Error(),
			Function: "commentHandler/CreateComment",
		}
	}

	cookie, err := c.Cookie("session")
	if err != nil {
		return sbErr.ErrNotLoggedin{
			Reason:   err.Error(),
			Function: "commentHandler/CreateComment",
		}
	}

	requestComment = SanitizeComment(requestComment)

	ctx := c.Request().Context()
	createCommentInput := &app.CreateCommentInput{
		Comment:              &requestComment,
		SessionID:            cookie.Value,
	}

	response, err := api.ComentUsecase.CreateComment(ctx, createCommentInput)
	if err != nil {
		return errors.Wrap(err, "commentHandler/CreateComment")
	}

	modelResponse := cmodels.Response{
		Status: uint(response.Status),
		Data:   cmodels.PreparedComment{
			Id:        response.Data.Id,
			DateTime:  response.Data.DateTime,
			Text:      response.Data.Text,
			ArticleId: response.Data.ArticleId,
			ParentId:  response.Data.ParentId,
			IsEdited:  response.Data.IsEdited,
			Likes:     int(response.Data.Likes),
			Author:    cmodels.Author{
				Login:     response.Data.Author.Login,
				Surname:   response.Data.Author.LastName,
				Name:      response.Data.Author.FirstName,
				Score:     int(response.Data.Author.Score),
				AvatarURL: response.Data.Author.AvatarUrl,
			},
		},
		Msg:    response.Msg,
	}

	return c.JSON(http.StatusOK, modelResponse)
}

func (api *CommentHandler) UpdateComment(c echo.Context) error {
	requestComment := app.Comment{}
	err := c.Bind(&requestComment)
	if err != nil {
		return sbErr.ErrUnpackingJSON{
			Reason:   err.Error(),
			Function: "commentHandler/UpdateComment",
		}
	}

	cookie, err := c.Cookie("session")
	if err != nil {
		return sbErr.ErrNotLoggedin{
			Reason:   err.Error(),
			Function: "commentHandler/UpdateComment",
		}
	}

	requestComment = SanitizeComment(requestComment)

	updateCommentInput := &app.UpdateCommentInput{
		Comment:              &requestComment,
		SessionID:            cookie.Value,
	}
	ctx := c.Request().Context()

	response, err := api.ComentUsecase.UpdateComment(ctx, updateCommentInput)
	if err != nil {
		return errors.Wrap(err, "commentHandler/UpdateComment")
	}

	modelResponse := cmodels.Response{
		Status: uint(response.Status),
		Data:   cmodels.PreparedComment{
			Id:        response.Data.Id,
			DateTime:  response.Data.DateTime,
			Text:      response.Data.Text,
			ArticleId: response.Data.ArticleId,
			ParentId:  response.Data.ParentId,
			IsEdited:  response.Data.IsEdited,
			Likes:     int(response.Data.Likes),
			Author:    cmodels.Author{
				Login:     response.Data.Author.Login,
				Surname:   response.Data.Author.LastName,
				Name:      response.Data.Author.FirstName,
				Score:     int(response.Data.Author.Score),
				AvatarURL: response.Data.Author.AvatarUrl,
			},
		},
		Msg:    response.Msg,
	}

	return c.JSON(http.StatusOK, modelResponse)
}

func (api *CommentHandler) GetCommentsByArticleID(c echo.Context) error {
	aricleID := c.QueryParam("id")
	ctx := c.Request().Context()

	id, err := strconv.ParseInt(aricleID, 10, 64)
	if err != nil {
		return sbErr.ErrInternal{
			Reason:   err.Error(),
			Function: "commentHandler/GetCommentsByArticleID",
		}
	}

	response, err := api.ComentUsecase.GetCommentsByArticleID(ctx, &app.ArticleID{ArticleID: id})
	if err != nil {
		return errors.Wrap(err, "commentHandler/GetCommentsByArticleID")
	}

	var prepComment []cmodels.PreparedComment

	for _, c := range response.Data {
		prepComment = append(prepComment, cmodels.PreparedComment{
			Id:        c.Id,
			DateTime:  c.DateTime,
			Text:      c.Text,
			ArticleId: c.ArticleId,
			ParentId:  c.ParentId,
			IsEdited:  c.IsEdited,
			Likes:     int(c.Likes),
			Author:    cmodels.Author{
				Login:     c.Author.Login,
				Surname:   c.Author.LastName,
				Name:      c.Author.FirstName,
				Score:     int(c.Author.Score),
				AvatarURL: c.Author.AvatarUrl,
			},
		})
	}

	modelResponse := cmodels.Response{
		Status: uint(response.Status),
		Data:   prepComment,
		Msg:    response.Msg,
	}

	return c.JSON(http.StatusOK, modelResponse)
}
