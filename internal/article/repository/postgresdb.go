package article

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	amodels "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/article/models"
	"github.com/go-park-mail-ru/2021_2_SaberDevs/internal/data"
	sbErr "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/syberErrors"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type psqlArticleRepository struct {
	Db *sqlx.DB
}

func NewArticleRepository(db *sqlx.DB) amodels.ArticleRepository {
	return &psqlArticleRepository{db}
}

const previewLen = 50

const tagsLoad = `select c.tag from categories c
inner join categories_articles ca  on c.Id = ca.categories_id
inner join articles a on a.Id = ca.articles_id
where a.Id = $1;`

const multiArtTags = `select a.Id, c.tag from categories c
inner join categories_articles ca  on c.Id = ca.categories_id
inner join articles a on a.Id = ca.articles_id
where a.Id in (`

const byTag = "articleRepository/GetByTag"

const byAuthor = "articleRepository/GetByAuthor"

const toFetch = "articleRepository/Fetch"

func previewConv(val amodels.DbArticle, auth amodels.Author) amodels.Preview {
	var article amodels.Preview
	article.Author = auth
	article.Comments = val.Comments
	article.DateTime = val.DateTime
	article.CommentsUrl = val.CommentsUrl
	article.Id = fmt.Sprint(val.Id)
	article.Likes = val.Likes
	article.PreviewUrl = val.PreviewUrl
	article.Title = val.Title
	if len(val.Text) <= previewLen {
		article.Text = val.Text
	} else {
		temp := strings.Split(val.Text[:50], "")
		article.Text = strings.Join(temp, "")
	}
	return article
}

func (m *psqlArticleRepository) uploadTags(ChunkData []amodels.Preview, funcName string) ([]amodels.Preview, error) {
	funcName = funcName + "/uploadTags"
	schema := multiArtTags
	var ids []interface{}
	for i, data := range ChunkData {
		ids = append(ids, data.Id)
		schema = schema + `$` + fmt.Sprint(i+1)
		if i < len(ChunkData)-1 {
			schema = schema + `,`
		}
	}
	schema = schema + `) order by a.Id, c.tag;`

	rows, err := m.Db.Queryx(schema, ids...)
	if err != nil {
		return ChunkData, sbErr.ErrDbError{
			Reason:   err.Error(),
			Function: funcName,
		}
	}
	var newtag string
	var id int
	i := 0
	for rows.Next() {
		err = rows.Scan(&id, &newtag)
		if err != nil {
			return ChunkData, sbErr.ErrDbError{
				Reason:   err.Error(),
				Function: funcName,
			}
		}
		myid, _ := strconv.Atoi(ChunkData[i].Id)
		if myid == id {
			ChunkData[i].Tags = append(ChunkData[i].Tags, newtag)
		} else {
			i++
			ChunkData[i].Tags = append(ChunkData[i].Tags, newtag)
		}
	}
	return ChunkData, nil
}

func (m *psqlArticleRepository) addTags(ChunkData []amodels.Preview, funcName string, rows *sqlx.Rows, overCount bool, arts []amodels.DbArticle) ([]amodels.Preview, error) {
	funcName = funcName + "/addTags"
	var auths []amodels.Author
	var newAuth amodels.Author
	for rows.Next() {
		err := rows.StructScan(&newAuth)
		if err != nil {
			return ChunkData, sbErr.ErrDbError{
				Reason:   err.Error(),
				Function: funcName,
			}
		}
		auths = append(auths, newAuth)
	}
	for i, article := range arts {
		outArticle := previewConv(article, auths[i])
		ChunkData = append(ChunkData, outArticle)
	}
	ChunkData, err := m.uploadTags(ChunkData, funcName)
	if err != nil {
		return ChunkData, sbErr.ErrDbError{
			Reason:   err.Error(),
			Function: funcName,
		}
	}
	if overCount {
		ChunkData = append(ChunkData, data.End)
	}
	return ChunkData, nil
}

