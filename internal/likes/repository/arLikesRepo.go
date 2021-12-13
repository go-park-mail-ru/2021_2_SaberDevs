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
	updateArticle := `UPDATE articles SET Likes = Likes + $  WHERE articles.Id  = $2 RETURNING Likes;`
	var Likes int
	err := m.Db.Get(&Likes, updateArticle)
	if err != nil {
		return 0, sbErr.ErrDbError{
			Reason:   err.Error(),
			Function: "articleRepository/Store",
		}
	}

	return Likes, nil
}

func (m *ArLikesRepository) Insert(ctx context.Context, a *amodels.LikeDb) (int, error) {
	ins := `INSERT INTO articles_likes(login, article_id, signum) VALUES
	($1, $2, $3) ON CONFLICT DO NOTHING;`
	_, err := m.Db.Exec(ins, a.Login, a.ArticleId, a.Signum)
	if err != nil {
		return 0, sbErr.ErrDbError{
			Reason:   err.Error(),
			Function: "/insert",
		}
	}
	return 0, nil
}

func (m *ArLikesRepository) Delete(ctx context.Context, a *amodels.LikeDb) error {
	delete := `delete from article_likes  WHERE article_id = $1 and login = $2;`
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
	check := `signum from article_likes  WHERE article_id = $1 and login = $2;`
	err := m.Db.Get(&sign, check, a.ArticleId, a.Login)
	if err != nil {
		return 0, sbErr.ErrBadImage{
			Reason:   err.Error(),
			Function: "articleRepository/Store",
		}
	}
	return sign, nil
}

func (m *ArLikesRepository) Cancel(ctx context.Context, a *amodels.LikeDb) (int, error) {
	sign, err := m.Check(ctx, a)
	if err != nil || sign == -3 {
		return 0, sbErr.ErrBadImage{
			Reason:   err.Error(),
			Function: "articleRepository/Store",
		}
	}
	err = m.Delete(ctx, a)
	if err != nil {
		return 0, sbErr.ErrBadImage{
			Reason:   err.Error(),
			Function: "articleRepository/Store",
		}
	}
	likes, err := m.UpdateCount(ctx, a.ArticleId, -sign)
	if err != nil {
		return 0, sbErr.ErrBadImage{
			Reason:   err.Error(),
			Function: "articleRepository/Store",
		}
	}
	return likes, nil
}

func (m *ArLikesRepository) InsertLike(ctx context.Context, a *amodels.LikeDb) (int, error) {
	sign, err := m.Check(ctx, a)
	if err != nil || sign == -3 {
		return 0, sbErr.ErrBadImage{
			Reason:   err.Error(),
			Function: "articleRepository/Store",
		}
	}
	if sign != a.Signum {
		err = m.Delete(ctx, a)
		if err != nil {
			return 0, sbErr.ErrBadImage{
				Reason:   err.Error(),
				Function: "articleRepository/Store",
			}
		}
		likes, err := m.UpdateCount(ctx, a.ArticleId, -sign)
		if err != nil {
			return 0, sbErr.ErrBadImage{
				Reason:   err.Error(),
				Function: "articleRepository/Store",
			}
		}
		return likes, nil
	}
	return 0, sbErr.ErrBadImage{
		Reason:   err.Error(),
		Function: "articleRepository/Store",
	}
}

func (m *ArLikesRepository) Like(ctx context.Context, a *amodels.LikeDb) (int, error) {
	likes, err := m.InsertLike(ctx, a)
	return likes, err
}

func (m *ArLikesRepository) Dislike(ctx context.Context, a *amodels.LikeDb) (int, error) {
	likes, err := m.InsertLike(ctx, a)
	return likes, err
}
