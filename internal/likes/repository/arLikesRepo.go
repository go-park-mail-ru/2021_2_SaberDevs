package repository

import (
	"context"

	amodels "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/likes/models"
	sbErr "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/syberErrors"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type ArLikesRepository struct {
	Db *sqlx.DB
}

func NewArLikesRepository(db *sqlx.DB) amodels.LikesRepository {
	return &ArLikesRepository{db}
}

func (m *ArLikesRepository) UpdateCount(ctx context.Context, articlesid int, change int) (int, error) {
	updateArticle := `UPDATE articles SET Likes = Likes + $1  WHERE articles.Id = $2 RETURNING Likes;`
	var Likes int
	err := m.Db.Get(&Likes, updateArticle, change, articlesid)
	if err != nil {
		return 0, sbErr.ErrDbError{
			Reason:   err.Error(),
			Function: "/update",
		}
	}

	return Likes, nil
}

func (m *ArLikesRepository) Insert(ctx context.Context, a *amodels.LikeDb) error {
	ins := `INSERT INTO article_likes(login, articleId, signum) VALUES ($1, $2, $3) ON CONFLICT DO NOTHING;`
	_, err := m.Db.Exec(ins, a.Login, a.ArticleId, a.Signum)
	if err != nil {
		return sbErr.ErrDbError{
			Reason:   err.Error(),
			Function: "/insert",
		}
	}
	return nil
}

func (m *ArLikesRepository) Delete(ctx context.Context, a *amodels.LikeDb) error {
	delete := `delete from article_likes  WHERE articleId = $1 and login = $2;`
	_, err := m.Db.Exec(delete, a.ArticleId, a.Login)
	if err != nil {
		return sbErr.ErrDbError{
			Reason:   err.Error(),
			Function: "/delete",
		}
	}
	return nil
}

func (m *ArLikesRepository) Check(ctx context.Context, a *amodels.LikeDb) (int, error) {
	sign := -3
	check := `select signum from article_likes  WHERE articleId = $1 and login = $2;`
	err := m.Db.Get(&sign, check, a.ArticleId, a.Login)
	if err != nil {
		return 0, sbErr.ErrBadImage{
			Reason:   err.Error(),
			Function: "/check",
		}
	}
	return sign, nil
}

func (m *ArLikesRepository) Cancel(ctx context.Context, a *amodels.LikeDb) (int, error) {
	sign, err := m.Check(ctx, a)
	if err != nil || sign == -3 {
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
	likes, err := m.UpdateCount(ctx, a.ArticleId, -sign)
	if err != nil {
		return 0, sbErr.ErrBadImage{
			Reason:   err.Error(),
			Function: "cancel",
		}
	}
	return likes, nil
}

func (m *ArLikesRepository) InsertLike(ctx context.Context, a *amodels.LikeDb) (int, error) {
	sign, err := m.Check(ctx, a)
	if sign != a.Signum && err == nil {
		err = m.Delete(ctx, a)
		if err != nil {
			return 0, sbErr.ErrBadImage{
				Reason:   err.Error(),
				Function: "inslike",
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
	likes, err := m.UpdateCount(ctx, a.ArticleId, sign)
	if err != nil {
		return 0, sbErr.ErrBadImage{
			Reason:   err.Error(),
			Function: "inslike",
		}
	}
	return likes, nil
}

func (m *ArLikesRepository) Like(ctx context.Context, a *amodels.LikeDb) (int, error) {
	likes, err := m.InsertLike(ctx, a)
	return likes, err
}

func (m *ArLikesRepository) Dislike(ctx context.Context, a *amodels.LikeDb) (int, error) {
	likes, err := m.InsertLike(ctx, a)
	return likes, err
}
