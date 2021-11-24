package main

import (
	"fmt"
	"log"
	"net"

	app "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/article/article_app"
	arepo "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/article/repository"
	ausecase "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/article/usecase"
	srepo "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/session/repository"

	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatalln("cant listen port", err)
	}

	server := grpc.NewServer()
	db, err := DbConnect()
	if err != nil {
		// e.Logger.Fatal(err)
	}

	tarantoolConn, err := TarantoolConnect()
	if err != nil {
		// e.Logger.Fatal(err)
	}

	defer DbClose(db)

	articleRepo := arepo.NewArticleRepository(db)
	sessionRepo := srepo.NewSessionRepository(tarantoolConn)
	articlesUsecase := ausecase.NewArticleUsecase(articleRepo, sessionRepo)
	app.RegisterArticleDeliveryServer(server, NewArticleManager(articlesUsecase))

	fmt.Println("starting server at :8081")
	server.Serve(lis)
}
