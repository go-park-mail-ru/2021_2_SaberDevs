package handler

import (
	imodels "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/image/models"
	sbErr "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/syberErrors"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
)

type ImageHandler struct {
	ImageUsecase imodels.ImageUsecase
}

func NewImageHandler(iu imodels.ImageUsecase) *ImageHandler {
	return &ImageHandler{iu}
}

func (api *ImageHandler) GetImage(c echo.Context) error {
	imgName := c.Param("name")
	if imgName == "" {
		return sbErr.ErrNoContent{
			Reason:  "no image",
			Function: "imageHandler/GetImage",
		}
	}

	ctx := c.Request().Context()
	name , err := api.ImageUsecase.GetImage(ctx, imgName)
	if err != nil {
		return errors.Wrap(err, "imageHandler/GetImage")
	}

	return c.File("static/" + name)
}
