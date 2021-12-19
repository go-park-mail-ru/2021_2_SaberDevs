package article

import (
	"context"
	"database/sql/driver"
	"regexp"
	"testing"

	amodels "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/article/models"
	"github.com/stretchr/testify/assert"
	sqlxmock "github.com/zhashkevych/go-sqlxmock"
)

func TestGetByID(t *testing.T) {
	db, mock, err := sqlxmock.Newx()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	rows := sqlxmock.NewRows([]string{"id", "previewurl", "datetime", "title", "category", "text", "authorname", "commentsurl", "comments", "likes"}).
		AddRow(1, "static/img/computer.png", "3", "7 Skills of Highly Effective Programmers", "SaberDevs",
			"Our team was inspired by the seven skills of highly effective", "mollenTEST1", "static/img/photo-elon-musk.jpg",
			97, 1001)

	query := "SELECT Id, PreviewUrl, DateTime, Title, Category, Text, AuthorName,  CommentsUrl, Comments, Likes FROM ARTICLES WHERE articles.Id = $1;"

	mock.ExpectQuery(regexp.QuoteMeta(query)).WithArgs(1).WillReturnRows(rows)

	rows2 := sqlxmock.NewRows([]string{"id", "login", "name", "surname", "avatarurl", "description", "email", "password", "score"}).
		AddRow(1, "mollenTEST1", "mollenTEST1", "7 Skills of Highly Effective Programmers", "SaberDevs",
			"Our team was inspired by the seven skills of highly effective", "mollenTEST1", "static/img/photo-elon-musk.jpg",
			97)

	query2 := "SELECT AU.ID, AU.LOGIN, AU.NAME, AU.SURNAME, AU.AVATARURL, AU.DESCRIPTION, AU.EMAIL, AU.PASSWORD, AU.SCORE FROM ARTICLES AS AR INNER JOIN AUTHOR AS AU ON AU.LOGIN = AR.AuthorName WHERE AR.ID = $1;"

	mock.ExpectQuery(regexp.QuoteMeta(query2)).WithArgs(1).WillReturnRows(rows2)
	rows0 := sqlxmock.NewRows([]string{"tag"}).
		AddRow("tag")

	query0 := "select c.tag from tags c inner join tags_articles ca on c.Id = ca.tags_id inner join articles a on a.Id = ca.articles_id where a.Id = $1;"

	mock.ExpectQuery(regexp.QuoteMeta(query0)).WithArgs(1).WillReturnRows(rows0)

	rows3 := sqlxmock.NewRows([]string{"signum"}).
		AddRow(1)

	query3 := "select signum from article_likes where articleId = $1 and Login = $2;"

	mock.ExpectQuery(regexp.QuoteMeta(query3)).WithArgs(1, "mollen").WillReturnRows(rows3)

	a := NewArticleRepository(db)

	anArticle, err := a.GetByID(context.TODO(), "mollen", 1)
	assert.NoError(t, err)
	assert.NotNil(t, anArticle)
}

