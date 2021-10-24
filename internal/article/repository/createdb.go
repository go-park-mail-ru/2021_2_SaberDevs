package main

import (
	"fmt"
	"log"

	amodels "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/article/models"
	data "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/data"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

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

	schema := `CREATE TABLE IF NOT EXISTS articles (
		Id varchar(45),
		PreviewUrl   varchar(45),
		Tags         varchar(45),
		Title        varchar(45),
		Text         text,
		AuthorUrl    varchar(45),
		AuthorName   varchar(45),
		AuthorAvatar varchar(45),
		CommentsUrl  varchar(45),
		Comments     Integer,
		Likes        Integer )`

	// execute a query on the server
	_, err = db.Exec(schema)
	// schema = `DROP TABLE articles`
	// _, err = db.Exec(schema)
	insert_article := `INSERT INTO articles (id, PreviewUrl, Tags, Title, Text, AuthorUrl, AuthorName, AuthorAvatar, CommentsUrl, Comments, Likes) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`
	for _, data := range data.TestData {
		//data = data
		db.Exec(insert_article, data.Id, data.PreviewUrl, data.Tags[0], data.Title, data.Text, data.AuthorUrl, data.AuthorName, data.AuthorAvatar, data.CommentsUrl, data.Comments, data.Likes)
	}
	//_, err = db.Exec(insert_article, "123", "data.PreviewUrl", "data.Tags[0]", "data.Title", "data.Text", "data.AuthorUrl", "data.AuthorName", "data.AuthorAvatar", "data.CommentsUrl", 1, 1)
	fmt.Println(err)
	rows, err := db.Queryx("SELECT * FROM ARTICLES")
	var newArticle amodels.Article
	for rows.Next() {
		err = rows.StructScan(&newArticle)
		fmt.Println(newArticle.Id, newArticle.PreviewUrl)
	}

}
