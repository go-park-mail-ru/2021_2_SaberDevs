package repository

// import (
// 	"context"
// 	wrapper "github.com/go-park-mail-ru/2021_2_SaberDevs/internal"
// 	cmodels "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/comment/models"
// 	"github.com/stretchr/testify/assert"
// 	sqlxmock "github.com/zhashkevych/go-sqlxmock"
// 	"regexp"
// 	"testing"
// )
//
// func TestStoreComment(t *testing.T) {
// 	log := wrapper.NewLogger()
// 	db, mock, err := sqlxmock.Newx()
// 	if err != nil {
// 		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
// 	}
//
// 	crepo := NewCommentRepository(db, log)
// 	comment := cmodels.Comment{
// 		Id:          0,
// 		DateTime:    "",
// 		Text:        "",
// 		AuthorLogin: "",
// 		ArticleId:   0,
// 		ParentId:    0,
// 		IsEdited:    false,
// 		Likes:       0,
// 	}
//
// 	rows := sqlxmock.NewRows([]string{"id"}).
// 		AddRow(1)
//
// 	query := `INSERT INTO comments (AuthorLogin, ArticleId, Likes, ParentId, Text, IsEdited, DateTime) values ($1, $2, $3, $4, $5, $6, $7) returning id;`
//
// 	mock.ExpectQuery(regexp.QuoteMeta(query)).WithArgs(comment.AuthorLogin, comment.ArticleId, comment.Likes, comment.ParentId, comment.Text, comment.IsEdited, comment.DateTime).WillReturnRows(rows)
//
// 	comm, err := crepo.StoreComment(context.TODO(), &comment)
// 	assert.NoError(t, err)
// 	assert.NotNil(t, comm)
// }