func TestGetbyTag(t *testing.T) {
	db, mock, err := sqlxmock.Newx()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	tag := "tag"
	from := 0
	chunkSize := 1
	login := "mollenTEST1"
	id := "1"
	rows := sqlxmock.NewRows([]string{"id", "previewurl", "datetime", "title", "category", "text", "authorname", "commentsurl", "comments", "likes"}).
		AddRow(1, "static/img/computer.png", "3", "7 Skills of Highly Effective Programmers", "SaberDevs",
			"Our team was inspired by the seven skills of highly effective", "mollenTEST1", "static/img/photo-elon-musk.jpg",
			97, 1001)

	query := "select a.Id, a.PreviewUrl, a.DateTime, a.Title, Category, a.Text, a.AuthorName, " +
		"a.CommentsUrl, a.Comments, a.Likes from tags c " +
		"inner join tags_articles ca  on c.Id = ca.tags_id " +
		"inner join articles a on a.Id = ca.articles_id " +
		"where c.tag = $1 and a.Id < $2 ORDER BY Id DESC LIMIT $3 "

	mock.ExpectQuery(regexp.QuoteMeta(query)).WithArgs(tag, from, chunkSize).WillReturnRows(rows)

	rows2 := sqlxmock.NewRows([]string{"id", "login", "name", "surname", "avatarurl", "description", "email", "password", "score"}).
		AddRow(1, "mollenTEST1", "mollenTEST1", "7 Skills of Highly Effective Programmers", "SaberDevs",
			"Our team was inspired by the seven skills of highly effective", "mollenTEST1", "static/img/photo-elon-musk.jpg",
			97)

	query2 := "SELECT AU.ID, AU.LOGIN, AU.NAME, AU.SURNAME, AU.AVATARURL, AU.DESCRIPTION, AU.EMAIL, AU.PASSWORD, AU.SCORE FROM AUTHOR AU WHERE AU.LOGIN IN ($1);"

	mock.ExpectQuery(regexp.QuoteMeta(query2)).WithArgs(login).WillReturnRows(rows2)
	rows0 := sqlxmock.NewRows([]string{"id", "tag"}).
		AddRow(1, "tag")

	query0 := "select a.Id, c.tag from tags c inner join tags_articles ca on c.Id = ca.tags_id inner join articles a on a.Id = ca.articles_id where a.Id in ($1) order by a.Id DESC;"

	mock.ExpectQuery(regexp.QuoteMeta(query0)).WithArgs(id).WillReturnRows(rows0)

	rows3 := sqlxmock.NewRows([]string{"signum"}).
		AddRow(1)

	query3 := "select signum from article_likes where articleId = $1 and Login = $2;"

	mock.ExpectQuery(regexp.QuoteMeta(query3)).WithArgs(1, login).WillReturnRows(rows3)

	a := NewArticleRepository(db)

	anArticle, err := a.GetByTag(context.TODO(), login, tag, from, chunkSize)
	assert.NoError(t, err)
	assert.NotNil(t, anArticle)
}

func TestFetch(t *testing.T) {
	db, mock, err := sqlxmock.Newx()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	from := 0
	chunkSize := 1
	login := "mollenTEST1"
	id := "1"
	rows := sqlxmock.NewRows([]string{"id", "previewurl", "datetime", "title", "category", "text", "authorname", "commentsurl", "comments", "likes"}).
		AddRow(1, "static/img/computer.png", "3", "7 Skills of Highly Effective Programmers", "SaberDevs",
			"Our team was inspired by the seven skills of highly effective", "mollenTEST1", "static/img/photo-elon-musk.jpg",
			97, 1001)

	query := "SELECT Id, PreviewUrl, DateTime,  Title, Category, Text, AuthorName,  CommentsUrl, Comments, Likes FROM ARTICLES WHERE Id < $1 ORDER BY Id DESC LIMIT $2;"

	mock.ExpectQuery(regexp.QuoteMeta(query)).WithArgs(from, chunkSize).WillReturnRows(rows)

	rows2 := sqlxmock.NewRows([]string{"id", "login", "name", "surname", "avatarurl", "description", "email", "password", "score"}).
		AddRow(1, "mollenTEST1", "mollenTEST1", "7 Skills of Highly Effective Programmers", "SaberDevs",
			"Our team was inspired by the seven skills of highly effective", "mollenTEST1", "static/img/photo-elon-musk.jpg",
			97)

	query2 := "SELECT AU.ID, AU.LOGIN, AU.NAME, AU.SURNAME, AU.AVATARURL, AU.DESCRIPTION, AU.EMAIL, AU.PASSWORD, AU.SCORE FROM AUTHOR AU WHERE AU.LOGIN IN ($1);"

	mock.ExpectQuery(regexp.QuoteMeta(query2)).WithArgs(login).WillReturnRows(rows2)
	rows0 := sqlxmock.NewRows([]string{"id", "tag"}).
		AddRow(1, "tag")

	query0 := "select a.Id, c.tag from tags c inner join tags_articles ca on c.Id = ca.tags_id inner join articles a on a.Id = ca.articles_id where a.Id in ($1) order by a.Id DESC;"

	mock.ExpectQuery(regexp.QuoteMeta(query0)).WithArgs(id).WillReturnRows(rows0)

	rows3 := sqlxmock.NewRows([]string{"signum"}).
		AddRow(1)

	query3 := "select signum from article_likes where articleId = $1 and Login = $2;"

	mock.ExpectQuery(regexp.QuoteMeta(query3)).WithArgs(1, login).WillReturnRows(rows3)

	a := NewArticleRepository(db)

	anArticle, err := a.Fetch(context.TODO(), login, from, chunkSize)
	assert.NoError(t, err)
	assert.NotNil(t, anArticle)
}

