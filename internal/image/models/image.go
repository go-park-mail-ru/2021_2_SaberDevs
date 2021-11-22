package models

import (
	"context"
	"image"
	"mime/multipart"
)

type Image struct {
	Name string
}

type SaveImageData struct {
	Name string `json:"imgId"`
}

type SaveImageResponse struct {
	Status uint          `json:"status"`
	Data   SaveImageData `json:"data"`
	Msg    string        `json:"msg"`
}

type ImageUsecase interface {
	GetImage(ctx context.Context, imageName string) (string, error)
	SaveImage(ctx context.Context, file *multipart.FileHeader) (SaveImageResponse, error)
}

type ImageRepository interface {
	GetImageByName(ctx context.Context, imageName string) (string, error)
	SaveImage(ctx context.Context, src *image.Image) (string, error)
}
