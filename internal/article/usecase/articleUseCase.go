package article

import (
	"context"
	"strconv"

	"github.com/pkg/errors"

	amodels "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/article/models"
)

type articleUseCase struct {
	articleRepo amodels.ArticleRepository
}

func NewArticleUseCase(articleRepo amodels.ArticleRepository) amodels.ArticleUseCase {
	return &articleUseCase{articleRepo}
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
		return nil, errors.Wrap(err, "articleUsecase/Fetch")
	}

	result, err = m.articleRepo.Fetch(ctx, from, chunkSize)
	return result, errors.Wrap(err, "articleUsecase/Fetch")
}

func (m *articleUseCase) GetByID(ctx context.Context, id int64) (result amodels.Article, err error) {
	result, err = m.articleRepo.GetByID(ctx, id)
	return result, errors.Wrap(err, "articleUsecase/GetByID")
}

func (m *articleUseCase) GetByTag(ctx context.Context, tag string) (result []amodels.Article, err error) {
	result, err = m.articleRepo.GetByTag(ctx, tag)
	return result, errors.Wrap(err, "articleUsecase/GetByTag")
}

func (m *articleUseCase) GetByAuthor(ctx context.Context, author string) (result []amodels.Article, err error) {
	result, err = m.articleRepo.GetByAuthor(ctx, author)
	return result, errors.Wrap(err, "articleUsecase/GetByAuthor")
}

func (m *articleUseCase) Store(ctx context.Context, a *amodels.Article) error {
	err := m.articleRepo.Store(ctx, a)
	return errors.Wrap(err, "articleUsecase/Store")
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
		return errors.Wrap(err, "articleUsecase/Delete")
	}
	err = m.articleRepo.Delete(ctx, int64(idInt))
	return errors.Wrap(err, "articleUsecase/Delete")
}
func (m *articleUseCase) Update(ctx context.Context, a *amodels.Article) error {
	err := m.articleRepo.Update(ctx, a)
	return errors.Wrap(err, "articleUsecase/Update")
}