func TestFindByTag(t *testing.T) {
	db, mock, err := sqlxmock.Newx()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	tag := "tag"
	from := 0
	chunkSize := 1
	login := "mollenTEST1"
	id := "1"
	rows := sqlxmock.NewRows([]string{"id", "previewurl", "datetime", "title", "category", "text", "authorname", "commentsurl", "comments", "likes"}).
		AddRow(1, "static/img/computer.png", "3", "7 Skills of Highly Effective Programmers", "SaberDevs",
			"Our team was inspired by the seven skills of highly effective", "mollenTEST1", "static/img/photo-elon-musk.jpg",
			97, 1001)

	query := "select DISTINCT a.Id, a.PreviewUrl, a.DateTime, a.Title, " +
		"a.Category, a.Text, a.AuthorName,  a.CommentsUrl, a.Comments, a.Likes from tags c " +
		"inner join tags_articles ca  on c.Id = ca.tags_id " +
		"inner join articles a on a.Id = ca.articles_id " +
		"where c.tag LIKE $1 and a.Id < $2 ORDER BY ID DESC LIMIT $3;"

	mock.ExpectQuery(regexp.QuoteMeta(query)).WithArgs(tag, from, chunkSize).WillReturnRows(rows)

	rows2 := sqlxmock.NewRows([]string{"id", "login", "name", "surname", "avatarurl", "description", "email", "password", "score"}).
		AddRow(1, "mollenTEST1", "mollenTEST1", "7 Skills of Highly Effective Programmers", "SaberDevs",
			"Our team was inspired by the seven skills of highly effective", "mollenTEST1", "static/img/photo-elon-musk.jpg",
			97)

	query2 := "SELECT AU.ID, AU.LOGIN, AU.NAME, AU.SURNAME, AU.AVATARURL, AU.DESCRIPTION, AU.EMAIL, AU.PASSWORD, AU.SCORE FROM AUTHOR AU WHERE AU.LOGIN IN ($1);"

	mock.ExpectQuery(regexp.QuoteMeta(query2)).WithArgs(login).WillReturnRows(rows2)
	rows0 := sqlxmock.NewRows([]string{"id", "tag"}).
		AddRow(1, "tag")

	query0 := "select a.Id, c.tag from tags c inner join tags_articles ca on c.Id = ca.tags_id inner join articles a on a.Id = ca.articles_id where a.Id in ($1) order by a.Id DESC;"

	mock.ExpectQuery(regexp.QuoteMeta(query0)).WithArgs(id).WillReturnRows(rows0)

	rows3 := sqlxmock.NewRows([]string{"signum"}).
		AddRow(1)

	query3 := "select signum from article_likes where articleId = $1 and Login = $2;"

	mock.ExpectQuery(regexp.QuoteMeta(query3)).WithArgs(1, login).WillReturnRows(rows3)

	a := NewArticleRepository(db)

	anArticle, err := a.FindByTag(context.TODO(), login, tag, from, chunkSize)
	assert.NoError(t, err)
	assert.NotNil(t, anArticle)
}

