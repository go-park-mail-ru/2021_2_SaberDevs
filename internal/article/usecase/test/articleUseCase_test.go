package testing

import (
	"context"
	"testing"

	amodels "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/article/models"
	mocks "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/article/models/mock"
	repo "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/article/usecase"
	smocks "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/session/models/mock"
	gomock "github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestGetByID(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()
	mockSesRepo := smocks.NewMockSessionRepository(ctrl)
	mockArticleRepo := mocks.NewMockArticleRepository(ctrl)
	mockArticle := amodels.FullArticle{
		Id:          "1",
		PreviewUrl:  "#",
		Tags:        []string{"fishing", "boat"},
		Title:       "fishing",
		Text:        "fishing is bad for business",
		Author:      amodels.Author{},
		CommentsUrl: "#",
		Comments:    11,
		Likes:       12,
	}

	t.Run("success", func(t *testing.T) {
		mockArticleRepo.EXPECT().GetByID(gomock.Eq(context.TODO()), gomock.Eq("author"), gomock.Eq(int64(1))).Return(mockArticle, nil).AnyTimes()
		mockSesRepo.EXPECT().GetSessionLogin(gomock.Eq(context.TODO()), gomock.Eq("author")).Return("author", nil).AnyTimes()
		u := repo.NewArticleUsecase(mockArticleRepo, mockSesRepo)
		a, err := u.GetByID(context.TODO(), "author", 1)
		assert.NoError(t, err)
		assert.Equal(t, a, mockArticle)

	})

}

func TestFetch(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()
	mockSesRepo := smocks.NewMockSessionRepository(ctrl)
	mockArticleRepo := mocks.NewMockArticleRepository(ctrl)
	mockArticle := amodels.Preview{
		Id:          "1",
		PreviewUrl:  "#",
		Tags:        []string{"fishing", "boat"},
		Title:       "fishing",
		Text:        "fishing is bad for business",
		Author:      amodels.Author{},
		CommentsUrl: "#",
		Comments:    11,
		Likes:       12,
	}
	mockArticles := []amodels.Preview{mockArticle}
	login := "Iam"
	maxInt := 999999
	chunkSize := 5
	t.Run("success", func(t *testing.T) {
		mockSesRepo.EXPECT().GetSessionLogin(gomock.Eq(context.TODO()), gomock.Eq(login)).Return(login, nil).AnyTimes()
		mockArticleRepo.EXPECT().Fetch(gomock.Any(), login, maxInt, chunkSize).Return(mockArticles, nil).AnyTimes()
		u := repo.NewArticleUsecase(mockArticleRepo, mockSesRepo)
		a, err := u.Fetch(context.TODO(), login, "", chunkSize)
		assert.NoError(t, err)
		assert.Equal(t, a, mockArticles)
	})

}

func TestGetByTag(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()
	mockSesRepo := smocks.NewMockSessionRepository(ctrl)
	mockArticleRepo := mocks.NewMockArticleRepository(ctrl)
	mockArticle := amodels.Preview{
		Id:          "1",
		PreviewUrl:  "#",
		Tags:        []string{"fishing", "boat"},
		Title:       "fishing",
		Text:        "fishing is bad for business",
		Author:      amodels.Author{},
		CommentsUrl: "#",
		Comments:    11,
		Likes:       12,
	}
	mockArticles := []amodels.Preview{mockArticle}
	login := "Iam"
	tag := "tag"
	maxInt := 999999
	chunkSize := 5
	t.Run("success", func(t *testing.T) {
		mockSesRepo.EXPECT().GetSessionLogin(gomock.Eq(context.TODO()), gomock.Eq(login)).Return(login, nil).AnyTimes()
		mockArticleRepo.EXPECT().GetByTag(gomock.Any(), login, tag, maxInt, chunkSize).Return(mockArticles, nil).AnyTimes()
		u := repo.NewArticleUsecase(mockArticleRepo, mockSesRepo)
		a, err := u.GetByTag(context.TODO(), login, tag, "", chunkSize)
		assert.NoError(t, err)
		assert.Equal(t, a, mockArticles)
	})
}

