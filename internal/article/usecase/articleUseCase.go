package article

import (
	"strconv"

	echo "github.com/labstack/echo/v4"

	repository "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/article/repository"

	amodels "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/article/models"
)

type articleUseCase struct {
	articleRepo amodels.ArticleRepository
}

func newArticleUsecase(repo amodels.ArticleRepository) amodels.ArticleUseCase {
	return &articleUseCase{repo}
}

// func NewArticleUsecase() amodels.ArticleUseCase {
// 	return &articleUseCase{repository.NewDataArticleRepository()}
// }
func NewArticleUsecase() amodels.ArticleUseCase {
	return &articleUseCase{repository.NewpsqlArticleRepository()}
}

func (m *articleUseCase) Fetch(ctx echo.Context, idLastLoaded string, chunkSize int) (result []amodels.Article, err error) {
	if idLastLoaded == "" {
		idLastLoaded = "0"
	}

	from, err := strconv.Atoi(idLastLoaded)
	if err != nil {
		//c.Logger().Printf("Error: %s", err.Error())
		//return c.JSON(http.StatusNotFound, errResp.ErrNotFeedNumber)
		return nil, err
	}

	result, err = m.articleRepo.Fetch(ctx, from, chunkSize)
	return result, err
}