func TestGetbyAuthor(t *testing.T) {
	db, mock, err := sqlxmock.Newx()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	author := "author"
	from := 0
	chunkSize := 1
	login := "mollenTEST1"
	id := "1"
	rows := sqlxmock.NewRows([]string{"id", "previewurl", "datetime", "title", "category", "text", "authorname", "commentsurl", "comments", "likes"}).
		AddRow(1, "static/img/computer.png", "3", "7 Skills of Highly Effective Programmers", "SaberDevs",
			"Our team was inspired by the seven skills of highly effective", "mollenTEST1", "static/img/photo-elon-musk.jpg",
			97, 1001)

	query := "SELECT Id, PreviewUrl, DateTime, Title, Category, Text, AuthorName,  CommentsUrl, Comments, Likes FROM ARTICLES WHERE articles.AuthorName = $1 and articles.Id < $2 ORDER BY Id DESC LIMIT $3;"

	mock.ExpectQuery(regexp.QuoteMeta(query)).WithArgs(author, from, chunkSize).WillReturnRows(rows)

	rows2 := sqlxmock.NewRows([]string{"id", "login", "name", "surname", "avatarurl", "description", "email", "password", "score"}).
		AddRow(1, "mollenTEST1", "mollenTEST1", "7 Skills of Highly Effective Programmers", "SaberDevs",
			"Our team was inspired by the seven skills of highly effective", "mollenTEST1", "static/img/photo-elon-musk.jpg",
			97)

	query2 := "SELECT AU.ID, AU.LOGIN, AU.NAME, AU.SURNAME, AU.AVATARURL, AU.DESCRIPTION, AU.EMAIL, AU.PASSWORD, AU.SCORE FROM AUTHOR AU WHERE AU.LOGIN IN ($1);"

	mock.ExpectQuery(regexp.QuoteMeta(query2)).WithArgs(author).WillReturnRows(rows2)
	rows0 := sqlxmock.NewRows([]string{"id", "tag"}).
		AddRow(1, "tag")

	query0 := "select a.Id, c.tag from tags c inner join tags_articles ca on c.Id = ca.tags_id inner join articles a on a.Id = ca.articles_id where a.Id in ($1) order by a.Id DESC;"

	mock.ExpectQuery(regexp.QuoteMeta(query0)).WithArgs(id).WillReturnRows(rows0)

	rows3 := sqlxmock.NewRows([]string{"signum"}).
		AddRow(1)

	query3 := "select signum from article_likes where articleId = $1 and Login = $2;"

	mock.ExpectQuery(regexp.QuoteMeta(query3)).WithArgs(1, login).WillReturnRows(rows3)

	a := NewArticleRepository(db)

	anArticle, err := a.GetByAuthor(context.TODO(), login, author, from, chunkSize)
	assert.NoError(t, err)
	assert.NotNil(t, anArticle)
}

