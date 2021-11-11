package models

import "context"

type Image struct {
	Name string
}

type ImageUsecase interface {
	GetImage(ctx context.Context, imageName string) (string, error)
}

type ImageRepository interface {
	GetImageByName(ctx context.Context, imageName string) (string, error)
}
