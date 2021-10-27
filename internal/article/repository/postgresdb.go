//package main

package article

import (
	"context"
	"fmt"

	amodels "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/article/models"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type psqlArticleRepository struct {
	Db *sqlx.DB
}

const previewLen = 50

func articleConv(val amodels.DbArticle, Db *sqlx.DB) amodels.Article {
	var article amodels.Article
	article.AuthorAvatar = val.AuthorAvatar
	article.AuthorName = val.AuthorName
	article.AuthorUrl = val.AuthorUrl
	article.Comments = val.Comments
	article.CommentsUrl = val.CommentsUrl
	article.Id = val.StringId
	article.Likes = val.Likes
	article.PreviewUrl = val.PreviewUrl
	if len(val.Text) <= previewLen {
		article.Text = val.Text
	} else {
		article.Text = val.Text[1:50]
	}
	article.Title = val.Title
	rows, err := Db.Queryx("SELECT * FROM TAGS")
	if err != nil {
		fmt.Println(err.Error())
	}

	for rows.Next() {
		//	err = rows.StructScan(&newArticle)
		if err != nil {
			fmt.Println(err.Error())
		}
	}
	return article
}

func NewpsqlArticleRepository() amodels.ArticleRepository {
	connStr := "user=postgres dbname=postgres password=yura11011 host=localhost sslmode=disable"
	db, err := sqlx.Open("postgres", connStr)
	if err != nil {
		fmt.Println(err.Error())
	}

	err = db.Ping()
	if err != nil {
		fmt.Println(err.Error())
	}
	//TODO defer db.Close()
	return &psqlArticleRepository{db}
}

func (m *psqlArticleRepository) Fetch(ctx context.Context, from, chunkSize int) (result []amodels.Article, err error) {

	var ChunkData []amodels.Article
	//rows, err := m.Db.Queryx("SELECT * FROM ARTICLES" LIMIT $1 OFFSET $2", from, from+chunkSize")
	rows, err := m.Db.Queryx("SELECT * FROM ARTICLES")
	if err != nil {
		fmt.Println(err.Error())
	}
	var newArticle amodels.DbArticle
	var outArticle amodels.Article
	for rows.Next() {
		err = rows.StructScan(&newArticle)
		if err != nil {
			fmt.Println(err.Error())
		}
		outArticle = articleConv(newArticle, m.Db)
		ChunkData = append(ChunkData, outArticle)
		//fmt.Println(newArticle.Id, newArticle.PreviewUrl)
	}

	return ChunkData, nil
}
