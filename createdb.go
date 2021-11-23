package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	server "github.com/go-park-mail-ru/2021_2_SaberDevs/cmd/sybernews"
	"github.com/go-park-mail-ru/2021_2_SaberDevs/internal/data"
	dataDB "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/data"
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
		Title        VARCHAR(350),
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
		_, err = db.Exec(insert_article, data.PreviewUrl, date, dataDB.CategoriesList[i], data.Title, data.Text, names[i/4].Login, data.CommentsUrl, data.Comments, data.Likes)
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
	for i := 1; i <= 11; i++ {
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