func TestFindAuthors(t *testing.T) {
	db, mock, err := sqlxmock.Newx()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	qauthor := "%author%"
	author := "author"
	from := 0
	chunkSize := 1
	login := "mollenTEST1"
	id := "1"

	rowspre := sqlxmock.NewRows([]string{"count(*)"}).
		AddRow(1)

	querypre := "SELECT count(*) FROM AUTHOR WHERE LOGIN LIKE $1 OR NAME LIKE $1 OR SURNAME LIKE $1;"

	mock.ExpectQuery(regexp.QuoteMeta(querypre)).WithArgs(qauthor).WillReturnRows(rowspre)

	rows2 := sqlxmock.NewRows([]string{"id", "login", "name", "surname", "avatarurl", "description", "email", "password", "score"}).
		AddRow(1, "mollenTEST1", "mollenTEST1", "7 Skills of Highly Effective Programmers", "SaberDevs",
			"Our team was inspired by the seven skills of highly effective", "mollenTEST1", "static/img/photo-elon-musk.jpg",
			97)

	query2 := "SELECT AU.ID, AU.LOGIN, AU.NAME, AU.SURNAME, AU.AVATARURL, AU.DESCRIPTION, AU.EMAIL, AU.PASSWORD, AU.SCORE FROM AUTHOR AU WHERE LOGIN LIKE $1 OR NAME LIKE $1 OR SURNAME LIKE $1 ORDER BY AU.Id DESC LIMIT $2 OFFSET $3;"

	mock.ExpectQuery(regexp.QuoteMeta(query2)).WithArgs(qauthor, chunkSize, from).WillReturnRows(rows2)
	rows0 := sqlxmock.NewRows([]string{"id", "tag"}).
		AddRow(1, "tag")

	query0 := "select a.Id, c.tag from tags c inner join tags_articles ca on c.Id = ca.tags_id inner join articles a on a.Id = ca.articles_id where a.Id in ($1) order by a.Id DESC;"

	mock.ExpectQuery(regexp.QuoteMeta(query0)).WithArgs(id).WillReturnRows(rows0)

	rows3 := sqlxmock.NewRows([]string{"signum"}).
		AddRow(1)

	query3 := "select signum from article_likes where articleId = $1 and Login = $2;"

	mock.ExpectQuery(regexp.QuoteMeta(query3)).WithArgs(1, login).WillReturnRows(rows3)

	a := NewArticleRepository(db)

	anAuth, err := a.FindAuthors(context.TODO(), author, from, chunkSize)
	assert.NoError(t, err)
	assert.NotNil(t, anAuth)
}

func TestFindArticles(t *testing.T) {
	db, mock, err := sqlxmock.Newx()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	qart := "author"
	from := 0
	chunkSize := 1
	login := "mollenTEST1"
	id := "1"

	rows := sqlxmock.NewRows([]string{"id", "previewurl", "datetime", "title", "category", "text", "authorname", "commentsurl", "comments", "likes"}).
		AddRow(1, "static/img/computer.png", "3", "7 Skills of Highly Effective Programmers", "SaberDevs",
			"Our team was inspired by the seven skills of highly effective", "mollenTEST1", "static/img/photo-elon-musk.jpg",
			97, 1001)

	query := "SELECT Id, PreviewUrl, DateTime, Title, Category, Text, AuthorName,  CommentsUrl, " +
		"Comments, Likes FROM ARTICLES WHERE articles.Id < $1 and " +
		"(en_tsvector(title, text) @@ plainto_tsquery('english', $2) or rus_tsvector(title, text) @@ plainto_tsquery('russian', $2)) " +
		"ORDER BY Id DESC LIMIT $3;"

	mock.ExpectQuery(regexp.QuoteMeta(query)).WithArgs(from, qart, chunkSize).WillReturnRows(rows)

	rows2 := sqlxmock.NewRows([]string{"id", "login", "name", "surname", "avatarurl", "description", "email", "password", "score"}).
		AddRow(1, "mollenTEST1", "mollenTEST1", "7 Skills of Highly Effective Programmers", "SaberDevs",
			"Our team was inspired by the seven skills of highly effective", "mollenTEST1", "static/img/photo-elon-musk.jpg",
			97)

	query2 := "SELECT AU.ID, AU.LOGIN, AU.NAME, AU.SURNAME, AU.AVATARURL, AU.DESCRIPTION, AU.EMAIL, AU.PASSWORD, AU.SCORE FROM AUTHOR AU WHERE AU.LOGIN IN ($1);"

	mock.ExpectQuery(regexp.QuoteMeta(query2)).WithArgs(login).WillReturnRows(rows2)
	rows0 := sqlxmock.NewRows([]string{"id", "tag"}).
		AddRow(1, "tag")

	query0 := "select a.Id, c.tag from tags c inner join tags_articles ca on c.Id = ca.tags_id inner join articles a on a.Id = ca.articles_id where a.Id in ($1) order by a.Id DESC;"

	mock.ExpectQuery(regexp.QuoteMeta(query0)).WithArgs(id).WillReturnRows(rows0)

	rows3 := sqlxmock.NewRows([]string{"signum"}).
		AddRow(1)

	query3 := "select signum from article_likes where articleId = $1 and Login = $2;"

	mock.ExpectQuery(regexp.QuoteMeta(query3)).WithArgs(1, login).WillReturnRows(rows3)

	a := NewArticleRepository(db)

	anArticle, err := a.FindArticles(context.TODO(), login, qart, from, chunkSize)
	assert.NoError(t, err)
	assert.NotNil(t, anArticle)
}

