package article

import (
	"context"
	"strconv"

	smodels "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/session/models"
	"github.com/pkg/errors"

	amodels "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/article/models"
)

type articleUsecase struct {
	sessionRepo smodels.SessionRepository
	articleRepo amodels.ArticleRepository
}

func NewArticleUsecase(articleRepo amodels.ArticleRepository, sessionRepo smodels.SessionRepository) amodels.ArticleUsecase {
	return &articleUsecase{articleRepo, sessionRepo}
}

func (m *articleUsecase) Fetch(ctx context.Context, idLastLoaded string, chunkSize int) (result []amodels.Article, err error) {
	if idLastLoaded == "" {
		idLastLoaded = "0"
	}
	if idLastLoaded == "end" {
		idLastLoaded = "12"
	}

	from, err := strconv.Atoi(idLastLoaded)
	if err != nil {
		return nil, errors.Wrap(err, "articleUsecase/Fetch")
	}

	result, err = m.articleRepo.Fetch(ctx, from, chunkSize)
	return result, errors.Wrap(err, "articleUsecase/Fetch")
}

func (m *articleUsecase) GetByID(ctx context.Context, id int64) (result amodels.Article, err error) {
	result, err = m.articleRepo.GetByID(ctx, id)
	return result, errors.Wrap(err, "articleUsecase/GetByID")
}

func (m *articleUsecase) GetByTag(ctx context.Context, tag string) (result []amodels.Article, err error) {
	result, err = m.articleRepo.GetByTag(ctx, tag)
	return result, errors.Wrap(err, "articleUsecase/GetByTag")
}

func (m *articleUsecase) GetByAuthor(ctx context.Context, author string) (result []amodels.Article, err error) {
	result, err = m.articleRepo.GetByAuthor(ctx, author)
	return result, errors.Wrap(err, "articleUsecase/GetByAuthor")
}

func (m *articleUsecase) Store(ctx context.Context, a *amodels.ArticleCreate) (int, error) {
	newArticle := amodels.Article{}
	newArticle.Text = a.Text
	newArticle.Tags = a.Tags
	newArticle.Title = a.Title
	newArticle.Id = "0"
	newArticle.AuthorName = session

	err := m.articleRepo.Store(ctx, newArticle)
	if err != nil {
		return 0, errors.Wrap(err, "articleUsecase/Store")
	}

}

func (m *articleUsecase) Delete(ctx context.Context, id string) error {
	if id == "" {
		id = "0"
	}
	if id == "end" {
		id = "12"
	}

	idInt, err := strconv.Atoi(id)
	if err != nil {
		return errors.Wrap(err, "articleUsecase/Delete")
	}
	err = m.articleRepo.Delete(ctx, int64(idInt))
	return errors.Wrap(err, "articleUsecase/Delete")
}
func (m *articleUsecase) Update(ctx context.Context, a *amodels.ArticleUpdate) error {

	if a.Id == "" {
		a.Id = "0"
	}
	if a.Id == "end" {
		a.Id = "12"
	}

	idInt, err := strconv.Atoi(a.Id)
	newArticle, err := m.GetByID(ctx, int64(idInt))
	if err != nil {
		return errors.Wrap(err, "articleUsecase/Delete")
	}
	newArticle.Text = a.Text
	newArticle.Tags = a.Tags
	newArticle.Title = a.Title
	newArticle.Id = a.Id
	err = m.articleRepo.Update(ctx, &newArticle)
	return errors.Wrap(err, "articleUsecase/Update")
}
