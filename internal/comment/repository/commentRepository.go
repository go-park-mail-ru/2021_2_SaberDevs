package repository

import (
	"context"
	"database/sql"
	cmodels "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/comment/models"
	sbErr "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/syberErrors"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type commentPsqlRepo struct {
	Db *sqlx.DB
}

type sqlComment struct {
	Id          int64         `json:"Id"  db:"id"`
	DateTime    string        `json:"datetime" db:"datetime"`
	Text        string        `json:"text" db:"text"`
	AuthorLogin string        `json:"authorLogin" db:"authorlogin"`
	ArticleId   int64         `json:"articleId" db:"articleid"`
	ParentId    sql.NullInt64 `json:"parentId" db:"parentid"`
	IsEdited    bool          `json:"isEdited" db:"isedited"`
}

type sqlPreparedComment struct {
	Id        int64          `json:"Id"  db:"id"`
	DateTime  string         `json:"datetime" db:"datetime"`
	Text      string         `json:"text" db:"text"`
	ArticleId int64          `json:"articleIdd" db:"articleid"`
	ParentId  sql.NullInt64          `json:"parentId" db:"parentid"`
	IsEdited  bool           `json:"isEdited" db:"isedited"`
	cmodels.Author `json:"author"`
}

func NewCommentRepository(db *sqlx.DB) cmodels.CommentRepository {
	return &commentPsqlRepo{db}
}

func (cr *commentPsqlRepo) StoreComment(ctx context.Context, comment *cmodels.Comment) (cmodels.Comment, error) {
	schema := `INSERT INTO comments (AuthorLogin, ArticleId, ParentId, Text, IsEdited, DateTime) values ($1, $2, $3, $4, $5, $6) returning id;`
	var result *sql.Rows

	if comment.ParentId == 0 {
		var err error
		result, err = cr.Db.Query(schema, comment.AuthorLogin, comment.ArticleId, sql.NullInt64{}, comment.Text, comment.IsEdited, comment.DateTime)
		if err != nil {
			return cmodels.Comment{}, sbErr.ErrInternal{
				Reason:   err.Error(),
				Function: "commentRepository/StoreComment",
			}
		}
	} else {
		var err error
		result, err = cr.Db.Query(schema, comment.AuthorLogin, comment.ArticleId, comment.ParentId, comment.Text, comment.IsEdited, comment.DateTime)
		if err != nil {
			return cmodels.Comment{}, sbErr.ErrInternal{
				Reason:   err.Error(),
				Function: "commentRepository/StoreComment",
			}
		}
	}

	var commentID int64
	for result.Next() {
		err := result.Scan(&commentID)
		if err != nil {
			return cmodels.Comment{}, sbErr.ErrInternal{
				Reason:   err.Error(),
				Function: "commentRepository/StoreComment",
			}
		}
	}

	comment.Id = commentID

	return *comment, nil
}

func (cr *commentPsqlRepo) UpdateComment(ctx context.Context, comment *cmodels.Comment) (cmodels.Comment, error) {
	// id, err := strconv.Atoi(comment.Id)
	// if err != nil {
	// 	return cmodels.Comment{}, sbErr.ErrInternal{
	// 		Reason:   err.Error(),
	// 		Function: "commentRepository/StoreComment",
	// 	}
	// }

	result, err := cr.Db.Query(`UPDATE comments SET text = $1, isedited = $2 WHERE id = $3 returning Id, AuthorLogin, ArticleId, ParentId, Text, IsEdited, DateTime`,
		comment.Text, comment.IsEdited, comment.Id)
	if err != nil {
		return cmodels.Comment{}, sbErr.ErrInternal{
			Reason:   err.Error(),
			Function: "commentRepository/StoreComment",
		}
	}

	var editedComment sqlComment
	for result.Next() {
		err = result.Scan(&editedComment)
		if err != nil {
			return cmodels.Comment{}, sbErr.ErrInternal{
				Reason:   err.Error(),
				Function: "commentRepository/StoreComment",
			}
		}
	}

	return cmodels.Comment{
		Id:          editedComment.Id,
		DateTime:    editedComment.DateTime,
		Text:        editedComment.Text,
		AuthorLogin: editedComment.AuthorLogin,
		ArticleId:   editedComment.ArticleId,
		ParentId:    editedComment.ParentId.Int64,
		IsEdited:    editedComment.IsEdited,
	}, nil
}

// Login     string `json:"login" db:"login"`
// 	Surname   string `json:"lastName" db:"surname"`
// 	Name      string `json:"firstName" db:"name"`
// 	Score     int    `json:"score" db:"score"`
// 	AvatarURL string `json:"avatarUrl" db:"avatarurl"`

func (cr *commentPsqlRepo) GetCommentsByArticleID(ctx context.Context, articleID int64, lastCommentID int64) ([]cmodels.PreparedComment, error) {
	var comments []sqlPreparedComment

	schema := `select c.id, c.articleid, c.parentid, c.text, c.isedited, c.datetime, a.login, a.surname, a.name, a.score, a.avatarurl  
               from comments c join author a on a.login = c.AuthorLogin where c.ArticleId = $1 limit 50`

	err := cr.Db.Select(&comments, schema, articleID)
	if err != nil {
		return []cmodels.PreparedComment{}, sbErr.ErrInternal{
			Reason:   err.Error(),
			Function: "commentRepository/GetCommentsByArticleID",
		}
	}
	return []cmodels.PreparedComment{}, nil
}

func (cr *commentPsqlRepo) GetCommentByID(ctx context.Context, commentID int64) (cmodels.Comment, error) {
	var comment sqlComment

	err := cr.Db.Get(&comment, `SELECT Id, AuthorLogin, ArticleId, ParentId, Text, IsEdited, DateTime FROM comments WHERE id = $1`, commentID)
	if err != nil {
		return cmodels.Comment{}, sbErr.ErrInternal{
			Reason:   err.Error(),
			Function: "commentRepository/StoreComment",
		}
	}

	return cmodels.Comment{
		Id:          comment.Id,
		DateTime:    comment.DateTime,
		Text:        comment.Text,
		AuthorLogin: comment.AuthorLogin,
		ArticleId:   comment.ArticleId,
		ParentId:    comment.ParentId.Int64,
		IsEdited:    comment.IsEdited,
	}, nil
}
