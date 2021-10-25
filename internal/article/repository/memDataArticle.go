package article

import (
	"context"

	amodels "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/article/models"
	data "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/data"
)

type memDataArticleRepository struct {
	Data []amodels.Article
}

func newDataArticleRepository(Data []amodels.Article) amodels.ArticleRepository {
	return &memDataArticleRepository{Data}
}

func NewDataArticleRepository() amodels.ArticleRepository {
	return &memDataArticleRepository{data.TestData[:]}
}

func (m *memDataArticleRepository) Fetch(ctx context.Context, from, chunkSize int) (result []amodels.Article, err error) {
	var ChunkData []amodels.Article
	if from >= 0 && from+chunkSize < len(m.Data) {
		ChunkData = m.Data[from : from+chunkSize]
	} else {
		start := 0
		if len(m.Data) > chunkSize {
			start = len(m.Data) - chunkSize
		}
		ChunkData = m.Data[start : len(m.Data)-1]

	}
	return ChunkData, nil
}
