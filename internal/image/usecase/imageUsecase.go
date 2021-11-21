package usecase

import (
	"context"
	imodels "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/image/models"
	sbErr "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/syberErrors"
	"github.com/pkg/errors"
	"mime/multipart"
	"net/http"
)

type imageUsecase struct {
	imageRepo imodels.ImageRepository
}

const MaxFileSize = 1024 * 1024 * 4

func NewImageUsecase(ir imodels.ImageRepository) imodels.ImageUsecase {
	return &imageUsecase{
		imageRepo: ir,
	}
}

func (iu *imageUsecase) GetImage(ctx context.Context, imageName string) (string, error) {
	name, err := iu.imageRepo.GetImageByName(ctx, imageName)
	if err != nil {
		return "", errors.Wrap(err, "imageUsecase/GetImage")
	}

	return name, nil
}

func (iu *imageUsecase) SaveImage(ctx context.Context, file *multipart.FileHeader) (imodels.SaveImageResponse, error) {
	if file.Size > MaxFileSize {
		return imodels.SaveImageResponse{}, sbErr.ErrBadImage{
			Reason:   "image too big",
			Function: "imageUsecase/SaveImage",
		}
	}

	src, err := file.Open()
	if err != nil {
		return imodels.SaveImageResponse{}, sbErr.ErrInternal{
			Reason:   err.Error(),
			Function: "imageUsecase/SaveImage",
		}
	}
	defer src.Close()
	// todo проверить на тип изображения, конвертировать и обрезать до 140*140

	savedImageName, err := iu.imageRepo.SaveImage(ctx, &src)
	if err != nil {
		return imodels.SaveImageResponse{}, errors.Wrap(err, "imageHandler/SaveImage")
	}

	data := imodels.SaveImageData{
		Name: savedImageName,
	}
	response := imodels.SaveImageResponse{
		Status: http.StatusOK,
		Data:   data,
		Msg:    "OK",
	}

	return response, nil
}
