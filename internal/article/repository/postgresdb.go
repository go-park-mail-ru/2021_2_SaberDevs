package article

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	wrapper "github.com/go-park-mail-ru/2021_2_SaberDevs/internal"
	"github.com/go-park-mail-ru/2021_2_SaberDevs/internal/article/models"
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

var PreviewLength = 50

const tagsLoad = `select c.tag from tags c
inner join tags_articles ca  on c.Id = ca.tags_id
inner join articles a on a.Id = ca.articles_id
where a.Id = $1;`

const multiArtTags = `select a.Id, c.tag from tags c
inner join tags_articles ca  on c.Id = ca.tags_id
inner join articles a on a.Id = ca.articles_id
where a.Id in (`

const deleteTags = `delete from tags_articles ta
where ta.articles_id  = $1;`

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
	article.Category = val.Category
	article.PreviewUrl = val.PreviewUrl
	article.Title = val.Title
	temp := strings.Split(val.Text, " ")
	previewLen := PreviewLength
	if len(temp) <= previewLen {
		article.Text = val.Text
	} else {
		article.Text = strings.Join(temp[:previewLen], " ")
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
	schema = schema + `) order by a.Id DESC;`
	rows, err := wrapper.MyQuery(m.Db, funcName, schema, ids...)
	if err != nil {
		return ChunkData, sbErr.ErrDbError{
			Reason:   err.Error(),
			Function: funcName,
		}
	}
	var newtag string
	var id int
	Tags := make(map[int][]string)
	for rows.Next() {
		err = rows.Scan(&id, &newtag)
		if err != nil {
			return ChunkData, sbErr.ErrDbError{
				Reason:   err.Error(),
				Function: funcName,
			}
		}
		slice := Tags[id]
		slice = append(slice, newtag)
		Tags[id] = slice
	}
	for i := range ChunkData {
		myid, _ := strconv.Atoi(ChunkData[i].Id)
		ChunkData[i].Tags = Tags[myid]
	}
	return ChunkData, nil
}

func (m *psqlArticleRepository) uploadAuthors(authors []string, funcName string) (map[string]amodels.Author, error) {
	funcName = funcName + "/uploadAuthors"
	ChunkData := make(map[string]amodels.Author)
	schema := "SELECT AU.ID, AU.LOGIN, AU.NAME, AU.SURNAME, AU.AVATARURL, AU.DESCRIPTION, AU.EMAIL, AU.PASSWORD, AU.SCORE FROM  AUTHOR AU   WHERE AU.LOGIN IN ("
	var ids []interface{}
	for i, data := range authors {
		ids = append(ids, data)
		schema = schema + `$` + fmt.Sprint(i+1)
		if i < len(authors)-1 {
			schema = schema + ","
		}
	}
	schema = schema + ");"
	rows, err := wrapper.MyQuery(m.Db, funcName, schema, ids...)
	if err != nil {
		return ChunkData, sbErr.ErrDbError{
			Reason:   err.Error(),
			Function: funcName,
		}
	}
	var newAuthor amodels.Author
	for rows.Next() {
		err = rows.StructScan(&newAuthor)
		if err != nil {
			return ChunkData, sbErr.ErrDbError{
				Reason:   err.Error(),
				Function: funcName,
			}
		}
		ChunkData[newAuthor.Login] = newAuthor
	}
	return ChunkData, nil
}

func (m *psqlArticleRepository) addTags(ChunkData []amodels.Preview, chunkSize int, authors map[string]amodels.Author, funcName string, arts []amodels.DbArticle) ([]amodels.Preview, error) {
	funcName = funcName + "/addTags"
	for _, article := range arts {
		outArticle := previewConv(article, authors[article.AuthorName])
		ChunkData = append(ChunkData, outArticle)
	}
	ChunkData, err := m.uploadTags(ChunkData, funcName)
	if err != nil {
		return ChunkData, sbErr.ErrDbError{
			Reason:   err.Error(),
			Function: funcName,
		}
	}
	if len(arts) < chunkSize {
		ChunkData = append(ChunkData, data.End)
	}
	return ChunkData, nil
}

