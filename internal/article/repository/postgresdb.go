//package main

package article

import (
	"log"

	amodels "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/article/models"
	"github.com/jmoiron/sqlx"
	echo "github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
)

type psqlArticleRepository struct {
	Db *sqlx.DB
}

func NewpsqlArticleRepository() amodels.ArticleRepository {
	connStr := "user=postgres dbname=postgres password=yura11011 host=localhost sslmode=disable"
	db, err := sqlx.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	//TODO defer db.Close()
	return &psqlArticleRepository{db}
}

func (m *psqlArticleRepository) Fetch(ctx echo.Context, from, chunkSize int) (result []amodels.Article, err error) {

	var ChunkData []amodels.Article
	rows, err := m.Db.Queryx("SELECT * FROM ARTICLES LIMIT $1 OFFSET $2", from, from+chunkSize)
	if err != nil {
		log.Fatal(err)
	}
	var newArticle amodels.Article
	for rows.Next() {
		err = rows.StructScan(&newArticle)
		if err != nil {
			log.Fatal(err)
		}
		ChunkData = append(ChunkData, newArticle)
		//fmt.Println(newArticle.Id, newArticle.PreviewUrl)
	}

	return ChunkData, nil
}
