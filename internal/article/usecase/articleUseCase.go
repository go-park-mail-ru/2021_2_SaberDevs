package article

import (
	"context"
	"net/http"
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
	return &articleUsecase{sessionRepo, articleRepo}
}

func IdToString(id string) (int, error) {
	if id == "" {
		id = "0"
	}
	idInt, err := strconv.Atoi(id)
	return idInt, err
}

func artOut(a *amodels.OutArticle) *amodels.Article {
	var out amodels.Article
	out.AuthorAvatar = a.Author.AvatarUrl
	out.AuthorName = a.Author.Name
	out.AuthorUrl = a.Author.AvatarUrl
	out.Comments = a.Comments
	out.CommentsUrl = a.CommentsUrl
	out.Id = a.Id
	out.Likes = a.Likes
	out.PreviewUrl = a.PreviewUrl
	out.Tags = a.Tags
	out.Text = a.Text
	out.Title = a.Title
	return &out
}

func (m *articleUsecase) Fetch(ctx context.Context, idLastLoaded string, chunkSize int) (result []amodels.OutArticle, err error) {
	from, err := IdToString(idLastLoaded)
	if err != nil {
		return nil, errors.Wrap(err, "articleUsecase/Fetch")
	}

	result, err = m.articleRepo.Fetch(ctx, from, chunkSize)
	return result, errors.Wrap(err, "articleUsecase/Fetch")
}

func (m *articleUsecase) GetByID(ctx context.Context, id int64) (result amodels.OutArticle, err error) {
	result, err = m.articleRepo.GetByID(ctx, id)
	return result, errors.Wrap(err, "articleUsecase/GetByID")
}

func (m *articleUsecase) GetByTag(ctx context.Context, tag string, idLastLoaded string, chunkSize int) (result []amodels.OutArticle, err error) {
	from, err := IdToString(idLastLoaded)
	if err != nil {
		return nil, errors.Wrap(err, "articleUsecase/GetByTag")
	}
	result, err = m.articleRepo.GetByTag(ctx, tag, from, chunkSize)
	return result, errors.Wrap(err, "articleUsecase/GetByTag")
}

func (m *articleUsecase) GetByAuthor(ctx context.Context, author string, idLastLoaded string, chunkSize int) (result []amodels.OutArticle, err error) {
	from, err := IdToString(idLastLoaded)
	if err != nil {
		return nil, errors.Wrap(err, "articleUsecase/GetByAuthor")
	}
	result, err = m.articleRepo.GetByAuthor(ctx, author, from, chunkSize)
	return result, errors.Wrap(err, "articleUsecase/GetByAuthor")
}

func (m *articleUsecase) Store(ctx context.Context, c *http.Cookie, a *amodels.ArticleCreate) (int, error) {
	newArticle := amodels.Article{}
	newArticle.Text = a.Text
	newArticle.Tags = a.Tags
	newArticle.Title = a.Title

	AuthorName, err := m.sessionRepo.GetSessionLogin(ctx, c.Value)
	if err != nil {
		return 0, errors.Wrap(err, "articleUsecase/Delete")
	}
	newArticle.AuthorName = AuthorName
	Id, err := m.articleRepo.Store(ctx, &newArticle)
	return Id, errors.Wrap(err, "articleUsecase/Store")
}

func (m *articleUsecase) Delete(ctx context.Context, id string) error {
	idInt, err := IdToString(id)
	if err != nil {
		return errors.Wrap(err, "articleUsecase/Delete")
	}
	err = m.articleRepo.Delete(ctx, int64(idInt))
	return errors.Wrap(err, "articleUsecase/Delete")
}

func (m *articleUsecase) Update(ctx context.Context, a *amodels.ArticleUpdate) error {
	idInt, err := IdToString(a.Id)
	if err != nil {
		return errors.Wrap(err, "articleUsecase/Delete")
	}
	newArticle, err := m.GetByID(ctx, int64(idInt))
	upArt := artOut(&newArticle)
	if err != nil {
		return errors.Wrap(err, "articleUsecase/Delete")
	}
	upArt.Text = a.Text
	upArt.Tags = a.Tags
	upArt.Title = a.Title
	err = m.articleRepo.Update(ctx, upArt)
	return errors.Wrap(err, "articleUsecase/Update")
}
