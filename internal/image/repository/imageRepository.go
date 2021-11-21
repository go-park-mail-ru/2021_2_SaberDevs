package repository

import (
	"context"
	imodels "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/image/models"
	sbErr "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/syberErrors"
	uuid "github.com/satori/go.uuid"
	"io"
	"mime/multipart"
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

func (ir *imageRepo) SaveImage(ctx context.Context, src *multipart.File) (string, error) {
	imgName := uuid.NewV4().String()
	imgFile, err := os.Create("static/" + imgName)
	if err != nil {
		return "", sbErr.ErrInternal{
			Reason:   err.Error(),
			Function: "imageRepository/SaveImage",
		}
	}
	defer imgFile.Close()

	if _, err = io.Copy(imgFile, *src); err != nil {
		return "", sbErr.ErrInternal{
			Reason:   err.Error(),
			Function: "imageRepository/SaveImage",
		}
	}

	return imgName, nil
}