func TestGetbyCategory(t *testing.T) {
	db, mock, err := sqlxmock.Newx()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	category := "author"
	from := 0
	chunkSize := 1
	login := "mollenTEST1"
	id := "1"
	rows := sqlxmock.NewRows([]string{"id", "previewurl", "datetime", "title", "category", "text", "authorname", "commentsurl", "comments", "likes"}).
		AddRow(1, "static/img/computer.png", "3", "7 Skills of Highly Effective Programmers", "SaberDevs",
			"Our team was inspired by the seven skills of highly effective", "mollenTEST1", "static/img/photo-elon-musk.jpg",
			97, 1001)

	query := "SELECT Id, PreviewUrl, DateTime, Title, Category, Text, AuthorName, " +
		"CommentsUrl, Comments, Likes FROM ARTICLES WHERE articles.Category = $1 " +
		"and articles.Id < $2 ORDER BY Id DESC LIMIT $3"

	mock.ExpectQuery(regexp.QuoteMeta(query)).WithArgs(category, from, chunkSize).WillReturnRows(rows)

	rows2 := sqlxmock.NewRows([]string{"id", "login", "name", "surname", "avatarurl", "description", "email", "password", "score"}).
		AddRow(1, "mollenTEST1", "mollenTEST1", "7 Skills of Highly Effective Programmers", "SaberDevs",
			"Our team was inspired by the seven skills of highly effective", "mollenTEST1", "static/img/photo-elon-musk.jpg",
			97)

	query2 := "SELECT AU.ID, AU.LOGIN, AU.NAME, AU.SURNAME, AU.AVATARURL, AU.DESCRIPTION, AU.EMAIL, AU.PASSWORD, AU.SCORE FROM AUTHOR AU WHERE AU.LOGIN IN ($1);"

	mock.ExpectQuery(regexp.QuoteMeta(query2)).WithArgs(login).WillReturnRows(rows2)
	rows0 := sqlxmock.NewRows([]string{"id", "tag"}).
		AddRow(1, "tag")

	query0 := "select a.Id, c.tag from tags c inner join tags_articles ca on c.Id = ca.tags_id inner join articles a on a.Id = ca.articles_id where a.Id in ($1) order by a.Id DESC;"

	mock.ExpectQuery(regexp.QuoteMeta(query0)).WithArgs(id).WillReturnRows(rows0)

	rows3 := sqlxmock.NewRows([]string{"signum"}).
		AddRow(1)

	query3 := "select signum from article_likes where articleId = $1 and Login = $2;"

	mock.ExpectQuery(regexp.QuoteMeta(query3)).WithArgs(1, login).WillReturnRows(rows3)

	a := NewArticleRepository(db)

	anArticle, err := a.GetByCategory(context.TODO(), login, category, from, chunkSize)
	assert.NoError(t, err)
	assert.NotNil(t, anArticle)
}

