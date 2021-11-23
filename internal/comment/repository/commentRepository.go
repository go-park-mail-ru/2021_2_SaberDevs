package repository

import (
	"context"
	"database/sql"
	cmodels "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/comment/models"
	sbErr "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/syberErrors"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"strconv"
)

type commentPsqlRepo struct {
	Db *sqlx.DB
}

func NewCommentRepository(db *sqlx.DB) cmodels.CommentRepository {
	return &commentPsqlRepo{db}
}

func (cr *commentPsqlRepo) StoreComment(ctx context.Context, comment *cmodels.Comment) (cmodels.Comment, error) {
	schema := `INSERT INTO comments (AuthorLogin, ArticleId, ParentId, Text, IsEdited, DateTime) values ($1, $2, $3, $4, $5, $6) returning id;`
	articleID, err := strconv.Atoi(comment.ArticleId)
	var result *sql.Rows

	if comment.ParentId == "" {
		result, err = cr.Db.Query(schema, comment.AuthorLogin, articleID, sql.NullInt64{}, comment.Text, comment.IsEdited, comment.DateTime)
		if err != nil {
			return cmodels.Comment{}, sbErr.ErrInternal{
				Reason:   err.Error(),
				Function: "commentRepository/StoreComment",
			}
		}
	} else {
		parentID, err := strconv.Atoi(comment.ParentId)
		result, err = cr.Db.Query(schema, comment.AuthorLogin, articleID, parentID, comment.Text, comment.IsEdited, comment.DateTime)
		if err != nil {
			return cmodels.Comment{}, sbErr.ErrInternal{
				Reason:   err.Error(),
				Function: "commentRepository/StoreComment",
			}
		}
	}

	var commentID int
	for result.Next() {
		err = result.Scan(&commentID)
		if err != nil {
			return cmodels.Comment{}, sbErr.ErrInternal{
				Reason:   err.Error(),
				Function: "commentRepository/StoreComment",
			}
		}
	}

	comment.Id = strconv.Itoa(commentID)

	return *comment, nil
}

func (cr *commentPsqlRepo) UpdateComment(ctx context.Context, comment *cmodels.Comment) (cmodels.Comment, error) {
	id, err := strconv.Atoi(comment.Id)
	if err != nil {
		return cmodels.Comment{}, sbErr.ErrInternal{
			Reason:   err.Error(),
			Function: "commentRepository/StoreComment",
		}
	}

	result, err := cr.Db.Query(`UPDATE comments SET text = $1, isedited = $2 WHERE id = $3 returning Id, AuthorLogin, ArticleId, ParentId, Text, IsEdited, DateTime`,
		comment.Text, comment.IsEdited, id)
	if err != nil {
		return cmodels.Comment{}, sbErr.ErrInternal{
			Reason:   err.Error(),
			Function: "commentRepository/StoreComment",
		}
	}

	var editedComment cmodels.Comment
	for result.Next() {
		err = result.Scan(&editedComment)
		if err != nil {
			return cmodels.Comment{}, sbErr.ErrInternal{
				Reason:   err.Error(),
				Function: "commentRepository/StoreComment",
			}
		}
	}

	return cmodels.Comment{}, nil
}

func (cr *commentPsqlRepo) GetCommentsByArticleID(ctx context.Context, articleID string, lastCommentID string) ([]cmodels.PreparedComment, error) {
	return []cmodels.PreparedComment{}, nil
}

func (cr *commentPsqlRepo) GetCommentByID(ctx context.Context, commentID string) (cmodels.Comment, error) {
	return cmodels.Comment{}, nil
}