func TestGetByAuthor(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()
	mockSesRepo := smocks.NewMockSessionRepository(ctrl)
	mockArticleRepo := mocks.NewMockArticleRepository(ctrl)
	mockArticle := amodels.Preview{
		Id:          "1",
		PreviewUrl:  "#",
		Tags:        []string{"fishing", "boat"},
		Title:       "fishing",
		Text:        "fishing is bad for business",
		Author:      amodels.Author{},
		CommentsUrl: "#",
		Comments:    11,
		Likes:       12,
	}
	mockArticles := []amodels.Preview{mockArticle}
	login := "Iam"
	author := "author"
	maxInt := 999999
	chunkSize := 5
	t.Run("success", func(t *testing.T) {
		mockSesRepo.EXPECT().GetSessionLogin(gomock.Eq(context.TODO()), gomock.Eq(login)).Return(login, nil).AnyTimes()
		mockArticleRepo.EXPECT().GetByAuthor(gomock.Any(), login, author, maxInt, chunkSize).Return(mockArticles, nil).AnyTimes()
		u := repo.NewArticleUsecase(mockArticleRepo, mockSesRepo)
		a, err := u.GetByAuthor(context.TODO(), login, author, "", chunkSize)
		assert.NoError(t, err)
		assert.Equal(t, a, mockArticles)
	})
}

func TestGetByCategory(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()
	mockSesRepo := smocks.NewMockSessionRepository(ctrl)
	mockArticleRepo := mocks.NewMockArticleRepository(ctrl)
	mockArticle := amodels.Preview{
		Id:          "1",
		PreviewUrl:  "#",
		Tags:        []string{"fishing", "boat"},
		Title:       "fishing",
		Text:        "fishing is bad for business",
		Author:      amodels.Author{},
		CommentsUrl: "#",
		Comments:    11,
		Likes:       12,
	}
	mockArticles := []amodels.Preview{mockArticle}
	login := "Iam"
	category := "category"
	maxInt := 999999
	chunkSize := 5
	t.Run("success", func(t *testing.T) {
		mockSesRepo.EXPECT().GetSessionLogin(gomock.Eq(context.TODO()), gomock.Eq(login)).Return(login, nil).AnyTimes()
		mockArticleRepo.EXPECT().GetByCategory(gomock.Any(), login, category, maxInt, chunkSize).Return(mockArticles, nil).AnyTimes()
		u := repo.NewArticleUsecase(mockArticleRepo, mockSesRepo)
		a, err := u.GetByCategory(context.TODO(), login, category, "", chunkSize)
		assert.NoError(t, err)
		assert.Equal(t, a, mockArticles)
	})
}

func TestFindByTag(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()
	mockSesRepo := smocks.NewMockSessionRepository(ctrl)
	mockArticleRepo := mocks.NewMockArticleRepository(ctrl)
	mockArticle := amodels.Preview{
		Id:          "1",
		PreviewUrl:  "#",
		Tags:        []string{"fishing", "boat"},
		Title:       "fishing",
		Text:        "fishing is bad for business",
		Author:      amodels.Author{},
		CommentsUrl: "#",
		Comments:    11,
		Likes:       12,
	}
	mockArticles := []amodels.Preview{mockArticle}
	login := "Iam"
	tag := "tag"
	maxInt := 999999
	chunkSize := 5
	t.Run("success", func(t *testing.T) {
		mockSesRepo.EXPECT().GetSessionLogin(gomock.Eq(context.TODO()), gomock.Eq(login)).Return(login, nil).AnyTimes()
		mockArticleRepo.EXPECT().FindByTag(gomock.Any(), login, tag, maxInt, chunkSize).Return(mockArticles, nil).AnyTimes()
		u := repo.NewArticleUsecase(mockArticleRepo, mockSesRepo)
		a, err := u.FindByTag(context.TODO(), login, tag, "", chunkSize)
		assert.NoError(t, err)
		assert.Equal(t, a, mockArticles)
	})
}

func TestFindAuthors(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()
	mockSesRepo := smocks.NewMockSessionRepository(ctrl)
	mockArticleRepo := mocks.NewMockArticleRepository(ctrl)
	mockAuthor := amodels.Author{Id: 1}
	mockAuths := []amodels.Author{mockAuthor}
	login := "Iam"
	query := "auth"
	maxInt := 999999
	chunkSize := 5
	t.Run("success", func(t *testing.T) {
		mockSesRepo.EXPECT().GetSessionLogin(gomock.Eq(context.TODO()), gomock.Eq(login)).Return(login, nil).AnyTimes()
		mockArticleRepo.EXPECT().FindAuthors(gomock.Any(), query, maxInt, chunkSize).Return(mockAuths, nil).AnyTimes()
		u := repo.NewArticleUsecase(mockArticleRepo, mockSesRepo)
		a, err := u.FindAuthors(context.TODO(), query, "", chunkSize)
		assert.NoError(t, err)
		assert.Equal(t, a, mockAuths)
	})
}