func TestStore(t *testing.T) {
	db, mock, err := sqlxmock.Newx()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	login := "mollenTEST1"
	id := "1"
	tid := int64(1)
	tag := "tag"
	art := amodels.Article{Id: id, PreviewUrl: "123", AuthorName: login, Tags: []string{tag}}

	rows := sqlxmock.NewRows([]string{"id"}).
		AddRow(1)

	query := "INSERT INTO articles (DateTime, PreviewUrl, Title, Category, Text, AuthorName, CommentsUrl, Comments, Likes) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING ID;"
	mock.ExpectQuery(regexp.QuoteMeta(query)).WithArgs(sqlxmock.AnyArg(), art.PreviewUrl, art.Title, art.Category, art.Text, art.AuthorName, art.CommentsUrl, art.Comments, art.Likes).WillReturnRows(rows)

	query2 := "INSERT INTO tags (tag) VALUES ($1) ON CONFLICT DO NOTHING;"
	mock.ExpectExec(regexp.QuoteMeta(query2)).WithArgs(tag).WillReturnResult(driver.RowsAffected(1))

	query3 := "INSERT INTO tags_articles (articles_id, tags_id) VALUES ((SELECT Id FROM articles WHERE Id = $1) , (SELECT Id FROM tags WHERE tag = $2)) ON CONFLICT DO NOTHING;"

	mock.ExpectExec(regexp.QuoteMeta(query3)).WithArgs(tid, tag).WillReturnResult(driver.RowsAffected(1))
	a := NewArticleRepository(db)
	aid, err := a.Store(context.TODO(), &art)
	assert.NoError(t, err)
	assert.NotNil(t, aid)
	assert.Equal(t, aid, 1)
}

func TestUpdate(t *testing.T) {
	db, mock, err := sqlxmock.Newx()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	login := "mollenTEST1"
	id := "1"
	tid := int64(1)
	tag := "tag"
	art := amodels.Article{Id: id, PreviewUrl: "123", AuthorName: login, Tags: []string{tag}}

	rows := sqlxmock.NewRows([]string{""})
	query := "UPDATE articles SET DateTime = $1, Title = $2, Text = $3, PreviewUrl = $4, Category = $5  WHERE articles.Id  = $6 and articles.Authorname = $7;"
	mock.ExpectQuery(regexp.QuoteMeta(query)).WithArgs(sqlxmock.AnyArg(), art.Title, art.Text, art.PreviewUrl, art.Category, tid, art.AuthorName).WillReturnRows(rows)

	query0 := "delete from tags_articles ta where ta.articles_id = $1;"
	mock.ExpectExec(regexp.QuoteMeta(query0)).WithArgs(tid).WillReturnResult(driver.RowsAffected(1))

	query2 := "INSERT INTO tags (tag) VALUES ($1) ON CONFLICT DO NOTHING;"
	mock.ExpectExec(regexp.QuoteMeta(query2)).WithArgs(tag).WillReturnResult(driver.RowsAffected(1))

	query3 := "INSERT INTO tags_articles (articles_id, tags_id) VALUES ((SELECT articles.Id FROM articles WHERE articles.Id = $1) , (SELECT tags.Id FROM tags WHERE tags.tag = $2)) ON CONFLICT DO NOTHING;"

	mock.ExpectExec(regexp.QuoteMeta(query3)).WithArgs(tid, tag).WillReturnResult(driver.RowsAffected(1))
	a := NewArticleRepository(db)
	err = a.Update(context.TODO(), &art)
	assert.NoError(t, err)
}
func TestDelete(t *testing.T) {
	db, mock, err := sqlxmock.Newx()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	login := "mollenTEST1"
	tid := int64(1)

	query0 := "DELETE FROM ARTICLES WHERE articles.Id = $1 and articles.Authorname = $2;"
	mock.ExpectExec(regexp.QuoteMeta(query0)).WithArgs(tid, login).WillReturnResult(driver.RowsAffected(1))

	a := NewArticleRepository(db)
	err = a.Delete(context.TODO(), login, tid)
	assert.NoError(t, err)
}
