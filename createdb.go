package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"time"

	server "github.com/go-park-mail-ru/2021_2_SaberDevs/cmd/sybernews"
	amodels "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/article/models"
	repo "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/article/repository"
	"github.com/go-park-mail-ru/2021_2_SaberDevs/internal/data"
	dataDB "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/data"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func Testing() {
	connStr, err := server.DbConfig()
	if err != nil {
		log.Fatal(err)
	}
	db, err := sqlx.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.Queryx("SELECT ID, LOGIN, NAME, SURNAME, EMAIL, PASSWORD, SCORE FROM author")
	if err != nil {
		fmt.Println(err.Error())
	}
	var names []string
	var author amodels.Author
	for rows.Next() {
		err = rows.StructScan(&author)
		if err != nil {
			fmt.Println(err.Error())
		}
		names = append(names, author.Login)
		fmt.Println(author.Name)
	}

	rows, err = db.Queryx("SELECT Id, PreviewUrl, DateTime,  Title, Text, AuthorName,  CommentsUrl, Comments, Likes FROM ARTICLES")
	if err != nil {
		log.Fatal(err)
	}
	var newArticle amodels.DbArticle
	for rows.Next() {
		err = rows.StructScan(&newArticle)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Print(newArticle.Id, "  ", newArticle.DateTime, "  ", newArticle.PreviewUrl, "  ", newArticle.AuthorName, "  ", newArticle.Likes, "\n")
	}
	fmt.Print("whereami", "\n")
	rows, err = db.Queryx("SELECT * FROM tags;")
	if err != nil {
		fmt.Println(err.Error())
	}
	var mytag string
	var tagid int
	for rows.Next() {
		err = rows.Scan(&tagid, &mytag)
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Print(tagid, "  ", mytag, "\n")
	}
	fmt.Print("whereami", "\n")
	rows, err = db.Queryx("SELECT * FROM categories;")
	if err != nil {
		fmt.Println(err.Error())
	}
	for rows.Next() {
		err = rows.Scan(&mytag)
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Print("  ", mytag, "\n")
	}
	fmt.Print("whereami", "\n")

	rows, err = db.Queryx("SELECT * FROM tags_articles;")
	if err != nil {
		fmt.Println(err.Error())
	}
	var tag amodels.СategoriesArticles
	for rows.Next() {
		err = rows.Scan(&tag.Articles_id, &tag.Tags_id)
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Print(tag.Articles_id, "  ", tag.Tags_id, "\n")
	}

	rows, err = db.Queryx(`select c.tag from tags c
	inner join tags_articles ca  on c.Id = ca.tags_id
	inner join articles a on a.Id = ca.articles_id
	where a.Id = $1;`, 11)
	if err != nil {
		fmt.Println(err.Error())
	}

	for rows.Next() {
		err = rows.Scan(&mytag)
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Printf("%s\n", mytag)
	}

	rows, err = db.Queryx("SELECT count(*) FROM articles;")
	if err != nil {
		fmt.Println(err.Error())
	}
	var count int
	for rows.Next() {
		err = rows.Scan(&count)
		if err != nil {
			fmt.Println(err.Error())
		}
	}
	fmt.Println("!", count)
	myRepo := repo.NewArticleRepository(db)
	result, err := myRepo.GetByID(context.TODO(), 10)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Print(result.Id, " ", result.Author.Name, " ", result.Tags, " ", result.Text, " ", result.Likes, "\n")
	fmt.Println()
	results, err := myRepo.GetByAuthor(context.TODO(), "dar", 0, 5)
	if err != nil {
		fmt.Println(err.Error())
	}
	for _, result := range results {
		fmt.Print(result.Id, " ", result.Author.Name, " ", result.Tags, " ", result.Text, " ", result.Likes, "\n")
	}
	fmt.Println()
	results, err = myRepo.GetByCategory(context.TODO(), "Маркетинг", 0, 5)
	if err != nil {
		fmt.Println(err.Error())
	}
	for _, result := range results {
		fmt.Print(result.Id, " ", result.Category, " ", result.Author.Name, " ", result.Tags, " ", result.Text, " ", result.Likes, "\n")
	}
	fmt.Println("fgh")
	results, err = myRepo.GetByCategory(context.TODO(), "Марке", 0, 5)
	if err != nil {
		fmt.Println(err.Error())
	}
	for _, result := range results {
		fmt.Print(result.Id, " ", result.Category, " ", result.Author.Name, " ", result.Tags, " ", result.Text, " ", result.Likes, "\n")
	}
	fmt.Println("fgh")
	fmt.Println()
	results, err = myRepo.GetByTag(context.TODO(), "finance", 0, 5)
	if err != nil {
		fmt.Println(err.Error())
	}
	for _, result := range results {
		fmt.Print(result.Id, " ", result.Author.Name, " ", result.Tags, " ", result.Text, " ", result.Likes, "\n")
	}

	fmt.Println()
	a := data.TestData[3]
	var ar amodels.Article
	ar.AuthorAvatar = a.AuthorAvatar
	ar.AuthorUrl = a.AuthorUrl
	ar.Comments = a.Comments
	ar.CommentsUrl = a.CommentsUrl
	ar.DateTime = time.Now().Format("2006/1/2 15:04")
	ar.Likes = a.Likes
	ar.PreviewUrl = a.PreviewUrl
	ar.Text = a.Text
	ar.Title = a.Title
	ar.Tags = a.Tags
	ar.Id = "13"
	ar.AuthorName = "dar"
	ar.Category = "SaberDevs"
	ar.Tags = append(ar.Tags, "finance")
	Id, err := myRepo.Store(context.TODO(), &ar)
	ar.Id = fmt.Sprint(Id)
	fmt.Println("IDDDDD=", Id)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("!", count)
	result, err = myRepo.GetByID(context.TODO(), int64(Id))
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Print(result.Id, " ", result.Category, " ", result.Author.Name, " ", result.Tags, " ", result.Text, " ", result.Likes, "\n")

	err = myRepo.Delete(context.TODO(), 3)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println()
	newresult, err := myRepo.Fetch(context.TODO(), 0, 12)
	if err != nil {
		fmt.Println(err.Error())
	}
	for _, result := range newresult {
		fmt.Print(result.Id, " ", result.Category, " ", result.Author.Name, " ", result.Tags, " ", result.Text, " ", result.Likes, "\n")
	}
	newresult, err = myRepo.Fetch(context.TODO(), 12, 7)
	if err != nil {
		fmt.Println(err.Error())
	}
	for _, result := range newresult {
		fmt.Print(result.Id, " ", result.Author.Name, " ", result.Tags, " ", result.Text, " ", result.Likes, "\n")
	}
	fmt.Println()
	fmt.Println()
	// ar.Text = `<a onblur="alert(secret)" href="http://www.google.com">Google</a>`
	ar.Tags = []string{"jojo", "finance"}
	err = myRepo.Update(context.TODO(), &ar)
	if err != nil {
		fmt.Println(err.Error())
	}
	result, err = myRepo.GetByID(context.TODO(), int64(Id))
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Print(result.Id, " ", result.Category, " ", result.Author.Name, " ", result.Tags, " ", result.Text, " ", result.Likes, "\n")

	fmt.Println("эксплойт для")
	newresult, err = myRepo.FindArticles(context.TODO(), "конкурс", 0, 4)
	if err != nil {
		fmt.Println(err.Error())
	}
	for _, result := range newresult {
		fmt.Print(result.Id, " ", result.Title, " ", result.Category, " ", result.Author.Name, " ", result.Tags, " ", result.Text, " ", result.Likes, "\n")
	}
	fmt.Println("Progra")
	newresult, err = myRepo.FindArticles(context.TODO(), "7 skill high", 0, 5)
	if err != nil {
		fmt.Println(err.Error())
	}
	for _, result := range newresult {
		fmt.Print(result.Id, " ", result.Title, " ", result.Category, " ", result.Author.Name, " ", result.Tags, " ", result.Text, " ", result.Likes, "\n")
	}

	fmt.Println("Program")
	newresult, err = myRepo.FindArticles(context.TODO(), "programm", 0, 15)
	if err != nil {
		fmt.Println(err.Error())
	}
	for _, result := range newresult {
		fmt.Print(result.Id, " ", result.Title, " ", result.Category, " ", result.Author.Name, " ", result.Tags, " ", result.Text, " ", result.Likes, "\n")
	}

	fmt.Println()
	newresult, err = myRepo.FindByTag(context.TODO(), "", 0, 10)
	if err != nil {
		fmt.Println(err.Error())
	}
	for _, result := range newresult {
		fmt.Print(result.Id, " ", result.Title, " ", result.Category, " ", result.Author.Name, " ", result.Tags, " ", result.Text, " ", result.Likes, "\n")
	}

	fmt.Println()

	newresult, err = myRepo.FindByTag(context.TODO(), "in", 12, 10)
	if err != nil {
		fmt.Println(err.Error())
	}
	for _, result := range newresult {
		fmt.Print(result.Id, " ", result.Title, " ", result.Category, " ", result.Author.Name, " ", result.Tags, " ", result.Text, " ", result.Likes, "\n")
	}

	fmt.Println()
	authr, err := myRepo.FindAuthors(context.TODO(), "en", 0, 10)
	if err != nil {
		fmt.Println(err.Error())
	}
	for _, result := range authr {
		fmt.Print(result.Id, " ", result.Login, " ", result.Name, " ", result.Surname, " ", "\n")
	}

	fmt.Println()
	authr, err = myRepo.FindAuthors(context.TODO(), "en", 10, 10)
	if err != nil {
		fmt.Println(err.Error())
	}
	for _, result := range authr {
		fmt.Print(result.Id, " ", result.Login, " ", result.Name, " ", result.Surname, " ", "\n")
	}

	fmt.Println()
	results, err = myRepo.GetByCategory(context.TODO(), "Офис", 0, 10)
	if err != nil {
		fmt.Println(err.Error())
	}
	for _, result := range results {
		fmt.Print(result.Id, " ", result.Title, " ", result.Category, " ", result.Author.Name, " ", result.Tags, " ", result.Text, " ", result.Likes, "\n")
	}
}

