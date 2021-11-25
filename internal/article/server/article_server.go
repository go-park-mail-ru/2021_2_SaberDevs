package main

import (
	"fmt"
	"net"
	"net/http"

	app "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/article/article_app"
	arepo "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/article/repository"
	ausecase "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/article/usecase"
	srepo "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/session/repository"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":8079")

	if err != nil {
		fmt.Println("cant listen port", err)
	}
	server := grpc.NewServer(
		grpc.StreamInterceptor(grpc_prometheus.StreamServerInterceptor),
		grpc.UnaryInterceptor(grpc_prometheus.UnaryServerInterceptor),
	)
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
	prometheus.MustRegister(arepo.Hits)
	grpc_prometheus.Register(server)
	// Register Prometheus metrics handler.
	http.Handle("/metrics", promhttp.Handler())
	fmt.Println("starting server at :8079")
	server.Serve(lis)

}
