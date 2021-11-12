package repository

import (
	"context"
	imodels "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/image/models"
)

type imageRepo struct {
}

func NewImageRepository() imodels.ImageRepository {
	return &imageRepo{}
}

func (ir *imageRepo)GetImageByName(ctx context.Context, imageName string) (string, error) {
	return imageName, nil
}