func main() {
	connStr, err := server.DbConfig()
	if err != nil {
		log.Fatal(err)
	}
	db, err := sqlx.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	schema := `DROP TABLE IF EXISTS articles CASCADE;
		DROP TABLE IF EXISTS author CASCADE;
		DROP TABLE IF EXISTS tags CASCADE;
		DROP TABLE IF EXISTS categories CASCADE;
		DROP TABLE IF EXISTS categories_articles CASCADE;
		DROP TABLE IF EXISTS tags_articles CASCADE;`

	schema0 := `CREATE TABLE author(
		Id          SERIAL PRIMARY KEY NOT NULL,
		Login       VARCHAR(45) NOT NULL UNIQUE,
		AvatarUrl   VARCHAR(75),
		Description TEXT NOT NULL,
		Name        VARCHAR(45),
		Surname     VARCHAR(45),
		Email       VARCHAR(45),
		Password    VARCHAR(45),
		Score       VARCHAR(45)
		);`

	schema1 := `CREATE TABLE categories (
		cat  VARCHAR(45) UNIQUE
		);`

	schema2 := `CREATE TABLE tags (
			Id   SERIAL PRIMARY KEY NOT NULL,
			tag  VARCHAR(45) UNIQUE
		);`

	schema3 := `CREATE TABLE articles (
		Id           SERIAL PRIMARY KEY,
		PreviewUrl   VARCHAR(45),
		Title        TEXT,
		Text         TEXT,
		DateTime     VARCHAR(45),
		Category     VARCHAR(45) REFERENCES categories (cat) ON DELETE CASCADE,
		AuthorName   VARCHAR(45) REFERENCES author(Login) ON DELETE CASCADE,
		CommentsUrl  VARCHAR(45),
		Comments     INT,
		Likes        INT 
		);`

	schema4 := `CREATE TABLE tags_articles (
		articles_id   INT REFERENCES articles(id) ON DELETE CASCADE,
		tags_id INT REFERENCES tags(id),
		CONSTRAINT id PRIMARY KEY (articles_id, tags_id) 
		   );`

	schema5 := `CREATE OR REPLACE FUNCTION en_tsvector(title TEXT, content TEXT)
		RETURNS tsvector AS $$
		BEGIN
		RETURN (setweight(to_tsvector('english', title),'A') ||
		setweight(to_tsvector('english', content), 'B'));
		END
		$$ LANGUAGE plpgsql;`

	schema6 := `CREATE OR REPLACE FUNCTION rus_tsvector(title TEXT, content TEXT)
		RETURNS tsvector AS $$
		BEGIN
		RETURN (setweight(to_tsvector('russian', title), 'A')||
		setweight(to_tsvector('russian', content), 'B'));
		END
		$$ LANGUAGE plpgsql;`

	// schema6 := `CREATE INDEX IF NOT EXISTS idx_fts_articles ON articles
	// 	USING gin(make_tsvector(title, Text))`

	// execute a query on the server
	_, err = db.Exec(schema)
	if err != nil {
		fmt.Println(err.Error())
	}
	_, err = db.Exec(schema0)
	if err != nil {
		fmt.Println(err.Error())
	}
	_, err = db.Exec(schema1)
	if err != nil {
		fmt.Println(err.Error())
	}
	_, err = db.Exec(schema2)
	if err != nil {
		fmt.Println(err.Error())
	}
	_, err = db.Exec(schema3)
	if err != nil {
		fmt.Println(err.Error())
	}
	_, err = db.Exec(schema4)
	if err != nil {
		fmt.Println(err.Error())
	}
	_, err = db.Exec(schema5)
	if err != nil {
		fmt.Println(err.Error())
	}
	_, err = db.Exec(schema6)
	if err != nil {
		fmt.Println(err.Error())
	}
	insert_cat := `INSERT INTO categories (cat) VALUES ($1);`
	for _, data := range dataDB.CategoriesList {
		_, err = db.Exec(insert_cat, data)
		if err != nil {
			fmt.Println(err.Error())
		}
	}
	insert_author := `INSERT INTO author (Login, Name, Surname, AvatarUrl, Email, Password, Score, DESCRIPTION) VALUES ($1, $2, $3, $4, $5, $6, $7, $8);`

	for _, data := range dataDB.TestUsers {
		_, err = db.Exec(insert_author, data.Login, data.Name, data.Surname, data.AvatarUrl, data.Email, data.Password, data.Score, "Something Strange")
		if err != nil {
			fmt.Println(err.Error())
		}
	}
	names := data.TestUsers
	insert_article := `INSERT INTO articles (PreviewUrl, DateTime, Category, Title, Text, AuthorName,  CommentsUrl, Comments, Likes) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9);`
	for i, data := range data.TestData {
		date := time.Now().Format("2006/1/2 15:04")
		di := i % 26
		fmt.Println(di, dataDB.CategoriesList[di])
		_, err = db.Exec(insert_article, data.PreviewUrl, date, dataDB.CategoriesList[di], data.Title, data.Text, names[i%4].Login, data.CommentsUrl, data.Comments, data.Likes)
		if err != nil {
			fmt.Println(err.Error())
		}
	}

	tags := []string{"personal", "marketing", "finance", "design", "career", "technical"}

	insert_tag := `INSERT INTO tags (tag) VALUES ($1);`
	for _, data := range tags {
		_, err = db.Exec(insert_tag, data)
		if err != nil {
			fmt.Println(err.Error())
		}
	}

	insert_junc := `INSERT INTO tags_articles (articles_id, tags_id) VALUES 
	((SELECT Id FROM articles WHERE articles.Id = $1) ,    
	(SELECT Id FROM tags WHERE tags.Id = $2));`

	rand.Seed(4)
	for i := 1; i <= 49; i++ {
		_, err = db.Exec(insert_junc, i, rand.Int63n(4)+2)
		if err != nil {
			fmt.Println(err.Error())
		}
		_, err = db.Exec(insert_junc, i, rand.Int63n(5)+1)
		if err != nil {
			fmt.Println(err.Error())
		}
	}
	Testing()
}
