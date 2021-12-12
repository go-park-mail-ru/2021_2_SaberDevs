package likes

import (
	"context"

	amodels "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/likes/models"
)

type comLikeUCase struct {
}

func NewComLikeUsecase() amodels.LikesUsecase {
	return &arLikeUCase{}
}

func (m *comLikeUCase) Rating(ctx context.Context, a *amodels.LikeData, cValue string) (int, error) {
	return 0, nil
}
