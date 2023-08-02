package likes

import (
	"context"
	"fmt"

	amodels "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/likes/models"
	smodels "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/session/models"
	sbErr "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/syberErrors"
	"github.com/pkg/errors"
)

type comLikeUCase struct {
	rep         amodels.LikesRepository
	sessionRepo smodels.SessionRepository
}

func NewComLikeUsecase(rep amodels.LikesRepository, sessionRepo smodels.SessionRepository) amodels.LikesUsecase {
	return &comLikeUCase{rep, sessionRepo}
}

func (m *comLikeUCase) Rating(ctx context.Context, a *amodels.LikeData, cValue string) (int, error) {
	login, err := m.sessionRepo.GetSessionLogin(ctx, cValue)
	if err != nil {
		return 0, errors.Wrap(err, "articleUsecase/Delete")
	}
	like := amodels.LikeDb{}
	like.ArticleId = a.Id
	like.Login = login
	like.Signum = a.Sign
	Id := -3
	if like.Signum == 1 {
		Id, err = m.rep.Like(ctx, &like)
	}
	if like.Signum == 0 {
		Id, err = m.rep.Cancel(ctx, &like)
	}
	if like.Signum == -1 {
		Id, err = m.rep.Dislike(ctx, &like)
	}
	if Id == -3 {
		fmt.Println("ID =", Id)
		return Id, sbErr.ErrNotFeedNumber{Reason: "ID =" + fmt.Sprint(Id)}
	}
	return Id, err
}
