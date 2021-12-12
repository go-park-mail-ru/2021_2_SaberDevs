package likes

import (
	"context"

	amodels "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/likes/models"
	smodels "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/session/models"
)

type arLikeUCase struct {
	rep         amodels.LikesRepository
	sessionRepo smodels.SessionRepository
}

func NewArLikeUsecase(rep amodels.LikesRepository, sessionRepo smodels.SessionRepository) amodels.LikesUsecase {
	return &arLikeUCase{rep, sessionRepo}
}

func (m *arLikeUCase) Rating(ctx context.Context, a *amodels.LikeData, cValue string) (int, error) {
	return 0, nil
}
