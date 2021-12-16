package article

import (
	"context"
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
	sqlxmock "github.com/zhashkevych/go-sqlxmock"
)

func TestGetByID(t *testing.T) {
	db, mock, err := sqlxmock.Newx()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	rows := sqlxmock.NewRows([]string{"id", "previewurl", "datetime", "title", "category", "text", "authorname", "commentsurl", "comments", "likes"}).
		AddRow(1, "static/img/computer.png", "3", "7 Skills of Highly Effective Programmers", "SaberDevs",
			"Our team was inspired by the seven skills of highly effective", "mollenTEST1", "static/img/photo-elon-musk.jpg",
			97, 1001)

	query := "SELECT Id, PreviewUrl, DateTime, Title, Category, Text, AuthorName,  CommentsUrl, Comments, Likes FROM ARTICLES WHERE articles.Id = $1;"

	mock.ExpectQuery(regexp.QuoteMeta(query)).WithArgs(1).WillReturnRows(rows)

	rows2 := sqlxmock.NewRows([]string{"id", "login", "name", "surname", "avatarurl", "description", "email", "password", "score"}).
		AddRow(1, "mollenTEST1", "mollenTEST1", "7 Skills of Highly Effective Programmers", "SaberDevs",
			"Our team was inspired by the seven skills of highly effective", "mollenTEST1", "static/img/photo-elon-musk.jpg",
			97)

	query2 := "SELECT AU.ID, AU.LOGIN, AU.NAME, AU.SURNAME, AU.AVATARURL, AU.DESCRIPTION, AU.EMAIL, AU.PASSWORD, AU.SCORE FROM ARTICLES AS AR INNER JOIN AUTHOR AS AU ON AU.LOGIN = AR.AuthorName WHERE AR.ID = $1;"

	mock.ExpectQuery(regexp.QuoteMeta(query2)).WithArgs(1).WillReturnRows(rows2)
	rows0 := sqlxmock.NewRows([]string{"tag"}).
		AddRow("tag")

	query0 := "select c.tag from tags c inner join tags_articles ca on c.Id = ca.tags_id inner join articles a on a.Id = ca.articles_id where a.Id = $1;"

	mock.ExpectQuery(regexp.QuoteMeta(query0)).WithArgs(1).WillReturnRows(rows0)

	rows3 := sqlxmock.NewRows([]string{"signum"}).
		AddRow(1)

	query3 := "select signum from article_likes where articleId = $1 and Login = $2;"

	mock.ExpectQuery(regexp.QuoteMeta(query3)).WithArgs(1, "mollen").WillReturnRows(rows3)

	a := NewArticleRepository(db)

	anArticle, err := a.GetByID(context.TODO(), "mollen", 1)
	assert.NoError(t, err)
	assert.NotNil(t, anArticle)
}
