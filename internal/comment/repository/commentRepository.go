package repository

import (
	"context"
	cmodels "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/comment/models"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type commentPsqlRepo struct {
	Db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) cmodels.CommentRepository {
	return &commentPsqlRepo{db}
}

func (cr *commentPsqlRepo) StoreComment(ctx context.Context, comment *cmodels.Comment) (cmodels.Comment, error) {

}

func (cr *commentPsqlRepo) UpdateComment(ctx context.Context, comment *cmodels.Comment) (cmodels.Comment, error) {

}

func (cr *commentPsqlRepo) GetCommentsByArticleID(ctx context.Context, articleID string, lastCommentID string) ([]cmodels.PreparedComment, error) {

}

func (cr *commentPsqlRepo) GetCommentByID(ctx context.Context, commentID string) (cmodels.Comment, error) {

}
