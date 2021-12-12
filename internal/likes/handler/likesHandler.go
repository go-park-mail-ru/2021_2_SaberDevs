package handlers

import (
	"fmt"
	"net/http"

	amodels "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/likes/models"
	uCase "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/likes/usecase"
	sbErr "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/syberErrors"
	echo "github.com/labstack/echo"
	"github.com/pkg/errors"
)

type LikesHandler struct {
	arUseCase  amodels.LikesUsecase
	comUseCase amodels.LikesUsecase
}

func NewLikesHandler() *LikesHandler {
	return &LikesHandler{uCase.NewArLikeUsecase(), uCase.NewComLikeUsecase()}
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
	ctx := c.context.context()
	Id, err := api.arUseCase.Rating(ctx, like, cVal)
	if err != nil {
		return errors.Wrap(err, "likesHandler/Rate")
	}

	response := amodels.GenericResponse{
		Status: http.StatusOK,
		Data:   fmt.Sprint(Id),
	}

	return c.JSON(http.StatusOK, response)
}
