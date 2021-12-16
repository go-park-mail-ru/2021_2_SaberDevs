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
	rows := sqlxmock.NewRows([]string{"id", "previewurl", "title", "category", "text", "authorname", "commentsurl", "comments", "likes"}).
		AddRow(1, "static/img/computer.png", "7 Skills of Highly Effective Programmers", "SaberDevs",
			"Our team was inspired by the seven skills of highly effective", "mollenTEST1", "static/img/photo-elon-musk.jpg",
			97, 1001)

	queryNew := `"SELECT Id, PreviewUrl, DateTime, Title, Category, Text, AuthorName,  CommentsUrl, Comments, Likes FROM ARTICLES WHERE articles.Id = $1"`

	mock.ExpectQuery(regexp.QuoteMeta(queryNew)).WithArgs(context.TODO(), "mollen", 1).WillReturnRows(rows)

	a := NewArticleRepository(db)

	anArticle, err := a.GetByID(context.TODO(), "mollen", 1)
	assert.NoError(t, err)
	assert.NotNil(t, anArticle)
}
