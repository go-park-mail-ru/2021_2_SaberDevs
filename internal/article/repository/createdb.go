package main

// package article

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

// func remain() {

func main() {
	connStr := "user=postgres dbname=postgres password=yura11011 host=localhost sslmode=disable"
	db, err := sqlx.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	// type Article struct {
	// 	Id           string   `json:"id"`
	// 	PreviewUrl   string   `json:"previewUrl"`
	// 	Tags         []string `json:"tags"`
	// 	Title        string   `json:"title"`
	// 	Text         string   `json:"text"`
	// 	AuthorUrl    string   `json:"authorUrl"`
	// 	AuthorName   string   `json:"authorName"`
	// 	AuthorAvatar string   `json:"authorAvatar"`
	// 	CommentsUrl  string   `json:"commentsUrl"`
	// 	Comments     uint     `json:"comments"`
	// 	Likes        uint     `json:"likes"`
	// }

	// type User struct {
	// 	Login    string `json:"login"`
	// 	Name     string `json:"name"`
	// 	Surname  string `json:"surname"`
	// 	Email    string `json:"email" valid:"email,optional"`
	// 	Password string `json:"password"`
	// 	Score    int    `json:"score"`
	// }
	schema := `DROP TABLE IF EXISTS articles;
		DROP TABLE IF EXISTS author;
		DROP TABLE IF EXISTS categories;
		DROP TABLE IF EXISTS categories_articles;`

	schema1 := `CREATE TABLE author(
		Id SERIAL PRIMARY KEY NOT NULL,
		Login    VARCHAR(45),
		Name     VARCHAR(45) NOT NULL UNIQUE,
		Surname  VARCHAR(45),
		Email    VARCHAR(45),
		Password VARCHAR(45),
		score    VARCHAR(45)
		);`

	schema2 := `CREATE TABLE categories (
		id SERIAL PRIMARY KEY NOT NULL,
		tag  VARCHAR(45)
		);`

	schema3 := `CREATE TABLE articles (
		    id SERIAL PRIMARY KEY NOT NULL,
			name_id VARCHAR(45),
			PreviewUrl   VARCHAR(45),
			Title        VARCHAR(45),
			Text         TEXT,
			AuthorUrl    VARCHAR(45),
			AuthorName   VARCHAR(45) REFERENCES author(Name),
			AuthorAvatar VARCHAR(45),
			CommentsUrl  VARCHAR(45),
			Comments     INT,
			Likes        INT );`

	schema4 := `CREATE TABLE categories_articles (
		articles_id INT REFERENCES articles(id),
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

	// insert_article := `INSERT INTO articles (id, PreviewUrl, Tags, Title, Text, AuthorUrl, AuthorName, AuthorAvatar, CommentsUrl, Comments, Likes) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`
	// for _, data := range data.TestData {
	// 	//data = data
	// 	db.Exec(insert_article, data.Id, data.PreviewUrl, data.Tags[0], data.Title, data.Text, data.AuthorUrl, data.AuthorName, data.AuthorAvatar, data.CommentsUrl, data.Comments, data.Likes)
	// }
	// //_, err = db.Exec(insert_article, "123", "data.PreviewUrl", "data.Tags[0]", "data.Title", "data.Text", "data.AuthorUrl", "data.AuthorName", "data.AuthorAvatar", "data.CommentsUrl", 1, 1)
	// fmt.Println(err)
	// rows, err := db.Queryx("SELECT * FROM ARTICLES")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// var newArticle amodels.Article
	// for rows.Next() {
	// 	err = rows.StructScan(&newArticle)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	fmt.Println(newArticle.Id, newArticle.PreviewUrl)
	// }

}
