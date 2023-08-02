package repository

import (
	"context"
	wrapper "github.com/go-park-mail-ru/2021_2_SaberDevs/internal"
	cmodels "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/comment/models"
	"github.com/stretchr/testify/assert"
	sqlxmock "github.com/zhashkevych/go-sqlxmock"
	"regexp"
	"testing"
)

func TestStoreComment(t *testing.T) {
	log := wrapper.NewLogger()
	db, mock, err := sqlxmock.Newx()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	crepo := NewCommentRepository(db, log)
	comment := cmodels.Comment{
		Id:          0,
		DateTime:    "",
		Text:        "",
		AuthorLogin: "",
		ArticleId:   0,
		ParentId:    1,
		IsEdited:    false,
		Likes:       0,
	}

	rows := sqlxmock.NewRows([]string{"id"}).
		AddRow(1)

	query := `INSERT INTO comments (AuthorLogin, ArticleId, Likes, ParentId, Text, IsEdited, DateTime) values ($1, $2, $3, $4, $5, $6, $7) returning id;`

	mock.ExpectQuery(regexp.QuoteMeta(query)).WithArgs(comment.AuthorLogin, comment.ArticleId, comment.Likes, comment.ParentId, comment.Text, comment.IsEdited, comment.DateTime).WillReturnRows(rows)

	comm, err := crepo.StoreComment(context.TODO(), &comment)
	assert.NoError(t, err)
	assert.NotNil(t, comm)
}

func TestUpdateComment(t *testing.T) {
	log := wrapper.NewLogger()
	db, mock, err := sqlxmock.Newx()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	crepo := NewCommentRepository(db, log)
	comment := cmodels.Comment{
		Id:          0,
		DateTime:    "",
		Text:        "",
		AuthorLogin: "",
		ArticleId:   0,
		ParentId:    1,
		IsEdited:    false,
		Likes:       0,
	}

	rows := sqlxmock.NewRows([]string{"id", "authorlogin", "articleid", "likes", "parentid", "text", "isedited", "datetime"}).
		AddRow(1, "", 1, 1, 1, "", true, "")

	query := `UPDATE comments SET text = $1, isedited = $2 WHERE id = $3 returning Id, AuthorLogin, ArticleId, Likes, ParentId, Text, IsEdited, DateTime`


	mock.ExpectQuery(regexp.QuoteMeta(query)).WithArgs(comment.Text, comment.IsEdited, comment.Id).WillReturnRows(rows)

	comm, err := crepo.UpdateComment(context.TODO(), &comment)
	assert.NoError(t, err)
	assert.NotNil(t, comm)
}

func TestGetCommentsByArticleID(t *testing.T) {
	log := wrapper.NewLogger()
	db, mock, err := sqlxmock.Newx()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	crepo := NewCommentRepository(db, log)

	rows := sqlxmock.NewRows([]string{"id", "articleid", "likes", "text", "parentid", "isedited", "datetime", "login", "surname", "name", "score", "avatarurl"}).
		AddRow(1, 1, 1, "1", 1, true, "true", "", "", "", 1, "")

	query := `select c.id, c.articleid, c.Likes, c.parentid, c.text, c.isedited, c.datetime, a.login, a.surname, a.name, a.score, a.avatarurl  
               from comments c join author a on a.login = c.AuthorLogin where c.ArticleId = $1 limit 50`


	mock.ExpectQuery(regexp.QuoteMeta(query)).WithArgs(1).WillReturnRows(rows)

	comm, err := crepo.GetCommentsByArticleID(context.TODO(), 1, 0)
	assert.NoError(t, err)
	assert.NotNil(t, comm)
}

func TestGetCommentByID(t *testing.T) {
	log := wrapper.NewLogger()
	db, mock, err := sqlxmock.Newx()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	crepo := NewCommentRepository(db, log)

	rows := sqlxmock.NewRows([]string{"id", "authorlogin", "articleid", "likes", "parentid", "text", "isedited", "datetime"}).
		AddRow(1, "", 1, 1, 1, "", true, "")

	query := `SELECT Id, AuthorLogin, ArticleId, Likes, ParentId, Text, IsEdited, DateTime FROM comments WHERE id = $1`


	mock.ExpectQuery(regexp.QuoteMeta(query)).WithArgs(1).WillReturnRows(rows)

	comm, err := crepo.GetCommentByID(context.TODO(), 1)
	assert.NoError(t, err)
	assert.NotNil(t, comm)
}

func TestGetCommentsStream(t *testing.T) {
	log := wrapper.NewLogger()
	db, mock, err := sqlxmock.Newx()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	crepo := NewCommentRepository(db, log)

	rows := sqlxmock.NewRows([]string{"id", "articleid", "text", "likes", "login", "surname", "name", "avatarurl"}).
		AddRow(1, 1, "1", 1, "1", "", "true", "")

	query := `select c.id, c.articleid, c.text,  c.likes, a.login, a.surname, a.name, a.avatarurl, a2.title
               from comments c join author a on a.login = c.AuthorLogin join articles a2 on c.articleid = a2.id where c.id > $1 order by c.id desc limit 5`


	mock.ExpectQuery(regexp.QuoteMeta(query)).WithArgs(1).WillReturnRows(rows)

	comm, err := crepo.GetCommentsStream(1)
	assert.NoError(t, err)
	assert.NotNil(t, comm)
}
