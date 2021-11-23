package main

import (
	"context"
	"fmt"
	"log"
	"time"

	server "github.com/go-park-mail-ru/2021_2_SaberDevs/cmd/sybernews"
	amodels "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/article/models"
	repo "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/article/repository"
	"github.com/go-park-mail-ru/2021_2_SaberDevs/internal/data"
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

	fmt.Println()
	newresult, err = myRepo.FindArticles(context.TODO(), "приз", 0, 4)
	if err != nil {
		fmt.Println(err.Error())
	}
	for _, result := range newresult {
		fmt.Print(result.Id, " ", result.Title, " ", result.Category, " ", result.Author.Name, " ", result.Tags, " ", result.Text, " ", result.Likes, "\n")
	}

	fmt.Println()
	newresult, err = myRepo.FindArticles(context.TODO(), "Prog", 0, 4)
	if err != nil {
		fmt.Println(err.Error())
	}
	for _, result := range newresult {
		fmt.Print(result.Id, " ", result.Title, " ", result.Category, " ", result.Author.Name, " ", result.Tags, " ", result.Text, " ", result.Likes, "\n")
	}

	fmt.Println()
	newresult, err = myRepo.FindByTag(context.TODO(), "in", 0, 10)
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
}
