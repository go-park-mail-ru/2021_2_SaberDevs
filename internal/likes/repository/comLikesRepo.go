package repository

import (
	"context"

	amodels "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/likes/models"
	sbErr "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/syberErrors"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type ComLikesRepository struct {
	Db *sqlx.DB
}

func NewComLikesRepository(db *sqlx.DB) amodels.LikesRepository {
	return &ComLikesRepository{db}
}

func (m *ComLikesRepository) UpdateCount(ctx context.Context, articlesid int) (int, error) {
	var Likes int
	count := "Select sum(signum) as s from comments_likes WHERE commentId = $1"
	err := m.Db.Get(&Likes, count, articlesid)
	if err != nil {
		Likes = 0
	}
	updateArticle := `UPDATE comments SET Likes = $1 WHERE comments.Id = $2;`

	_, err = m.Db.Exec(updateArticle, Likes, articlesid)
	if err != nil {
		return 0, sbErr.ErrDbError{
			Reason:   err.Error(),
			Function: "/insert",
		}
	}
	return Likes, nil
}

func (m *ComLikesRepository) Insert(ctx context.Context, a *amodels.LikeDb) error {
	ins := `INSERT INTO comments_likes(login, commentId, signum) VALUES ($1, $2, $3) ON CONFLICT DO NOTHING;`
	_, err := m.Db.Exec(ins, a.Login, a.ArticleId, a.Signum)
	if err != nil {
		return sbErr.ErrDbError{
			Reason:   err.Error(),
			Function: "/insert",
		}
	}
	return nil
}

func (m *ComLikesRepository) Delete(ctx context.Context, a *amodels.LikeDb) error {
	delete := `delete from comments_likes  WHERE commentId = $1 and login = $2;`
	_, err := m.Db.Exec(delete, a.ArticleId, a.Login)
	if err != nil {
		return sbErr.ErrDbError{
			Reason:   err.Error(),
			Function: "/delete",
		}
	}
	return nil
}

func (m *ComLikesRepository) Check(ctx context.Context, a *amodels.LikeDb) (int, error) {
	check := `select signum from comments_likes  WHERE commentId = $1 and login = $2;`
	var sign []int
	err := m.Db.Select(&sign, check, a.ArticleId, a.Login)
	if err != nil {
		return 0, sbErr.ErrDbError{
			Reason:   err.Error(),
			Function: "/update",
		}
	}
	if len(sign) == 0 {
		return 0, nil
	}
	return sign[0], nil
}

func (m *ComLikesRepository) Cancel(ctx context.Context, a *amodels.LikeDb) (int, error) {
	sign, err := m.Check(ctx, a)
	if err != nil || sign == 0 {
		return 0, sbErr.ErrBadImage{
			Reason:   err.Error(),
			Function: "/cancel",
		}
	}
	err = m.Delete(ctx, a)
	if err != nil {
		return 0, sbErr.ErrBadImage{
			Reason:   err.Error(),
			Function: "cancel",
		}
	}
	likes, err := m.UpdateCount(ctx, a.ArticleId)
	if err != nil {
		return 0, sbErr.ErrBadImage{
			Reason:   err.Error(),
			Function: "cancel",
		}
	}
	return likes, nil
}

func (m *ComLikesRepository) InsertLike(ctx context.Context, a *amodels.LikeDb) (int, error) {
	sign, err := m.Check(ctx, a)
	if err != nil || sign == a.Signum {
		return 0, sbErr.ErrNoContent{
			Reason:   err.Error(),
			Function: "inslike",
		}
	}

	if sign != 0 && sign != a.Signum {
		err = m.Delete(ctx, a)
		if err != nil {
			return 0, sbErr.ErrBadImage{
				Reason:   err.Error(),
				Function: "cancel",
			}
		}
	}

	err = m.Insert(ctx, a)
	if err != nil {
		return 0, sbErr.ErrBadImage{
			Reason:   err.Error(),
			Function: "inslike",
		}
	}

	likes, err := m.UpdateCount(ctx, a.ArticleId)

	return likes, err
}

func (m *ComLikesRepository) Like(ctx context.Context, a *amodels.LikeDb) (int, error) {
	likes, err := m.InsertLike(ctx, a)
	return likes, err
}

func (m *ComLikesRepository) Dislike(ctx context.Context, a *amodels.LikeDb) (int, error) {
	likes, err := m.InsertLike(ctx, a)
	return likes, err
}
