package article

import (
	"context"
	"regexp"
	"testing"

	amodels "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/likes/models"
	repo "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/likes/repository"
	"github.com/stretchr/testify/assert"
	sqlxmock "github.com/zhashkevych/go-sqlxmock"
)

func TestStore(t *testing.T) {
	db, mock, err := sqlxmock.Newx()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	login := "mollenTEST1"

	art := amodels.LikeDb{Login: login, ArticleId: 1, Signum: 1}

	rows := sqlxmock.NewRows([]string{"signum"}).
		AddRow(1)

	query := "select signum from article_likes  WHERE articleId = $1 and login = $2;"
	mock.ExpectQuery(regexp.QuoteMeta(query)).WithArgs(art.ArticleId, art.Login).WillReturnRows(rows)

	a := repo.NewArLikesRepository(db)
	aid, err := a.Like(context.TODO(), &art)
	assert.NoError(t, err)
	assert.NotNil(t, aid)
	assert.Equal(t, aid, 1)
}
