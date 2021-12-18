package testing

import (
	"context"
	"testing"

	app "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/article/article_app"
	"github.com/go-park-mail-ru/2021_2_SaberDevs/internal/article/models"
	mocks "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/article/models/mock"
	ser "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/article/server/serve"
	gomock "github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestGetByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockUcase := mocks.NewMockArticleUsecase(ctrl)
	mockArticle := models.FullArticle{
		Id:          "1",
		PreviewUrl:  "#",
		Tags:        []string{"fishing", "boat"},
		Title:       "fishing",
		Text:        "fishing is bad for business",
		Author:      models.Author{Id: 1},
		CommentsUrl: "#",
		Comments:    11,
		Likes:       12,
	}
	resArticle := &app.FullArticle{
		Id:          "1",
		PreviewUrl:  "#",
		Tags:        []string{"fishing", "boat"},
		Title:       "fishing",
		Text:        "fishing is bad for business",
		Author:      &app.Author{Id: 1},
		CommentsUrl: "#",
		Comments:    11,
		Likes:       12,
	}
	u := ser.NewArticleManager(mockUcase)
	login := "Iam"
	id := &app.Id{Id: "1", Value: login}
	t.Run("success", func(t *testing.T) {
		mockUcase.EXPECT().GetByID(gomock.Eq(context.TODO()), gomock.Eq(login), gomock.Eq(int64(1))).Return(mockArticle, nil).AnyTimes()
		a, err := u.GetByID(context.TODO(), id)
		assert.NoError(t, err)
		assert.Equal(t, a, resArticle)

	})
}

func TestFetch(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockUcase := mocks.NewMockArticleUsecase(ctrl)
	mockArticle := models.Preview{
		Id:          "1",
		PreviewUrl:  "#",
		Tags:        []string{"fishing", "boat"},
		Title:       "fishing",
		Text:        "fishing is bad for business",
		Author:      models.Author{Id: 1},
		CommentsUrl: "#",
		Comments:    11,
		Likes:       12,
	}
	preArticle := []models.Preview{mockArticle}
	resArticle := &app.Preview{
		Id:          "1",
		PreviewUrl:  "#",
		Tags:        []string{"fishing", "boat"},
		Title:       "fishing",
		Text:        "fishing is bad for business",
		Author:      &app.Author{Id: 1},
		CommentsUrl: "#",
		Comments:    11,
		Likes:       12,
	}
	repArticle := &app.Repview{Preview: []*app.Preview{resArticle}}
	login := "Iam"
	ch := &app.Chunk{IdLastLoaded: "999999", ChunkSize: 5, Value: login}
	u := ser.NewArticleManager(mockUcase)
	t.Run("success", func(t *testing.T) {
		mockUcase.EXPECT().Fetch(gomock.Eq(context.TODO()), gomock.Eq(login), "999999", 5).Return(preArticle, nil).AnyTimes()
		a, err := u.Fetch(context.TODO(), ch)
		assert.NoError(t, err)
		assert.Equal(t, a, repArticle)

	})
}

func TestFindArticles(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockUcase := mocks.NewMockArticleUsecase(ctrl)
	mockArticle := models.Preview{
		Id:          "1",
		PreviewUrl:  "#",
		Tags:        []string{"fishing", "boat"},
		Title:       "fishing",
		Text:        "fishing is bad for business",
		Author:      models.Author{Id: 1},
		CommentsUrl: "#",
		Comments:    11,
		Likes:       12,
	}

	preArticle := []models.Preview{mockArticle}
	resArticle := &app.Preview{
		Id:          "1",
		PreviewUrl:  "#",
		Tags:        []string{"fishing", "boat"},
		Title:       "fishing",
		Text:        "fishing is bad for business",
		Author:      &app.Author{Id: 1},
		CommentsUrl: "#",
		Comments:    11,
		Likes:       12,
	}
	login := "Iam"
	ch := &app.Chunk{
		IdLastLoaded: "999999",
		ChunkSize:    5,
		Value:        login,
	}
	q := &app.Queries{
		Query: "query",
		Chunk: ch,
	}
	query := "query"
	repArticle := &app.Repview{Preview: []*app.Preview{resArticle}}

	u := ser.NewArticleManager(mockUcase)
	t.Run("success", func(t *testing.T) {
		mockUcase.EXPECT().FindArticles(gomock.Eq(context.TODO()), login, query, "999999", 5).Return(preArticle, nil).AnyTimes()
		a, err := u.FindArticles(context.TODO(), q)
		assert.NoError(t, err)
		assert.Equal(t, a, repArticle)

	})
}