func fullArticleConv(val amodels.DbArticle, Db *sqlx.DB, auth amodels.Author) (amodels.FullArticle, error) {
	var article amodels.FullArticle
	article.Author = auth
	article.Comments = val.Comments
	article.DateTime = val.DateTime
	article.CommentsUrl = val.CommentsUrl
	article.Id = fmt.Sprint(val.Id)
	article.Likes = val.Likes
	article.PreviewUrl = val.PreviewUrl
	article.Title = val.Title
	article.Text = val.Text
	rows, err := Db.Queryx(tagsLoad, val.Id)
	if err != nil {
		return article, sbErr.ErrDbError{
			Reason:   err.Error(),
			Function: "articleRepository/fullArticleConv",
		}
	}
	var mytag string
	for rows.Next() {
		err = rows.Scan(&mytag)
		if err != nil {
			return article, sbErr.ErrDbError{
				Reason:   err.Error(),
				Function: "articleRepository/fullArticleConv",
			}
		}
		article.Tags = append(article.Tags, mytag)
	}
	return article, nil
}

func (m *psqlArticleRepository) limitChecker(schemaCount string, from, chunkSize int, args ...interface{}) (int, []amodels.Preview, bool, error) {
	var ChunkData []amodels.Preview
	overCount := false
	var count int
	err := m.Db.Get(&count, schemaCount, args...)
	if err != nil {
		return chunkSize, ChunkData, overCount, sbErr.ErrDbError{
			Reason:   err.Error(),
			Function: "articleRepository/limitChecker",
		}
	}

	if count <= from {
		ChunkData = append(ChunkData, data.End)
		overCount = true
	}

	if (count > from) && (count < from+chunkSize) {
		chunkSize = count - from
		overCount = true
	}
	return chunkSize, ChunkData, overCount, nil
}

func (m *psqlArticleRepository) Fetch(ctx context.Context, from, chunkSize int) (result []amodels.Preview, err error) {
	schemaCount := "SELECT count(*) FROM articles;"
	chunkSize, ChunkData, overCount, err := m.limitChecker(schemaCount, from, chunkSize)
	if err != nil || len(ChunkData) > 0 {
		return ChunkData, err
	}
	rows, err := m.Db.Queryx("SELECT * FROM ARTICLES ORDER BY Id LIMIT $1 OFFSET $2", chunkSize, from)
	if err != nil {
		return ChunkData, sbErr.ErrDbError{
			Reason:   err.Error(),
			Function: toFetch,
		}
	}
	var newArticle amodels.DbArticle
	var arts []amodels.DbArticle
	for rows.Next() {
		err = rows.StructScan(&newArticle)
		if err != nil {
			return ChunkData, sbErr.ErrDbError{
				Reason:   err.Error(),
				Function: toFetch,
			}
		}
		arts = append(arts, newArticle)
	}

	rows, err = m.Db.Queryx("SELECT AU.* FROM ARTICLES AS AR INNER JOIN AUTHOR AS AU ON AU.LOGIN = AR.AuthorName ORDER BY AR.Id LIMIT $1 OFFSET $2", chunkSize, from)
	if err != nil {
		return ChunkData, sbErr.ErrDbError{
			Reason:   err.Error(),
			Function: "articleRepository/Fetch",
		}
	}
	ChunkData, err = m.addTags(ChunkData, toFetch, rows, overCount, arts)
	return ChunkData, err
}

func (m *psqlArticleRepository) GetByID(ctx context.Context, id int64) (result amodels.FullArticle, err error) {
	var newArticle amodels.DbArticle
	err = m.Db.Get(&newArticle, "SELECT * FROM ARTICLES WHERE articles.Id = $1", id)
	var outArticle amodels.FullArticle
	if err != nil {
		return outArticle, sbErr.ErrDbError{
			Reason:   err.Error(),
			Function: "articleRepository/GetByID",
		}
	}
	var newAuth amodels.Author
	err = m.Db.Get(&newAuth, `SELECT AU.* FROM ARTICLES AS AR INNER JOIN AUTHOR AS AU ON AU.LOGIN = AR.AuthorName WHERE AR.ID = $1`, id)
	if err != nil {
		return outArticle, sbErr.ErrDbError{
			Reason:   err.Error(),
			Function: "articleRepository/GetByID",
		}
	}

	outArticle, err = fullArticleConv(newArticle, m.Db, newAuth)
	if err != nil {
		return outArticle, sbErr.ErrDbError{
			Reason:   err.Error(),
			Function: "articleRepository/GetByID",
		}
	}
	return outArticle, nil
}