func fullArticleConv(val amodels.DbArticle, Db *sqlx.DB, auth amodels.Author) (amodels.FullArticle, error) {
	fPath := "fullArticleConv"
	var article amodels.FullArticle
	article.Author = auth
	article.Comments = val.Comments
	article.DateTime = val.DateTime
	article.CommentsUrl = val.CommentsUrl
	article.Id = fmt.Sprint(val.Id)
	article.Likes = val.Likes
	article.PreviewUrl = val.PreviewUrl
	article.Title = val.Title
	article.Category = val.Category
	article.Text = val.Text
	rows, err := wrapper.MyQuery(Db, fPath, tagsLoad, val.Id)
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

func (m *psqlArticleRepository) authLimitChecker(schemaCount string, from, chunkSize int, args ...interface{}) (int, []amodels.Author, bool, error) {
	var ChunkData []amodels.Author
	path := "authLimitChecker"
	overCount := false
	var count int
	err := wrapper.MyGet(m.Db, path, schemaCount, &count, args...)
	if err != nil {
		return chunkSize, ChunkData, overCount, sbErr.ErrDbError{
			Reason:   err.Error(),
			Function: "articleRepository/limitChecker",
		}
	}

	if count <= from {
		ChunkData = append(ChunkData, models.Author{Login: "end"})
		overCount = true
	}

	if (count > from) && (count < from+chunkSize) {
		chunkSize = count - from
		overCount = true
	}
	return chunkSize, ChunkData, overCount, nil
}

func (m *psqlArticleRepository) addLiked(chunkData []amodels.Preview, login string) ([]amodels.Preview, error) {
	fName := "addLiked"
	schema := `select signum from article_likes where articleId = $1 and Login = $2`
	for i := range chunkData {
		err := wrapper.MyGet(m.Db, fName, schema, &chunkData[i].Liked, chunkData[i].Id, login)
		if err != nil {
			chunkData[i].Liked = 0
		}
	}
	return chunkData, nil
}

func (m *psqlArticleRepository) Fetch(ctx context.Context, login string, from, chunkSize int) (result []amodels.Preview, err error) {
	fName := toFetch
	var arts []amodels.DbArticle
	var ChunkData []amodels.Preview
	schema := "SELECT Id, PreviewUrl, DateTime,  Title, Category, Text, AuthorName,  CommentsUrl, Comments, Likes FROM ARTICLES WHERE Id < $1 ORDER BY Id DESC LIMIT $2;"
	err = wrapper.MySelect(m.Db, fName, schema, &arts, from, chunkSize)
	if err != nil {
		return ChunkData, sbErr.ErrDbError{
			Reason:   err.Error(),
			Function: fName,
		}
	}
	if len(arts) == 0 {
		ChunkData = append(ChunkData, data.End)
		return ChunkData, nil
	}
	var authors []string
	for _, a := range arts {
		authors = append(authors, a.AuthorName)
	}

	authorRes, err := m.uploadAuthors(authors, "articleRepository/Fetch")
	if err != nil {
		return ChunkData, sbErr.ErrDbError{
			Reason:   err.Error(),
			Function: fName,
		}
	}
	ChunkData, err = m.addTags(ChunkData, chunkSize, authorRes, fName, arts)
	if err != nil {
		return ChunkData, err
	}
	ChunkData, err = m.addLiked(ChunkData, login)
	for _, a := range ChunkData {
		fmt.Println(a.Liked)

	}
	return ChunkData, err
}

func (m *psqlArticleRepository) GetByID(ctx context.Context, login string, id int64) (result amodels.FullArticle, err error) {
	var newArticle amodels.DbArticle
	fName := "articleRepository/GetbyID"

	schema := "SELECT Id, PreviewUrl, DateTime, Title, Category, Text, AuthorName,  CommentsUrl, Comments, Likes FROM ARTICLES WHERE articles.Id = $1;"
	err = wrapper.MyGet(m.Db, fName, schema, &newArticle, id)
	var outArticle amodels.FullArticle
	if err != nil {
		return outArticle, sbErr.ErrDbError{
			Reason:   err.Error(),
			Function: fName,
		}
	}
	var newAuth amodels.Author
	schema = "SELECT AU.ID, AU.LOGIN, AU.NAME, AU.SURNAME, AU.AVATARURL, AU.DESCRIPTION, AU.EMAIL, AU.PASSWORD, AU.SCORE FROM ARTICLES AS AR INNER JOIN AUTHOR AS AU ON AU.LOGIN = AR.AuthorName WHERE AR.ID = $1;"
	err = wrapper.MyGet(m.Db, fName, schema, &newAuth, id)
	if err != nil {
		return outArticle, sbErr.ErrDbError{
			Reason:   err.Error(),
			Function: fName,
		}
	}
	outArticle, err = fullArticleConv(newArticle, m.Db, newAuth)
	if err != nil {
		return outArticle, sbErr.ErrDbError{
			Reason:   err.Error(),
			Function: fName,
		}
	}
	schema = `select signum from article_likes where articleId = $1 and Login = $2;`
	err = wrapper.MyGet(m.Db, fName, schema, &outArticle.Liked, id, login)
	if err != nil {
		outArticle.Liked = 0
	}
	return outArticle, nil
}

func (m *psqlArticleRepository) GetByTag(ctx context.Context, login string, tag string, from, chunkSize int) (result []amodels.Preview, err error) {
	fName := "articleRepository/GetbyTag"
	var arts []amodels.DbArticle
	var ChunkData []amodels.Preview
	schema := `select a.Id, a.PreviewUrl, a.DateTime, a.Title, Category, a.Text, a.AuthorName,  
	a.CommentsUrl, a.Comments, a.Likes from tags c
	inner join tags_articles ca  on c.Id = ca.tags_id
	inner join articles a on a.Id = ca.articles_id
	where c.tag = $1 and a.Id < $2 ORDER BY Id DESC LIMIT $3`
	err = wrapper.MySelect(m.Db, fName, schema, &arts, tag, from, chunkSize)
	if err != nil {
		return ChunkData, sbErr.ErrDbError{
			Reason:   err.Error(),
			Function: fName,
		}
	}
	if len(arts) == 0 {
		ChunkData = append(ChunkData, data.End)
		return ChunkData, nil
	}
	var authors []string
	for _, a := range arts {
		authors = append(authors, a.AuthorName)
	}

	authorRes, err := m.uploadAuthors(authors, fName)
	if err != nil {
		return ChunkData, sbErr.ErrDbError{
			Reason:   err.Error(),
			Function: fName,
		}
	}
	ChunkData, err = m.addTags(ChunkData, chunkSize, authorRes, fName, arts)
	if err != nil {
		return ChunkData, err
	}
	ChunkData, err = m.addLiked(ChunkData, login)
	return ChunkData, err
}

func (m *psqlArticleRepository) FindByTag(ctx context.Context, login string, query string, from, chunkSize int) (result []amodels.Preview, err error) {
	fName := "articleRepository/FindByTag"
	var arts []amodels.DbArticle
	var ChunkData []amodels.Preview
	schema := `select DISTINCT a.Id, a.PreviewUrl, a.DateTime, a.Title, 
	a.Category, a.Text, a.AuthorName,  a.CommentsUrl, a.Comments, a.Likes from tags c
	inner join tags_articles ca  on c.Id = ca.tags_id
	inner join articles a on a.Id = ca.articles_id
	where c.tag LIKE $1 and a.Id < $2 ORDER BY ID DESC LIMIT $3;`
	err = wrapper.MySelect(m.Db, fName, schema, &arts, query, from, chunkSize)
	if err != nil {
		return ChunkData, sbErr.ErrDbError{
			Reason:   err.Error(),
			Function: fName,
		}
	}
	var authors []string
	if len(arts) == 0 {
		ChunkData = append(ChunkData, data.End)
		return ChunkData, nil
	}
	for _, a := range arts {
		authors = append(authors, a.AuthorName)
	}

	authorRes, err := m.uploadAuthors(authors, "getByTag")
	if err != nil {
		return ChunkData, sbErr.ErrDbError{
			Reason:   err.Error(),
			Function: fName,
		}
	}
	ChunkData, err = m.addTags(ChunkData, chunkSize, authorRes, fName, arts)
	if err != nil {
		return ChunkData, err
	}
	ChunkData, err = m.addLiked(ChunkData, login)
	return ChunkData, err
}

func (m *psqlArticleRepository) GetByAuthor(ctx context.Context, login string, author string, from, chunkSize int) (result []amodels.Preview, err error) {
	fName := "articleRepository/GetByAuthor"
	var arts []amodels.DbArticle
	var ChunkData []amodels.Preview
	schema := "SELECT Id, PreviewUrl, DateTime, Title, Category, Text, AuthorName,  CommentsUrl, Comments, Likes FROM ARTICLES WHERE articles.AuthorName = $1 and articles.Id < $2 ORDER BY Id DESC LIMIT $3;"
	err = wrapper.MySelect(m.Db, fName, schema, &arts, author, from, chunkSize)
	if err != nil {
		return ChunkData, sbErr.ErrDbError{
			Reason:   err.Error(),
			Function: byAuthor,
		}
	}
	authors := []string{author}
	if len(arts) == 0 {
		ChunkData = append(ChunkData, data.End)
		return ChunkData, nil
	}
	if err != nil {
		return ChunkData, sbErr.ErrDbError{
			Reason:   err.Error(),
			Function: fName,
		}
	}

	authorRes, err := m.uploadAuthors(authors, fName)
	if err != nil {
		return ChunkData, sbErr.ErrDbError{
			Reason:   err.Error(),
			Function: fName,
		}
	}
	ChunkData, err = m.addTags(ChunkData, chunkSize, authorRes, fName, arts)
	if err != nil {
		return ChunkData, err
	}
	ChunkData, err = m.addLiked(ChunkData, login)
	return ChunkData, err
}

func (m *psqlArticleRepository) FindAuthors(ctx context.Context, query string, from, chunkSize int) (result []amodels.Author, err error) {
	fPath := "articleRepository/FindAuthors"
	query = "%" + query + "%"
	schemaCount := `SELECT count(*) FROM AUTHOR WHERE LOGIN LIKE $1 OR NAME LIKE $1 OR SURNAME LIKE $1;`
	chunkSize, ChunkData, overCount, err := m.authLimitChecker(schemaCount, from, chunkSize, query)
	if err != nil || len(ChunkData) > 0 {
		return ChunkData, err
	}
	schema := "SELECT AU.ID, AU.LOGIN, AU.NAME, AU.SURNAME, AU.AVATARURL, AU.DESCRIPTION, AU.EMAIL, AU.PASSWORD, AU.SCORE FROM AUTHOR AU WHERE LOGIN LIKE $1 OR NAME LIKE $1 OR SURNAME LIKE $1 ORDER BY AU.Id DESC LIMIT $2 OFFSET $3;"
	rows, err := wrapper.MyQuery(m.Db, fPath, schema, query, chunkSize, from)
	if err != nil {
		return ChunkData, sbErr.ErrDbError{
			Reason:   err.Error(),
			Function: fPath,
		}
	}
	var newAuthor models.Author
	for rows.Next() {
		err = rows.StructScan(&newAuthor)
		if err != nil {
			return ChunkData, sbErr.ErrDbError{
				Reason:   err.Error(),
				Function: fPath,
			}
		}
		ChunkData = append(ChunkData, newAuthor)
	}
	if overCount {
		ChunkData = append(ChunkData, models.Author{Login: "end"})
	}
	return ChunkData, err
}

func (m *psqlArticleRepository) FindArticles(ctx context.Context, login string, query string, from, chunkSize int) (result []amodels.Preview, err error) {
	fName := "articleRepository/FindAuthors"
	var arts []amodels.DbArticle
	var ChunkData []amodels.Preview
	schema := `SELECT Id, PreviewUrl, DateTime, Title, Category, Text, AuthorName,  CommentsUrl, 
	Comments, Likes FROM ARTICLES WHERE articles.Id < $1 and
	(en_tsvector(title, text) @@ plainto_tsquery('english', $2) or rus_tsvector(title, text) @@ plainto_tsquery('russian', $2)) 
	ORDER BY Id DESC LIMIT $3;`
	err = wrapper.MySelect(m.Db, fName, schema, &arts, from, query, chunkSize)
	if err != nil {
		return ChunkData, sbErr.ErrDbError{
			Reason:   err.Error(),
			Function: fName,
		}
	}
	var authors []string
	if len(arts) == 0 {
		ChunkData = append(ChunkData, data.End)
		return ChunkData, nil
	}
	for _, a := range arts {
		authors = append(authors, a.AuthorName)
	}

	authorRes, err := m.uploadAuthors(authors, fName)
	if err != nil {
		return ChunkData, sbErr.ErrDbError{
			Reason:   err.Error(),
			Function: fName,
		}
	}
	ChunkData, err = m.addTags(ChunkData, chunkSize, authorRes, fName, arts)
	if err != nil {
		return ChunkData, err
	}
	ChunkData, err = m.addLiked(ChunkData, login)
	return ChunkData, err
}

func (m *psqlArticleRepository) GetByCategory(ctx context.Context, login string, category string, from, chunkSize int) (result []amodels.Preview, err error) {
	fName := "articleRepository/GetByCategory"
	var arts []amodels.DbArticle
	var ChunkData []amodels.Preview
	schema := `SELECT Id, PreviewUrl, DateTime, Title, Category, Text, AuthorName,  
	CommentsUrl, Comments, Likes FROM ARTICLES WHERE articles.Category = $1
	and articles.Id < $2 ORDER BY Id DESC LIMIT $3`
	err = wrapper.MySelect(m.Db, fName, schema, &arts, category, from, chunkSize)
	if err != nil {
		return ChunkData, sbErr.ErrDbError{
			Reason:   err.Error(),
			Function: fName,
		}
	}
	if len(arts) == 0 {
		ChunkData = append(ChunkData, data.End)
		return ChunkData, nil
	}
	var authors []string
	for _, a := range arts {
		authors = append(authors, a.AuthorName)
	}

	authorRes, err := m.uploadAuthors(authors, fName)
	if err != nil {
		return ChunkData, sbErr.ErrDbError{
			Reason:   err.Error(),
			Function: fName,
		}
	}
	ChunkData, err = m.addTags(ChunkData, chunkSize, authorRes, fName, arts)
	if err != nil {
		return ChunkData, err
	}
	ChunkData, err = m.addLiked(ChunkData, login)
	return ChunkData, err
}

func (m *psqlArticleRepository) Store(ctx context.Context, a *amodels.Article) (int, error) {
	fPath := "store"
	insertArticle := `INSERT INTO articles (DateTime, PreviewUrl, Title, Category, Text, AuthorName, CommentsUrl, Comments, Likes) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING ID;`
	rows, err := wrapper.MyQuery(m.Db, fPath, insertArticle, a.DateTime, a.PreviewUrl, a.Title, a.Category, a.Text, a.AuthorName, a.CommentsUrl, a.Comments, a.Likes)
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

	insertCat := `INSERT INTO tags (tag) VALUES ($1) ON CONFLICT DO NOTHING;`
	for _, data := range a.Tags {
		_, err = wrapper.MyExec(m.Db, fPath, insertCat, data)
		fPath = "newtag"
		if err != nil {
			return Id, sbErr.ErrDbError{
				Reason:   err.Error(),
				Function: "articleRepository/Store",
			}
		}
	}
	insert_junc := `INSERT INTO tags_articles (articles_id, tags_id) VALUES 
	((SELECT Id FROM articles WHERE Id = $1) ,    
	(SELECT Id FROM tags WHERE tag = $2)) ON CONFLICT DO NOTHING;`
	for _, v := range a.Tags {
		_, err = wrapper.MyExec(m.Db, fPath, insert_junc, Id, v)
		fPath = "newconstraint"
		if err != nil {
			return Id, sbErr.ErrDbError{
				Reason:   err.Error(),
				Function: "articleRepository/Store",
			}
		}
	}
	return Id, nil
}

func (m *psqlArticleRepository) Delete(ctx context.Context, author string, id int64) error {
	fPath := "delete"
	schema := "DELETE FROM ARTICLES WHERE articles.Id = $1 and articles.Authorname = $2;"
	_, err := wrapper.MyExec(m.Db, fPath, schema, id, author)
	if err != nil {
		return sbErr.ErrDbError{
			Reason:   err.Error(),
			Function: "articleRepository/Delete",
		}
	}
	return nil
}
func (m *psqlArticleRepository) Update(ctx context.Context, a *amodels.Article) error {
	fPath := "update"
	uniqId, err := strconv.Atoi(a.Id)
	if err != nil {
		return sbErr.ErrDbError{
			Reason:   err.Error(),
			Function: "articleRepository/Update",
		}
	}
	updateArticle := `UPDATE articles SET DateTime = $1, Title = $2, Text = $3, PreviewUrl = $4, Category = $5  WHERE articles.Id  = $6 and articles.Authorname = $7;`
	_, err = wrapper.MyQuery(m.Db, fPath, updateArticle, time.Now().Format("2006/1/2 15:04"), a.Title, a.Text, a.PreviewUrl, a.Category, uniqId, a.AuthorName)
	if err != nil {
		return sbErr.ErrDbError{
			Reason:   err.Error(),
			Function: "articleRepository/Update",
		}
	}
	schema := deleteTags
	_, err = wrapper.MyExec(m.Db, fPath, schema, uniqId)
	if err != nil {
		return sbErr.ErrDbError{
			Reason:   err.Error(),
			Function: "articleRepository/Update",
		}
	}
	insertCat := `INSERT INTO tags (tag) VALUES ($1) ON CONFLICT DO NOTHING;`
	for _, data := range a.Tags {
		_, err = wrapper.MyExec(m.Db, fPath, insertCat, data)
		fPath = "newtag"
		if err != nil {
			return sbErr.ErrDbError{
				Reason:   err.Error(),
				Function: "articleRepository/Update",
			}
		}
	}
	insert_junc := `INSERT INTO tags_articles (articles_id, tags_id) VALUES
	((SELECT articles.Id FROM articles WHERE articles.Id = $1) ,
	(SELECT tags.Id FROM tags WHERE tags.tag = $2)) ON CONFLICT DO NOTHING;`
	for _, v := range a.Tags {
		_, err = wrapper.MyExec(m.Db, fPath, insert_junc, uniqId, v)
		fPath = "newconstraint"
		if err != nil {
			return sbErr.ErrDbError{
				Reason:   err.Error(),
				Function: "articleRepository/Store",
			}
		}
	}
	return nil
}