func TestFindbyTag(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockUcase := mocks.NewMockArticleUsecase(ctrl)
	mockArticle := models.Preview{
		Id:          "1",
		PreviewUrl:  "#",
		Tags:        []string{"fishing", "boat"},
		Title:       "fishing",
		Text:        "fishing is bad for business",
		Author:      models.Author{Id: 1},
		CommentsUrl: "#",
		Comments:    11,
		Likes:       12,
	}

	preArticle := []models.Preview{mockArticle}
	resArticle := &app.Preview{
		Id:          "1",
		PreviewUrl:  "#",
		Tags:        []string{"fishing", "boat"},
		Title:       "fishing",
		Text:        "fishing is bad for business",
		Author:      &app.Author{Id: 1},
		CommentsUrl: "#",
		Comments:    11,
		Likes:       12,
	}
	login := "Iam"
	ch := &app.Chunk{
		IdLastLoaded: "999999",
		ChunkSize:    5,
		Value:        login,
	}
	q := &app.Queries{
		Query: "query",
		Chunk: ch,
	}
	query := "query"
	repArticle := &app.Repview{Preview: []*app.Preview{resArticle}}

	u := ser.NewArticleManager(mockUcase)
	t.Run("success", func(t *testing.T) {
		mockUcase.EXPECT().FindByTag(gomock.Eq(context.TODO()), login, query, "999999", 5).Return(preArticle, nil).AnyTimes()
		a, err := u.FindByTag(context.TODO(), q)
		assert.NoError(t, err)
		assert.Equal(t, a, repArticle)

	})
}

func TestGetbyAuthor(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockUcase := mocks.NewMockArticleUsecase(ctrl)
	mockArticle := models.Preview{
		Id:          "1",
		PreviewUrl:  "#",
		Tags:        []string{"fishing", "boat"},
		Title:       "fishing",
		Text:        "fishing is bad for business",
		Author:      models.Author{Id: 1},
		CommentsUrl: "#",
		Comments:    11,
		Likes:       12,
	}

	preArticle := []models.Preview{mockArticle}
	resArticle := &app.Preview{
		Id:          "1",
		PreviewUrl:  "#",
		Tags:        []string{"fishing", "boat"},
		Title:       "fishing",
		Text:        "fishing is bad for business",
		Author:      &app.Author{Id: 1},
		CommentsUrl: "#",
		Comments:    11,
		Likes:       12,
	}
	login := "Iam"
	ch := &app.Chunk{
		IdLastLoaded: "999999",
		ChunkSize:    5,
		Value:        login,
	}
	author := "author"
	au := &app.Authors{
		Author: author,
		Chunk:  ch,
	}
	repArticle := &app.Repview{Preview: []*app.Preview{resArticle}}

	u := ser.NewArticleManager(mockUcase)
	t.Run("success", func(t *testing.T) {
		mockUcase.EXPECT().GetByAuthor(gomock.Eq(context.TODO()), login, author, "999999", 5).Return(preArticle, nil).AnyTimes()
		a, err := u.GetByAuthor(context.TODO(), au)
		assert.NoError(t, err)
		assert.Equal(t, a, repArticle)

	})
}

