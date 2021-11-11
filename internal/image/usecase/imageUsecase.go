package usecase

import (
	"context"
	imodels "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/image/models"
	"github.com/pkg/errors"
)

type imageUsecase struct {
	imageRepo imodels.ImageRepository
}

func NewImageUsecase(ir imodels.ImageRepository) imodels.ImageUsecase {
	return &imageUsecase{
		imageRepo: ir,
	}
}

func (iu *imageUsecase)GetImage(ctx context.Context, imageName string) (string, error) {
	name , err := iu.imageRepo.GetImageByName(ctx, imageName)
	if err != nil {
		return "", errors.Wrap(err, "imageUsecase/GetImage")
	}

	return name, nil
}
