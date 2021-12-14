package handlers

import (
	"fmt"
	"net/http"

	amodels "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/likes/models"
	sbErr "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/syberErrors"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
)

type LikesHandler struct {
	arUseCase  amodels.LikesUsecase
	comUseCase amodels.LikesUsecase
}

func NewLikesHandler(arUseCase amodels.LikesUsecase, comUseCase amodels.LikesUsecase) *LikesHandler {
	return &LikesHandler{arUseCase, comUseCase}
}

func (api *LikesHandler) Rate(c echo.Context) error {
	like := new(amodels.LikeData)
	err := c.Bind(like)
	if err != nil {
		return sbErr.ErrUnpackingJSON{
			Reason:   err.Error(),
			Function: "likesHandler/Rate",
		}
	}
	cookie, err := c.Cookie("session")
	if err != nil {
		return sbErr.ErrAuthorised{
			Reason:   err.Error(),
			Function: "likesHandler/Rate",
		}
	}
	cVal := cookie.Value
	ctx := c.Request().Context()
	num := -1
	if like.Ltype == 0 {
		num, err = api.arUseCase.Rating(ctx, like, cVal)
		if err != nil {
			return errors.Wrap(err, "likesHandler/Rate")
		}
	}
	if like.Ltype == 1 {
		num, err = api.comUseCase.Rating(ctx, like, cVal)
		if err != nil {
			return errors.Wrap(err, "likesHandler/Rate")
		}
	}
	if num == -1 {
		return sbErr.ErrNotFeedNumber{
			Reason:   err.Error(),
			Function: "likesHandler/Rate",
		}
	}
	response := amodels.GenericResponse{
		Status: http.StatusOK,
		Data:   fmt.Sprint(num),
	}

	return c.JSON(http.StatusOK, response)
}
