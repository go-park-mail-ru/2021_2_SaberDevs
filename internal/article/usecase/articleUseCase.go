package article

import (
	"strconv"

	"context"

	"github.com/microcosm-cc/bluemonday"

	repository "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/article/repository"
	"github.com/jmoiron/sqlx"

	amodels "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/article/models"
)

type articleUseCase struct {
	articleRepo amodels.ArticleRepository
}

func SanitizeArticle(a *amodels.Article) *amodels.Article {
	s := bluemonday.StrictPolicy()
	l := bluemonday.UGCPolicy()
	a.AuthorAvatar = s.Sanitize(a.AuthorAvatar)
	a.AuthorName = s.Sanitize(a.AuthorName)
	a.AuthorUrl = s.Sanitize(a.AuthorUrl)
	//a.Comments = s.Sanitize(a.Comments) //not a string
	a.CommentsUrl = s.Sanitize(a.CommentsUrl)
	a.Id = s.Sanitize(a.Id)
	// a.Likes = s.Sanitize(a.Likes)//not a string
	a.PreviewUrl = s.Sanitize(a.PreviewUrl)
	for i := range a.Tags {
		a.Tags[i] = s.Sanitize(a.Tags[i])
	}
	a.Text = l.Sanitize(a.Text)
	a.Title = s.Sanitize(a.Title)
	return a
}

func NewArticleUsecase(db *sqlx.DB) amodels.ArticleUseCase {
	return &articleUseCase{repository.NewpsqlArticleRepository(db)}
}

func (m *articleUseCase) Fetch(ctx context.Context, idLastLoaded string, chunkSize int) (result []amodels.Article, err error) {
	if idLastLoaded == "" {
		idLastLoaded = "0"
	}
	if idLastLoaded == "end" {
		idLastLoaded = "12"
	}

	from, err := strconv.Atoi(idLastLoaded)
	if err != nil {
		return nil, err
	}

	result, err = m.articleRepo.Fetch(ctx, from, chunkSize)
	return result, err
}

func (m *articleUseCase) GetByID(ctx context.Context, id int64) (result amodels.Article, err error) {
	result, err = m.articleRepo.GetByID(ctx, id)
	return result, err
}

func (m *articleUseCase) GetByTag(ctx context.Context, tag string) (result []amodels.Article, err error) {
	result, err = m.articleRepo.GetByTag(ctx, tag)
	return result, err
}

func (m *articleUseCase) GetByAuthor(ctx context.Context, author string) (result []amodels.Article, err error) {
	result, err = m.articleRepo.GetByTag(ctx, author)
	return result, err
}

func (m *articleUseCase) Store(ctx context.Context, a *amodels.Article) error {
	a = SanitizeArticle(a)
	err := m.articleRepo.Store(ctx, a)
	return err
}

func (m *articleUseCase) Delete(ctx context.Context, id string) error {
	if id == "" {
		id = "0"
	}
	if id == "end" {
		id = "12"
	}

	idInt, err := strconv.Atoi(id)
	if err != nil {
		return err
	}
	err = m.articleRepo.Delete(ctx, int64(idInt))
	return err
}
func (m *articleUseCase) Update(ctx context.Context, a *amodels.Article) error {
	err := m.articleRepo.Update(ctx, a)
	return err
}