func (m *psqlArticleRepository) GetByTag(ctx context.Context, tag string, from, chunkSize int) (result []amodels.Preview, err error) {

	schemaCount := `SELECT count(*) FROM  categories c
	inner join categories_articles ca  on c.Id = ca.categories_id
	inner join articles a on a.Id = ca.articles_id
	where c.tag = $1;`
	chunkSize, ChunkData, overCount, err := m.limitChecker(schemaCount, from, chunkSize, tag)
	if err != nil || len(ChunkData) > 0 {
		return ChunkData, err
	}
	rows, err := m.Db.Queryx(`select a.* from categories c
	inner join categories_articles ca  on c.Id = ca.categories_id
	inner join articles a on a.Id = ca.articles_id
	where c.tag = $1 LIMIT $2 OFFSET $3`, tag, chunkSize, from)
	if err != nil {
		return ChunkData, sbErr.ErrDbError{
			Reason:   err.Error(),
			Function: byTag,
		}
	}
	var newArticle amodels.DbArticle
	var arts []amodels.DbArticle
	for rows.Next() {
		err = rows.StructScan(&newArticle)
		if err != nil {
			return ChunkData, sbErr.ErrDbError{
				Reason:   err.Error(),
				Function: byTag,
			}
		}
		arts = append(arts, newArticle)
	}

	rows, err = m.Db.Queryx(`SELECT AU.* FROM categories c
	inner join categories_articles ca  on c.Id = ca.categories_id
	inner join articles a on a.Id = ca.articles_id
	INNER JOIN AUTHOR AS AU ON AU.LOGIN = A.AUTHORNAME where c.tag = $1 ORDER BY a.Id LIMIT $2 OFFSET $3`, tag, chunkSize, from)
	if err != nil {
		return ChunkData, sbErr.ErrDbError{
			Reason:   err.Error(),
			Function: byTag,
		}
	}
	ChunkData, err = m.addTags(ChunkData, byTag, rows, overCount, arts)
	return ChunkData, err
}
func (m *psqlArticleRepository) GetByAuthor(ctx context.Context, author string, from, chunkSize int) (result []amodels.Preview, err error) {
	schemaCount := `SELECT count(*) FROM ARTICLES WHERE articles.AuthorName = $1`
	chunkSize, ChunkData, overCount, err := m.limitChecker(schemaCount, from, chunkSize, author)
	if err != nil || len(ChunkData) > 0 {
		return ChunkData, err
	}
	rows, err := m.Db.Queryx("SELECT * FROM ARTICLES WHERE articles.AuthorName = $1 ORDER BY Id LIMIT $2 OFFSET $3", author, chunkSize, from)
	if err != nil {
		return ChunkData, sbErr.ErrDbError{
			Reason:   err.Error(),
			Function: byAuthor,
		}
	}
	var newArticle amodels.DbArticle
	var arts []amodels.DbArticle
	for rows.Next() {
		err = rows.StructScan(&newArticle)
		if err != nil {
			return ChunkData, sbErr.ErrDbError{
				Reason:   err.Error(),
				Function: byAuthor,
			}
		}
		arts = append(arts, newArticle)
	}
	rows, err = m.Db.Queryx("SELECT AU.* FROM ARTICLES AR JOIN AUTHOR AU ON AU.LOGIN = AR.AUTHORNAME  WHERE AU.LOGIN = $1 ORDER BY AR.Id LIMIT $2 OFFSET $3", author, chunkSize, from)
	if err != nil {
		return ChunkData, sbErr.ErrDbError{
			Reason:   err.Error(),
			Function: byAuthor,
		}
	}
	ChunkData, err = m.addTags(ChunkData, byAuthor, rows, overCount, arts)
	return ChunkData, err
}

