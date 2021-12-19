package article

import (
	"context"
	"database/sql/driver"
	"regexp"
	"testing"

	amodels "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/likes/models"
	repo "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/likes/repository"
	"github.com/stretchr/testify/assert"
	sqlxmock "github.com/zhashkevych/go-sqlxmock"
)

func TestArLike(t *testing.T) {
	db, mock, err := sqlxmock.Newx()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	login := "mollenTEST1"

	art := amodels.LikeDb{Login: login, ArticleId: 1, Signum: 1}

	rows := sqlxmock.NewRows([]string{"signum"}).
		AddRow(-1)

	query := "select signum from article_likes  WHERE articleId = $1 and login = $2;"
	mock.ExpectQuery(regexp.QuoteMeta(query)).WithArgs(art.ArticleId, art.Login).WillReturnRows(rows)

	query5 := "delete from article_likes WHERE articleId = $1 and login = $2;"
	mock.ExpectExec(regexp.QuoteMeta(query5)).WithArgs(1, art.Login).WillReturnResult(driver.RowsAffected(1))

	query2 := "INSERT INTO article_likes(login, articleId, signum) VALUES ($1, $2, $3) ON CONFLICT DO NOTHING;"
	mock.ExpectExec(regexp.QuoteMeta(query2)).WithArgs(art.Login, art.ArticleId, art.Signum).WillReturnResult(driver.RowsAffected(1))

	query3 := "UPDATE articles SET Likes = $1 WHERE articles.Id = $2;"
	mock.ExpectExec(regexp.QuoteMeta(query3)).WithArgs(0, art.ArticleId).WillReturnResult(driver.RowsAffected(1))

	rows4 := sqlxmock.NewRows([]string{"s"}).
		AddRow(0)
	query4 := "UPDATE articles SET Likes = $1 WHERE articles.Id = $2;"
	mock.ExpectQuery(regexp.QuoteMeta(query4)).WithArgs(0, art.ArticleId).WillReturnRows(rows4)

	a := repo.NewArLikesRepository(db)
	aid, err := a.Like(context.TODO(), &art)
	assert.NoError(t, err)
	assert.NotNil(t, aid)
	assert.Equal(t, aid, 0)
}

func TestArDisLike(t *testing.T) {
	db, mock, err := sqlxmock.Newx()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	login := "mollenTEST1"

	art := amodels.LikeDb{Login: login, ArticleId: 1, Signum: -1}

	rows := sqlxmock.NewRows([]string{"signum"}).
		AddRow(1)

	query := "select signum from article_likes  WHERE articleId = $1 and login = $2;"
	mock.ExpectQuery(regexp.QuoteMeta(query)).WithArgs(art.ArticleId, art.Login).WillReturnRows(rows)

	query5 := "delete from article_likes WHERE articleId = $1 and login = $2;"
	mock.ExpectExec(regexp.QuoteMeta(query5)).WithArgs(1, art.Login).WillReturnResult(driver.RowsAffected(1))

	query2 := "INSERT INTO article_likes(login, articleId, signum) VALUES ($1, $2, $3) ON CONFLICT DO NOTHING;"
	mock.ExpectExec(regexp.QuoteMeta(query2)).WithArgs(art.Login, art.ArticleId, art.Signum).WillReturnResult(driver.RowsAffected(1))

	query3 := "UPDATE articles SET Likes = $1 WHERE articles.Id = $2;"
	mock.ExpectExec(regexp.QuoteMeta(query3)).WithArgs(0, art.ArticleId).WillReturnResult(driver.RowsAffected(1))

	rows4 := sqlxmock.NewRows([]string{"s"}).
		AddRow(0)
	query4 := "UPDATE articles SET Likes = $1 WHERE articles.Id = $2;"
	mock.ExpectQuery(regexp.QuoteMeta(query4)).WithArgs(0, art.ArticleId).WillReturnRows(rows4)

	a := repo.NewArLikesRepository(db)
	aid, err := a.Dislike(context.TODO(), &art)
	assert.NoError(t, err)
	assert.NotNil(t, aid)
	assert.Equal(t, aid, 0)
}

func TestArCancel(t *testing.T) {
	db, mock, err := sqlxmock.Newx()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	login := "mollenTEST1"

	art := amodels.LikeDb{Login: login, ArticleId: 1, Signum: 0}

	rows := sqlxmock.NewRows([]string{"signum"}).
		AddRow(1)

	query := "select signum from article_likes  WHERE articleId = $1 and login = $2;"
	mock.ExpectQuery(regexp.QuoteMeta(query)).WithArgs(art.ArticleId, art.Login).WillReturnRows(rows)

	query5 := "delete from article_likes WHERE articleId = $1 and login = $2;"
	mock.ExpectExec(regexp.QuoteMeta(query5)).WithArgs(1, art.Login).WillReturnResult(driver.RowsAffected(1))

	query3 := "UPDATE articles SET Likes = $1 WHERE articles.Id = $2;"
	mock.ExpectExec(regexp.QuoteMeta(query3)).WithArgs(0, art.ArticleId).WillReturnResult(driver.RowsAffected(1))

	a := repo.NewArLikesRepository(db)
	aid, err := a.Cancel(context.TODO(), &art)
	assert.NoError(t, err)
	assert.NotNil(t, aid)
	assert.Equal(t, aid, 0)
}

