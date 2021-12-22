package main

import (
	"fmt"
	"net"
	"net/http"

	wrapper "github.com/go-park-mail-ru/2021_2_SaberDevs/internal"
	app "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/article/article_app"
	arepo "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/article/repository"
	ser "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/article/server/serve"
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
	db, err := ser.DbConnect()
	if err != nil {
		fmt.Println(err)
	}

	tarantoolConn, err := ser.TarantoolConnect()
	if err != nil {
		fmt.Println(err)
	}

	defer ser.DbClose(db)

	articleRepo := arepo.NewArticleRepository(db)
	sessionRepo := srepo.NewSessionRepository(tarantoolConn)
	articlesUsecase := ausecase.NewArticleUsecase(articleRepo, sessionRepo)
	app.RegisterArticleDeliveryServer(server, ser.NewArticleManager(articlesUsecase))
	grpc_prometheus.Register(server)
	prometheus.MustRegister(wrapper.Hits)
	// Register Prometheus metrics handler.
	http.Handle("/metrics", promhttp.Handler())
	fmt.Println("starting server at :8079")
	// go log.Fatal(http.ListenAndServe(":8078", nil))
	server.Serve(lis)

}
