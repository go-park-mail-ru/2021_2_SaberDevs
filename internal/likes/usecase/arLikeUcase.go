package likes

import (
	"context"

	amodels "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/likes/models"
)

type arLikeUCase struct {
}

func NewArLikeUsecase() amodels.LikesUsecase {
	return &arLikeUCase{}
}

func (m *arLikeUCase) Rating(ctx context.Context, a *amodels.LikeData, cValue string) (int, error) {
	return 0, nil
}