func TestGetbyCategory(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockUcase := mocks.NewMockArticleUsecase(ctrl)
	mockArticle := models.Preview{
		Id:          "1",
		PreviewUrl:  "#",
		Tags:        []string{"fishing", "boat"},
		Title:       "fishing",
		Text:        "fishing is bad for business",
		Author:      models.Author{Id: 1},
		CommentsUrl: "#",
		Comments:    11,
		Likes:       12,
	}

	preArticle := []models.Preview{mockArticle}
	resArticle := &app.Preview{
		Id:          "1",
		PreviewUrl:  "#",
		Tags:        []string{"fishing", "boat"},
		Title:       "fishing",
		Text:        "fishing is bad for business",
		Author:      &app.Author{Id: 1},
		CommentsUrl: "#",
		Comments:    11,
		Likes:       12,
	}
	login := "Iam"
	ch := &app.Chunk{
		IdLastLoaded: "999999",
		ChunkSize:    5,
		Value:        login,
	}
	category := "category"
	cat := &app.Categories{
		Category: category,
		Chunk:    ch,
	}
	repArticle := &app.Repview{Preview: []*app.Preview{resArticle}}

	u := ser.NewArticleManager(mockUcase)
	t.Run("success", func(t *testing.T) {
		mockUcase.EXPECT().GetByCategory(gomock.Eq(context.TODO()), login, category, "999999", 5).Return(preArticle, nil).AnyTimes()
		a, err := u.GetByCategory(context.TODO(), cat)
		assert.NoError(t, err)
		assert.Equal(t, a, repArticle)

	})
}

func TestGetByTag(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockUcase := mocks.NewMockArticleUsecase(ctrl)
	mockArticle := models.Preview{
		Id:          "1",
		PreviewUrl:  "#",
		Tags:        []string{"fishing", "boat"},
		Title:       "fishing",
		Text:        "fishing is bad for business",
		Author:      models.Author{Id: 1},
		CommentsUrl: "#",
		Comments:    11,
		Likes:       12,
	}

	preArticle := []models.Preview{mockArticle}
	resArticle := &app.Preview{
		Id:          "1",
		PreviewUrl:  "#",
		Tags:        []string{"fishing", "boat"},
		Title:       "fishing",
		Text:        "fishing is bad for business",
		Author:      &app.Author{Id: 1},
		CommentsUrl: "#",
		Comments:    11,
		Likes:       12,
	}
	login := "Iam"
	ch := &app.Chunk{
		IdLastLoaded: "999999",
		ChunkSize:    5,
		Value:        login,
	}
	tag := "tag"
	tags := &app.Tags{
		Tag:   tag,
		Chunk: ch,
	}
	repArticle := &app.Repview{Preview: []*app.Preview{resArticle}}

	u := ser.NewArticleManager(mockUcase)
	t.Run("success", func(t *testing.T) {
		mockUcase.EXPECT().GetByTag(gomock.Eq(context.TODO()), login, tag, "999999", 5).Return(preArticle, nil).AnyTimes()
		a, err := u.GetByTag(context.TODO(), tags)
		assert.NoError(t, err)
		assert.Equal(t, a, repArticle)

	})
}

func TestFindAuthors(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockUcase := mocks.NewMockArticleUsecase(ctrl)

	authors := []models.Author{models.Author{Id: 12}}

	login := "Iam"
	ch := &app.Chunk{
		IdLastLoaded: "999999",
		ChunkSize:    5,
		Value:        login,
	}
	query := "query"
	q := &app.Queries{
		Query: query,
		Chunk: ch,
	}
	Auth := &app.Author{Id: 12}
	repArticle := &app.AView{Author: []*app.Author{Auth}}

	u := ser.NewArticleManager(mockUcase)
	t.Run("success", func(t *testing.T) {
		mockUcase.EXPECT().FindAuthors(gomock.Eq(context.TODO()), query, "999999", 5).Return(authors, nil).AnyTimes()
		a, err := u.FindAuthors(context.TODO(), q)
		assert.NoError(t, err)
		assert.Equal(t, a, repArticle)

	})
}
