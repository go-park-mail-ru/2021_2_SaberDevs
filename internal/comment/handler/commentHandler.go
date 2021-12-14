package handler

import (
	"net/http"
	"strconv"

	cmodels "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/comment/models"
	sbErr "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/syberErrors"
	"github.com/labstack/echo/v4"
	"github.com/microcosm-cc/bluemonday"
	"github.com/pkg/errors"
)

type CommentHandler struct {
	ComentUsecase cmodels.CommentUsecase
}

func NewCommentHandler(cu cmodels.CommentUsecase) *CommentHandler {
	return &CommentHandler{cu}
}

func SanitizeComment(c cmodels.Comment) cmodels.Comment {
	s := bluemonday.StrictPolicy()
	c.Text = s.Sanitize(c.Text)
	return c
}

func (api *CommentHandler) CreateComment(c echo.Context) error {
	requestComment := cmodels.Comment{}
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
	response, err := api.ComentUsecase.CreateComment(ctx, &requestComment, cookie.Value)
	if err != nil {
		return errors.Wrap(err, "commentHandler/CreateComment")
	}

	return c.JSON(http.StatusOK, response)
}

func (api *CommentHandler) UpdateComment(c echo.Context) error {
	requestComment := cmodels.Comment{}
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

	ctx := c.Request().Context()
	response, err := api.ComentUsecase.UpdateComment(ctx, &requestComment, cookie.Value)
	// response, err := api.ComentUsecase.UpdateComment(ctx, &requestComment, "cookie.Value")
	if err != nil {
		return errors.Wrap(err, "commentHandler/UpdateComment")
	}

	return c.JSON(http.StatusOK, response)
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

	response, err := api.ComentUsecase.GetCommentsByArticleID(ctx, id)
	if err != nil {
		return errors.Wrap(err, "commentHandler/GetCommentsByArticleID")
	}

	return c.JSON(http.StatusOK, response)
}
