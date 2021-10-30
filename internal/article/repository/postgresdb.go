//package main

package article

import (
	"context"
	"fmt"
	"strconv"

	"github.com/microcosm-cc/bluemonday"

	amodels "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/article/models"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type psqlArticleRepository struct {
	Db *sqlx.DB
}

func NewpsqlArticleRepository(db *sqlx.DB) amodels.ArticleRepository {
	//TODO defer db.Close()
	return &psqlArticleRepository{db}
}

func SanitizeArticle(a *amodels.Article) *amodels.Article {
	// s := bluemonday.NewPolicy()
	//s.AllowStandardURLs()
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
		a.Tags[i] = l.Sanitize(a.Tags[i])
	}
	a.Text = s.Sanitize(a.Text)
	a.Title = s.Sanitize(a.Title)
	return a
}

const previewLen = 50

func articleShortConv(val amodels.DbArticle, Db *sqlx.DB) (amodels.Article, error) {
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
		article.Text = val.Text[:50]
	}
	article.Title = val.Title
	return article, nil
}
func fullArticleConv(val amodels.DbArticle, Db *sqlx.DB) (amodels.Article, error) {
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
		article.Text = val.Text
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
		from = count - chunkSize - 1
	}

	rows, err = m.Db.Queryx("SELECT * FROM ARTICLES ORDER BY Id LIMIT $1 OFFSET $2", chunkSize, from)
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
		outArticle, err = articleShortConv(newArticle, m.Db)
		if err != nil {
			return ChunkData, err
		}
		ChunkData = append(ChunkData, outArticle)
	}
	schema := `select a.StringId, a.Id, c.tag from categories c
	inner join categories_articles ca  on c.Id = ca.categories_id
	inner join articles a on a.Id = ca.articles_id
	where a.StringId in (`
	var ids []interface{}
	for i, data := range ChunkData {
		ids = append(ids, data.Id)
		schema = schema + `$` + fmt.Sprint(i+1)
		if i < len(ChunkData)-1 {
			schema = schema + `,`
		}
	}
	schema = schema + `) order by a.Id, c.tag;`

	rows, err = m.Db.Queryx(schema, ids...)
	if err != nil {
		return ChunkData, err
	}
	var newtag string
	var strid string
	var id int
	i := 0
	for rows.Next() {
		err = rows.Scan(&strid, &id, &newtag)
		if err != nil {
			return ChunkData, err
		}
		if ChunkData[i].Id == strid {
			ChunkData[i].Tags = append(ChunkData[i].Tags, newtag)
		} else {
			i++
			ChunkData[i].Tags = append(ChunkData[i].Tags, newtag)
		}
	}
	return ChunkData, nil
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
	outArticle, err = fullArticleConv(newArticle, m.Db)
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
		outArticle, err = fullArticleConv(newArticle, m.Db)
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
		outArticle, err = fullArticleConv(newArticle, m.Db)
		if err != nil {
			return articles, err
		}
		articles = append(articles, outArticle)
	}
	return articles, nil
}

func (m *psqlArticleRepository) Store(ctx context.Context, a *amodels.Article) error {
	a = SanitizeArticle(a)
	insertArticle := `INSERT INTO articles (StringId, PreviewUrl, Title, Text, AuthorUrl, AuthorName, AuthorAvatar, CommentsUrl, Comments, Likes) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10);`
	_, err := m.Db.Exec(insertArticle, a.Id, a.PreviewUrl, a.Title, a.Text, a.AuthorUrl, a.AuthorName, a.AuthorAvatar, a.CommentsUrl, a.Comments, a.Likes)
	if err != nil {
		return err
	}
	var tags []interface{}
	exTags := make(map[string]int)
	//find existing tags
	schema := `SELECT tag FROM categories WHERE tag IN (`
	for i, tag := range a.Tags {
		exTags[tag] = 0
		tags = append(tags, tag)
		schema = schema + `$` + fmt.Sprint(i+1)
		if i < len(a.Tags)-1 {
			schema = schema + `,`
		}
	}
	schema = schema + `)`
	rows, err := m.Db.Queryx(schema, tags...)
	if err != nil {
		return err
	}
	var tag string
	for rows.Next() {
		err = rows.Scan(&tag)
		if err != nil {
			return err
		}
		exTags[tag] = 1
	}
	//create new tags
	var newtags []interface{}
	for _, v := range a.Tags {
		if exTags[v] == 0 {
			newtags = append(newtags, v)
		}
	}

	insertCat := `INSERT INTO categories (tag) VALUES ($1);`
	for _, data := range newtags {
		_, err = m.Db.Exec(insertCat, data)
		if err != nil {
			return err
		}
	}
	insert_junc := `INSERT INTO categories_articles (articles_id, categories_id) VALUES 
	((SELECT Id FROM articles WHERE StringId = $1) ,    
	(SELECT Id FROM categories WHERE tag = $2));`
	for _, v := range a.Tags {
		_, err = m.Db.Exec(insert_junc, a.Id, v)
		if err != nil {
			return err
		}
	}
	return nil
}

func (m *psqlArticleRepository) Delete(ctx context.Context, id int64) error {
	_, err := m.Db.Exec("DELETE FROM ARTICLES WHERE articles.StringId = $1", id)
	if err != nil {
		return err
	}
	return nil
}
func (m *psqlArticleRepository) Update(ctx context.Context, a *amodels.Article) error {
	if a.Id == "" {
		a.Id = "0"
	}
	if a.Id == "end" {
		a.Id = "12"
	}

	uniqId, err := strconv.Atoi(a.Id)
	if err != nil {
		return err
	}
	err = m.Delete(ctx, int64(uniqId))
	if err != nil {
		return err
	}
	err = m.Store(ctx, a)
	if err != nil {
		return err
	}
	return nil
}
