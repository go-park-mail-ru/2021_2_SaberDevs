package repository

import (
	"context"
	imodels "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/image/models"
	sbErr "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/syberErrors"
	"github.com/kolesa-team/go-webp/encoder"
	"github.com/kolesa-team/go-webp/webp"
	uuid "github.com/satori/go.uuid"
	"image"
	"os"
)

type imageRepo struct {
}

func NewImageRepository() imodels.ImageRepository {
	return &imageRepo{}
}

func (ir *imageRepo) GetImageByName(ctx context.Context, imageName string) (string, error) {
	return imageName, nil
}

func (ir *imageRepo) SaveImage(ctx context.Context, src *image.Image) (string, error) {
	imgName := uuid.NewV4().String()
	imgFile, err := os.Create("static/" + imgName + ".webp")
	if err != nil {
		return "", sbErr.ErrInternal{
			Reason:   err.Error(),
			Function: "imageRepository/SaveImage",
		}
	}
	defer imgFile.Close()

	options, err := encoder.NewLossyEncoderOptions(encoder.PresetDefault, 75)
	if err != nil {
		return "", sbErr.ErrInternal{
			Reason:   err.Error(),
			Function: "imageRepository/SaveImage",
		}
	}

	if err := webp.Encode(imgFile, *src, options); err != nil {
		return "", sbErr.ErrInternal{
			Reason:   err.Error(),
			Function: "imageRepository/SaveImage",
		}
	}

	return imgName + ".webp", nil
}
