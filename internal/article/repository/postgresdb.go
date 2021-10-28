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

func articleConv(val amodels.DbArticle, Db *sqlx.DB) (amodels.Article, error) {
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
	article.Tags = append(article.Tags, "FUBAR")
	rows, err := Db.Queryx(`select c.tag from categories c
	inner join categories_articles ca  on c.Id = ca.categories_id
	inner join articles a on a.Id = ca.articles_id
	where a.StringId = $1;`, val.StringId)
	if err != nil {
		return article, err
	}
	var mytag string
	for rows.Next() {
		err = rows.Scan(&mytag)
		if err != nil {
			return article, err
		}
		article.Tags = append(article.Tags, mytag)
		//fmt.Printf("%s\n", mytag)
	}
	return article, nil
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
	rows, err := m.Db.Queryx("SELECT count(*) FROM articles;")
	if err != nil {
		return ChunkData, err
	}
	var count int
	for rows.Next() {
		err = rows.Scan(&count)
		if err != nil {
			return ChunkData, err
		}
	}
	// fmt.Println(count)
	if count <= from+chunkSize {
		from = count - chunkSize
	}

	rows, err = m.Db.Queryx("SELECT * FROM ARTICLES LIMIT $1 OFFSET $2", chunkSize, from)
	// rows, err = m.Db.Queryx("SELECT * FROM ARTICLES")
	if err != nil {
		return ChunkData, err
	}
	var newArticle amodels.DbArticle
	var outArticle amodels.Article
	for rows.Next() {
		err = rows.StructScan(&newArticle)
		if err != nil {
			return ChunkData, err
		}
		outArticle, err = articleConv(newArticle, m.Db)
		if err != nil {
			return ChunkData, err
		}
		ChunkData = append(ChunkData, outArticle)
		//fmt.Println(newArticle.Id, newArticle.PreviewUrl)
	}

	return ChunkData[:5], nil
}

func (m *psqlArticleRepository) GetByID(ctx context.Context, id int64) (result amodels.Article, err error) {
	rows, err := m.Db.Queryx("SELECT * FROM ARTICLES WHERE articles.StringId = $1", id)
	var outArticle amodels.Article
	if err != nil {
		return outArticle, err
	}
	var newArticle amodels.DbArticle
	for rows.Next() {
		err = rows.StructScan(&newArticle)
		if err != nil {
			return outArticle, err
		}
	}
	outArticle, err = articleConv(newArticle, m.Db)
	if err != nil {
		return outArticle, err
	}
	return outArticle, nil
}

func (m *psqlArticleRepository) GetByTag(ctx context.Context, tag string) (result []amodels.Article, err error) {
	rows, err := m.Db.Queryx(`select a.* from categories c
	inner join categories_articles ca  on c.Id = ca.categories_id
	inner join articles a on a.Id = ca.articles_id
	where c.tag = $1;`, tag)
	var articles []amodels.Article
	if err != nil {
		return articles, err
	}
	var outArticle amodels.Article
	var newArticle amodels.DbArticle
	for rows.Next() {
		err = rows.StructScan(&newArticle)
		if err != nil {
			return articles, err
		}
		outArticle, err = articleConv(newArticle, m.Db)
		if err != nil {
			return articles, err
		}
		articles = append(articles, outArticle)
	}
	return articles, nil
}
func (m *psqlArticleRepository) GetByAuthor(ctx context.Context, author string) (result []amodels.Article, err error) {
	rows, err := m.Db.Queryx("SELECT * FROM ARTICLES WHERE articles.AuthorName = $1", author)
	var articles []amodels.Article
	if err != nil {
		return articles, err
	}
	var outArticle amodels.Article
	var newArticle amodels.DbArticle
	for rows.Next() {
		err = rows.StructScan(&newArticle)
		if err != nil {
			return articles, err
		}
		outArticle, err = articleConv(newArticle, m.Db)
		if err != nil {
			return articles, err
		}
		articles = append(articles, outArticle)
	}
	return articles, nil
}
func (m *psqlArticleRepository) Update(ctx context.Context, a *amodels.Article) error {
	return nil
}
func (m *psqlArticleRepository) Store(ctx context.Context, a *amodels.Article) error {
	return nil
}
func (m *psqlArticleRepository) Delete(ctx context.Context, id int64) error {
	return nil
}
