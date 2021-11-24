package main

import (
	"fmt"
	"net"

	app "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/article/article_app"
	arepo "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/article/repository"
	ausecase "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/article/usecase"
	srepo "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/session/repository"

	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":8079")
	if err != nil {
		fmt.Println("cant listen port", err)
	}

	server := grpc.NewServer()
	db, err := DbConnect()
	if err != nil {
		fmt.Println(err)
	}

	tarantoolConn, err := TarantoolConnect()
	if err != nil {
		fmt.Println(err)
	}

	defer DbClose(db)

	articleRepo := arepo.NewArticleRepository(db)
	sessionRepo := srepo.NewSessionRepository(tarantoolConn)
	articlesUsecase := ausecase.NewArticleUsecase(articleRepo, sessionRepo)
	app.RegisterArticleDeliveryServer(server, NewArticleManager(articlesUsecase))

	fmt.Println("starting server at :8079")
	server.Serve(lis)
}
