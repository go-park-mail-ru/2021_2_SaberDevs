package server

import (
	"fmt"
	app "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/article/article_app"
	arepo "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/article/repository"
	krepo "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/keys/repository"
	srepo "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/session/repository"
	urepo "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/user/repository"
	uusecase "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/user/usecase"
	"google.golang.org/grpc"
	"net"
)

func main() {
	lis, err := net.Listen("tcp", ":8078")
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

	userRepo := urepo.NewUserRepository(db)
	sessionRepo := srepo.NewSessionRepository(tarantoolConn)
	keyRepo := krepo.NewKeyRepository(tarantoolConn)
	articleRepo := arepo.NewArticleRepository(db)

	userUsecase := uusecase.NewUserUsecase(userRepo, sessionRepo, keyRepo, articleRepo)

	app.

	// userRepo := arepo.NewArticleRepository(db)
	// sessionRepo := srepo.NewSessionRepository(tarantoolConn)
	// articlesUsecase := ausecase.NewArticleUsecase(articleRepo, sessionRepo)
	// app.RegisterArticleDeliveryServer(server, NewArticleManager(articlesUsecase))

	fmt.Println("user starting server at :8079")
	server.Serve(lis)
}