func TestComLike(t *testing.T) {
	db, mock, err := sqlxmock.Newx()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	login := "mollenTEST1"

	art := amodels.LikeDb{Login: login, ArticleId: 1, Signum: 1}

	rows := sqlxmock.NewRows([]string{"signum"}).
		AddRow(-1)

	query := "select signum from comments_likes  WHERE commentId = $1 and login = $2;"
	mock.ExpectQuery(regexp.QuoteMeta(query)).WithArgs(art.ArticleId, art.Login).WillReturnRows(rows)

	query5 := "delete from comments_likes WHERE commentId = $1 and login = $2;"
	mock.ExpectExec(regexp.QuoteMeta(query5)).WithArgs(1, art.Login).WillReturnResult(driver.RowsAffected(1))

	query2 := "INSERT INTO comments_likes(login, commentId, signum) VALUES ($1, $2, $3) ON CONFLICT DO NOTHING;"
	mock.ExpectExec(regexp.QuoteMeta(query2)).WithArgs(art.Login, art.ArticleId, art.Signum).WillReturnResult(driver.RowsAffected(1))

	query3 := "UPDATE comments SET Likes = $1 WHERE comments.Id = $2;"
	mock.ExpectExec(regexp.QuoteMeta(query3)).WithArgs(0, art.ArticleId).WillReturnResult(driver.RowsAffected(1))

	rows4 := sqlxmock.NewRows([]string{"s"}).
		AddRow(0)
	query4 := "UPDATE comments SET Likes = $1 WHERE comments.Id = $2;"
	mock.ExpectQuery(regexp.QuoteMeta(query4)).WithArgs(0, art.ArticleId).WillReturnRows(rows4)

	a := repo.NewComLikesRepository(db)
	aid, err := a.Like(context.TODO(), &art)
	assert.NoError(t, err)
	assert.NotNil(t, aid)
	assert.Equal(t, aid, 0)
}

func TestComDisLike(t *testing.T) {
	db, mock, err := sqlxmock.Newx()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	login := "mollenTEST1"

	art := amodels.LikeDb{Login: login, ArticleId: 1, Signum: -1}

	rows := sqlxmock.NewRows([]string{"signum"}).
		AddRow(1)

	query := "select signum from comments_likes  WHERE commentId = $1 and login = $2;"
	mock.ExpectQuery(regexp.QuoteMeta(query)).WithArgs(art.ArticleId, art.Login).WillReturnRows(rows)

	query5 := "delete from comments_likes WHERE commentId = $1 and login = $2;"
	mock.ExpectExec(regexp.QuoteMeta(query5)).WithArgs(1, art.Login).WillReturnResult(driver.RowsAffected(1))

	query2 := "INSERT INTO comments_likes(login, commentId, signum) VALUES ($1, $2, $3) ON CONFLICT DO NOTHING;"
	mock.ExpectExec(regexp.QuoteMeta(query2)).WithArgs(art.Login, art.ArticleId, art.Signum).WillReturnResult(driver.RowsAffected(1))

	query3 := "UPDATE comments SET Likes = $1 WHERE comments.Id = $2;"
	mock.ExpectExec(regexp.QuoteMeta(query3)).WithArgs(0, art.ArticleId).WillReturnResult(driver.RowsAffected(1))

	rows4 := sqlxmock.NewRows([]string{"s"}).
		AddRow(0)
	query4 := "UPDATE comments SET Likes = $1 WHERE comments.Id = $2;"
	mock.ExpectQuery(regexp.QuoteMeta(query4)).WithArgs(0, art.ArticleId).WillReturnRows(rows4)

	a := repo.NewComLikesRepository(db)
	aid, err := a.Dislike(context.TODO(), &art)
	assert.NoError(t, err)
	assert.NotNil(t, aid)
	assert.Equal(t, aid, 0)
}

func TestComCancel(t *testing.T) {
	db, mock, err := sqlxmock.Newx()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	login := "mollenTEST1"

	art := amodels.LikeDb{Login: login, ArticleId: 1, Signum: 0}

	rows := sqlxmock.NewRows([]string{"signum"}).
		AddRow(1)

	query := "select signum from comments_likes  WHERE commentId = $1 and login = $2;"
	mock.ExpectQuery(regexp.QuoteMeta(query)).WithArgs(art.ArticleId, art.Login).WillReturnRows(rows)

	query5 := "delete from comments_likes WHERE commentId = $1 and login = $2;"
	mock.ExpectExec(regexp.QuoteMeta(query5)).WithArgs(1, art.Login).WillReturnResult(driver.RowsAffected(1))

	query3 := "UPDATE comments SET Likes = $1 WHERE comments.Id = $2;"
	mock.ExpectExec(regexp.QuoteMeta(query3)).WithArgs(0, art.ArticleId).WillReturnResult(driver.RowsAffected(1))

	a := repo.NewComLikesRepository(db)
	aid, err := a.Cancel(context.TODO(), &art)
	assert.NoError(t, err)
	assert.NotNil(t, aid)
	assert.Equal(t, aid, 0)
}
