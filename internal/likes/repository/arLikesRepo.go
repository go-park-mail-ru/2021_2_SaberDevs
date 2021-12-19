package repository

import (
	"context"
	"fmt"

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

func (m *ArLikesRepository) UpdateCount(ctx context.Context, articlesid int) (int, error) {
	var Likes int
	count := "Select sum(signum) as s from article_likes WHERE articleId = $1"
	err := m.Db.Get(&Likes, count, articlesid)
	if err != nil {
		Likes = 0
	}
	updateArticle := `UPDATE articles SET Likes = $1 WHERE articles.Id = $2;`

	_, err = m.Db.Exec(updateArticle, Likes, articlesid)
	if err != nil {
		return 0, sbErr.ErrDbError{
			Reason:   err.Error(),
			Function: "/insert",
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
	check := `select signum from article_likes  WHERE articleId = $1 and login = $2;`
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

func (m *ArLikesRepository) Cancel(ctx context.Context, a *amodels.LikeDb) (int, error) {
	sign, err := m.Check(ctx, a)
	if err != nil || sign == 0 {
		return 0, sbErr.ErrBadImage{
			Reason:   err.Error(),
			Function: "/cancel",
		}
	}
	fmt.Println("signum = ", sign, "id =", a.ArticleId)
	err = m.Delete(ctx, a)
	if err != nil {
		fmt.Println("err = ", err.Error())
		return 0, sbErr.ErrBadImage{
			Reason:   err.Error(),
			Function: "cancel",
		}
	}
	fmt.Println("deleted = ", sign)
	likes, err := m.UpdateCount(ctx, a.ArticleId)
	if err != nil {
		fmt.Println("err = ", err.Error())
		return 0, sbErr.ErrBadImage{
			Reason:   err.Error(),
			Function: "cancel",
		}
	}
	return likes, nil
}

func (m *ArLikesRepository) InsertLike(ctx context.Context, a *amodels.LikeDb) (int, error) {
	sign, err := m.Check(ctx, a)
	if err != nil || sign == a.Signum {
		return 0, sbErr.ErrNoContent{
			Reason:   fmt.Sprint(sign),
			Function: "inslike",
		}
	}
	var likes int

	if sign != 0 && sign != a.Signum {
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

	likes, err = m.UpdateCount(ctx, a.ArticleId)
	if err != nil {
		return 0, sbErr.ErrBadImage{
			Reason:   err.Error(),
			Function: "inslike",
		}
	}
	fmt.Println("REPO LIKES =", likes)
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
