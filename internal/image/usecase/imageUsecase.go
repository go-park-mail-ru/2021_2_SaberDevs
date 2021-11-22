package usecase

import (
	"context"
	imodels "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/image/models"
	sbErr "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/syberErrors"
	"github.com/pkg/errors"
	"image"
	"image/jpeg"
	"image/png"
	"io"
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

	buff := make([]byte, 512)
	_, err = src.Read(buff)
	if err != nil {
		return imodels.SaveImageResponse{}, sbErr.ErrInternal{
			Reason:   err.Error(),
			Function: "imageUsecase/SaveImage",
		}
	}

	filetype := http.DetectContentType(buff)
	if filetype != "image/jpeg" && filetype != "image/png" {
		return imodels.SaveImageResponse{}, sbErr.ErrBadImage{
			Reason:   err.Error(),
			Function: "imageUsecase/SaveImage",
		}
	}

	_, err = src.Seek(0, io.SeekStart)
	if err != nil {
		return imodels.SaveImageResponse{}, sbErr.ErrInternal{
			Reason:   err.Error(),
			Function: "imageUsecase/SaveImage",
		}
	}

	var img image.Image
	switch filetype {
	case "image/jpeg":
		img, err = jpeg.Decode(src)
		if err != nil {
			return imodels.SaveImageResponse{}, sbErr.ErrInternal{
				Reason:   err.Error(),
				Function: "imageUsecase/SaveImage",
			}
		}
	case "image/png":
		img, err = png.Decode(src)
		if err != nil {
			return imodels.SaveImageResponse{}, sbErr.ErrInternal{
				Reason:   err.Error(),
				Function: "imageUsecase/SaveImage",
			}
		}
	default:
		return imodels.SaveImageResponse{}, sbErr.ErrInternal{
			Reason:   "switch error in default",
			Function: "imageUsecase/SaveImage",
		}
	}

	savedImageName, err := iu.imageRepo.SaveImage(ctx, &img)
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