func TestFindArticles(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()
	mockSesRepo := smocks.NewMockSessionRepository(ctrl)
	mockArticleRepo := mocks.NewMockArticleRepository(ctrl)
	mockArticle := amodels.Preview{
		Id:          "1",
		PreviewUrl:  "#",
		Tags:        []string{"fishing", "boat"},
		Title:       "fishing",
		Text:        "fishing is bad for business",
		Author:      amodels.Author{},
		CommentsUrl: "#",
		Comments:    11,
		Likes:       12,
	}
	mockArticles := []amodels.Preview{mockArticle}
	login := "Iam"
	query := "tag"
	maxInt := 999999
	chunkSize := 5
	t.Run("success", func(t *testing.T) {
		mockSesRepo.EXPECT().GetSessionLogin(gomock.Eq(context.TODO()), gomock.Eq(login)).Return(login, nil).AnyTimes()
		mockArticleRepo.EXPECT().FindArticles(gomock.Any(), login, query, maxInt, chunkSize).Return(mockArticles, nil).AnyTimes()
		u := repo.NewArticleUsecase(mockArticleRepo, mockSesRepo)
		a, err := u.FindArticles(context.TODO(), login, query, "", chunkSize)
		assert.NoError(t, err)
		assert.Equal(t, a, mockArticles)
	})
}

func TestStore(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()
	mockSesRepo := smocks.NewMockSessionRepository(ctrl)
	mockArticleRepo := mocks.NewMockArticleRepository(ctrl)
	mockArticle := amodels.ArticleCreate{
		Tags:  []string{"fishing", "boat"},
		Title: "fishing",
		Text:  "fishing is bad for business",
	}
	login := "Iam"

	t.Run("success", func(t *testing.T) {
		mockSesRepo.EXPECT().GetSessionLogin(gomock.Eq(context.TODO()), gomock.Eq(login)).Return(login, nil).AnyTimes()
		mockArticleRepo.EXPECT().Store(gomock.Any(), gomock.Any()).Return(1, nil).AnyTimes()
		u := repo.NewArticleUsecase(mockArticleRepo, mockSesRepo)
		a, err := u.Store(context.TODO(), login, &mockArticle)
		assert.NoError(t, err)
		assert.Equal(t, a, 1)
	})
}

func TestUpdate(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()
	mockSesRepo := smocks.NewMockSessionRepository(ctrl)
	mockArticleRepo := mocks.NewMockArticleRepository(ctrl)
	mockArticle := amodels.ArticleUpdate{
		Id:    "1",
		Tags:  []string{"fishing", "boat"},
		Title: "fishing",
		Text:  "fishing is bad for business",
	}
	mockID := amodels.FullArticle{
		Id:          "1",
		PreviewUrl:  "#",
		Tags:        []string{"fishing", "boat"},
		Title:       "fishing",
		Text:        "fishing is bad for business",
		Author:      amodels.Author{},
		CommentsUrl: "#",
		Comments:    11,
		Likes:       12,
	}
	login := "Iam"

	t.Run("success", func(t *testing.T) {
		mockSesRepo.EXPECT().GetSessionLogin(gomock.Eq(context.TODO()), gomock.Eq(login)).Return(login, nil).AnyTimes()
		mockArticleRepo.EXPECT().GetByID(gomock.Eq(context.TODO()), login, gomock.Eq(int64(1))).Return(mockID, nil).AnyTimes()
		mockArticleRepo.EXPECT().Update(gomock.Any(), gomock.Any()).Return(nil).AnyTimes()
		u := repo.NewArticleUsecase(mockArticleRepo, mockSesRepo)
		err := u.Update(context.TODO(), login, &mockArticle)
		assert.NoError(t, err)
	})
}

func TestDelete(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()
	mockSesRepo := smocks.NewMockSessionRepository(ctrl)
	mockArticleRepo := mocks.NewMockArticleRepository(ctrl)
	login := "Iam"

	t.Run("success", func(t *testing.T) {
		mockSesRepo.EXPECT().GetSessionLogin(gomock.Eq(context.TODO()), gomock.Eq(login)).Return(login, nil).AnyTimes()
		mockArticleRepo.EXPECT().Delete(gomock.Any(), login, int64(1)).Return(nil).AnyTimes()
		u := repo.NewArticleUsecase(mockArticleRepo, mockSesRepo)
		err := u.Delete(context.TODO(), login, "1")
		assert.NoError(t, err)
	})
}
