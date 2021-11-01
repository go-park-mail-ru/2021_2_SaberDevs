package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"

	server "github.com/go-park-mail-ru/2021_2_SaberDevs/cmd/sybernews"
	amodels "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/article/models"
	repo "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/article/repository"
	data "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/data"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

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
		DROP TABLE IF EXISTS categories CASCADE;
		DROP TABLE IF EXISTS categories_articles CASCADE;`

	schema1 := `CREATE TABLE author(
		Id       SERIAL PRIMARY KEY NOT NULL,
		Login    VARCHAR(45) NOT NULL UNIQUE,
		Name     VARCHAR(45),
		Surname  VARCHAR(45),
		Email    VARCHAR(45),
		Password VARCHAR(45),
		Score    VARCHAR(45)
		);`

	schema2 := `CREATE TABLE categories (
		Id   SERIAL PRIMARY KEY NOT NULL,
		tag  VARCHAR(45) UNIQUE
		);`

	schema3 := `CREATE TABLE articles (
		Id           SERIAL PRIMARY KEY,
		PreviewUrl   VARCHAR(45),
		Title        VARCHAR(45),
		Text         TEXT,
		AuthorUrl    VARCHAR(45),
		AuthorName   VARCHAR(45) REFERENCES author(Login) ON DELETE CASCADE,
		AuthorAvatar VARCHAR(45),
		CommentsUrl  VARCHAR(45),
		Comments     INT,
		Likes        INT 
		);`

	schema4 := `CREATE TABLE categories_articles (
		articles_id   INT REFERENCES articles(id) ON DELETE CASCADE,
		categories_id INT REFERENCES categories(id),
		CONSTRAINT id PRIMARY KEY (articles_id, categories_id) 
		   );`

	// execute a query on the server
	_, err = db.Exec(schema)
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

	insert_author := `INSERT INTO author (Login, Name, Surname, Email, Password, Score) VALUES ($1, $2, $3, $4, $5, $6);`

	for _, data := range data.TestUsers {
		_, err = db.Exec(insert_author, data.Login, data.Name, data.Surname, data.Email, data.Password, data.Score)
		if err != nil {
			fmt.Println(err.Error())
		}
	}
	rows, err := db.Queryx("SELECT * FROM author")
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

	insert_article := `INSERT INTO articles (PreviewUrl, Title, Text, AuthorUrl, AuthorName, AuthorAvatar, CommentsUrl, Comments, Likes) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9);`
	for i, data := range data.TestData {
		_, err = db.Exec(insert_article, data.PreviewUrl, data.Title, data.Text, data.AuthorUrl, names[i/4], data.AuthorAvatar, data.CommentsUrl, data.Comments, data.Likes)
		if err != nil {
			fmt.Println(err.Error())
		}
	}

	rows, err = db.Queryx("SELECT * FROM ARTICLES")
	if err != nil {
		log.Fatal(err)
	}
	var newArticle amodels.DbArticle
	for rows.Next() {
		err = rows.StructScan(&newArticle)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Print(newArticle.Id, "  ", newArticle.PreviewUrl, "  ", newArticle.AuthorName, "  ", newArticle.Likes, "\n")
	}

	categories := []string{"personal", "marketing", "finance", "design", "career", "technical"}

	insert_cat := `INSERT INTO categories (tag) VALUES ($1);`
	for _, data := range categories {
		_, err = db.Exec(insert_cat, data)
		if err != nil {
			fmt.Println(err.Error())
		}
	}
	fmt.Print("whereami", "\n")
	rows, err = db.Queryx("SELECT * FROM categories;")
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

	insert_junc := `INSERT INTO categories_articles (articles_id, categories_id) VALUES 
	((SELECT Id FROM articles WHERE Id = $1) ,    
	(SELECT Id FROM categories WHERE Id = $2));`

	rand.Seed(4)
	for i := 1; i <= 12; i++ {
		_, err = db.Exec(insert_junc, i, rand.Int63n(4)+2)
		if err != nil {
			fmt.Println(err.Error())
		}
		_, err = db.Exec(insert_junc, i, rand.Int63n(5)+1)
		if err != nil {
			fmt.Println(err.Error())
		}
	}

	rows, err = db.Queryx("SELECT * FROM categories_articles;")
	if err != nil {
		fmt.Println(err.Error())
	}
	var tag amodels.СategoriesArticles
	for rows.Next() {
		err = rows.Scan(&tag.Articles_id, &tag.Categories_id)
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Print(tag.Articles_id, "  ", tag.Categories_id, "\n")
	}

	rows, err = db.Queryx(`select c.tag from categories c
	inner join categories_articles ca  on c.Id = ca.categories_id
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
	fmt.Print(result.Id, " ", result.AuthorName, " ", result.Tags, " ", result.Text, " ", result.Likes, "\n")

	results, err := myRepo.GetByAuthor(context.TODO(), "dar")
	if err != nil {
		fmt.Println(err.Error())
	}
	for _, result := range results {
		fmt.Print(result.Id, " ", result.AuthorName, " ", result.Tags, " ", result.Text, " ", result.Likes, "\n")
	}
	fmt.Println()
	results, err = myRepo.GetByTag(context.TODO(), "finance")
	if err != nil {
		fmt.Println(err.Error())
	}
	for _, result := range results {
		fmt.Print(result.Id, " ", result.AuthorName, " ", result.Tags, " ", result.Text, " ", result.Likes, "\n")
	}

	fmt.Println()
	ar := data.TestData[3]
	ar.Id = "13"
	ar.AuthorName = "dar"
	ar.Tags = append(ar.Tags, "finance")
	Id, err := myRepo.Store(context.TODO(), &ar)
	ar.Id = fmt.Sprint(Id)
	fmt.Println("IDDDDD=", Id)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("!", count)
	result, err = myRepo.GetByID(context.TODO(), 13)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Print(result.Id, " ", result.AuthorName, " ", result.Tags, " ", result.Text, " ", result.Likes, "\n")

	err = myRepo.Delete(context.TODO(), 3)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println()
	newresult, err := myRepo.Fetch(context.TODO(), 0, 6)
	if err != nil {
		fmt.Println(err.Error())
	}
	for _, result := range newresult {
		fmt.Print(result.Id, " ", result.AuthorName, " ", result.Tags, " ", result.Text, " ", result.Likes, "\n")
	}
	newresult, err = myRepo.Fetch(context.TODO(), 5, 6)
	if err != nil {
		fmt.Println(err.Error())
	}
	for _, result := range newresult {
		fmt.Print(result.Id, " ", result.AuthorName, " ", result.Tags, " ", result.Text, " ", result.Likes, "\n")
	}
	fmt.Println()
	fmt.Println()
	ar.Text = `<a onblur="alert(secret)" href="http://www.google.com">Google</a>`
	ar.Tags = append(ar.Tags, "jojo")
	err = myRepo.Update(context.TODO(), &ar)
	if err != nil {
		fmt.Println(err.Error())
	}
	result, err = myRepo.GetByID(context.TODO(), int64(Id))
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Print(result.Id, " ", result.AuthorName, " ", result.Tags, " ", result.Text, " ", result.Likes, "\n")
}
