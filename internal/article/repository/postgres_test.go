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
	rows := sqlxmock.NewRows([]string{"id", "previewurl", "title", "text", "authorurl", "authorname", "authoravatar", "commentsurl", "comments", "likes"}).
		AddRow(1, "static/img/computer.png", "7 Skills of Highly Effective Programmers",
			"Our team was inspired by the seven skills of highly effective", "#", "mollenTEST1", "static/img/photo-elon-musk.jpg",
			"#", 97, 1001)

	query := `SELECT * FROM ARTICLES WHERE articles.Id = $1`

	articleID := int64(1)
	mock.ExpectQuery(regexp.QuoteMeta(query)).WithArgs(articleID).WillReturnRows(rows)

	rowsNew := sqlxmock.NewRows([]string{"tag"}).AddRow("design").AddRow("finance")

	queryNew := `select c.tag from categories c
	inner join categories_articles ca  on c.Id = ca.categories_id
	inner join articles a on a.Id = ca.articles_id
	where a.Id = $1;`

	articleIDString := 1
	mock.ExpectQuery(regexp.QuoteMeta(queryNew)).WithArgs(articleIDString).WillReturnRows(rowsNew)

	a := NewArticleRepository(db)

	anArticle, err := a.GetByID(context.TODO(), 1)
	assert.NoError(t, err)
	assert.NotNil(t, anArticle)
}