func (m *psqlArticleRepository) Store(ctx context.Context, a *amodels.Article) (int, error) {
	insertArticle := `INSERT INTO articles (DateTime, PreviewUrl, Title, Text, AuthorName, CommentsUrl, Comments, Likes) VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING ID;`
	rows, err := m.Db.Query(insertArticle, a.DateTime, a.PreviewUrl, a.Title, a.Text, a.AuthorName, a.CommentsUrl, a.Comments, a.Likes)
	if err != nil {
		return 0, sbErr.ErrDbError{
			Reason:   err.Error(),
			Function: "articleRepository/Store",
		}
	}
	var Id int
	for rows.Next() {
		err = rows.Scan(&Id)
		if err != nil {
			return Id, sbErr.ErrDbError{
				Reason:   err.Error(),
				Function: "articleRepository/Store",
			}
		}
	}

	insertCat := `INSERT INTO categories (tag) VALUES ($1) ON CONFLICT DO NOTHING;`
	for _, data := range a.Tags {
		_, err = m.Db.Exec(insertCat, data)
		if err != nil {
			return Id, sbErr.ErrDbError{
				Reason:   err.Error(),
				Function: "articleRepository/Store",
			}
		}
	}
	insert_junc := `INSERT INTO categories_articles (articles_id, categories_id) VALUES 
	((SELECT Id FROM articles WHERE Id = $1) ,    
	(SELECT Id FROM categories WHERE tag = $2)) ON CONFLICT DO NOTHING;`
	for _, v := range a.Tags {
		_, err = m.Db.Exec(insert_junc, Id, v)
		if err != nil {
			return Id, sbErr.ErrDbError{
				Reason:   err.Error(),
				Function: "articleRepository/Store",
			}
		}
	}
	return Id, nil
}

func (m *psqlArticleRepository) Delete(ctx context.Context, id int64) error {
	_, err := m.Db.Exec("DELETE FROM ARTICLES WHERE articles.Id = $1", id)
	if err != nil {
		return sbErr.ErrDbError{
			Reason:   err.Error(),
			Function: "articleRepository/Delete",
		}
	}
	return nil
}
func (m *psqlArticleRepository) Update(ctx context.Context, a *amodels.Article) error {
	uniqId, err := strconv.Atoi(a.Id)
	if err != nil {
		return sbErr.ErrDbError{
			Reason:   err.Error(),
			Function: "articleRepository/Update",
		}
	}
	updateArticle := `UPDATE articles SET DateTime = $1, Title = $2, Text = $3 WHERE articles.Id  = $4`
	_, err = m.Db.Query(updateArticle, time.Now().Format("2006/1/2 15:04"), a.Title, a.Text, uniqId)
	if err != nil {
		return sbErr.ErrDbError{
			Reason:   err.Error(),
			Function: "articleRepository/Update",
		}
	}
	insertCat := `INSERT INTO categories (tag) VALUES ($1) ON CONFLICT DO NOTHING;`
	for _, data := range a.Tags {
		_, err = m.Db.Exec(insertCat, data)
		if err != nil {
			return sbErr.ErrDbError{
				Reason:   err.Error(),
				Function: "articleRepository/Store",
			}
		}
	}
	insert_junc := `INSERT INTO categories_articles (articles_id, categories_id) VALUES
	((SELECT articles.Id FROM articles WHERE articles.Id = $1) ,
	(SELECT categories.Id FROM categories WHERE categories.tag = $2)) ON CONFLICT DO NOTHING;`
	for _, v := range a.Tags {
		_, err = m.Db.Exec(insert_junc, uniqId, v)
		if err != nil {
			return sbErr.ErrDbError{
				Reason:   err.Error(),
				Function: "articleRepository/Store",
			}
		}
	}
	return nil
}
