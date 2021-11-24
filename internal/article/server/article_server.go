package main

import (
	"fmt"
	"log"
	"net"

	app "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/article/article_app"
	ahandler "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/article/handler"

	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatalln("cant listen port", err)
	}

	server := grpc.NewServer()
	handler := &ahandler.ArticlesHandler{}

	app.RegisterArticleDeliveryServer(server, NewArticleManager(handler))

	fmt.Println("starting server at :8081")
	server.Serve(lis)
}
