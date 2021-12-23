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
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

var (
	zapLogger  *zap.Logger
	customFunc grpc_zap.CodeToLevel
)

func main() {
	lis, err := net.Listen("tcp", ":8079")
	if err != nil {
		fmt.Println("cant listen port", err)
	}
	server := grpc.NewServer()
	db, err := ser.DbConnect()
	if err != nil {
		fmt.Println(err)
	}

	tarantoolConn, err := ser.TarantoolConnect()
	if err != nil {
		fmt.Println(err)
	}

	defer ser.DbClose(db)

	log := wrapper.NewLogger()
	articleRepo := arepo.NewArticleRepository(db, log)
	sessionRepo := srepo.NewSessionRepository(tarantoolConn, log)
	articlesUsecase := ausecase.NewArticleUsecase(articleRepo, sessionRepo)
	app.RegisterArticleDeliveryServer(server, ser.NewArticleManager(articlesUsecase))
	//prometheus.MustRegister(wrapper.Hits, wrapper.Duration, wrapper.Errors)
	http.Handle("/metrics", promhttp.Handler())
	prometheus.MustRegister(wrapper.Hits, wrapper.Duration, wrapper.Errors)
	// Register Prometheus metrics handler.
	go func() {
		err := http.ListenAndServe(":8074", nil)
		log.Logger.Fatal(err.Error())
	}()
	fmt.Println("starting server at :8079")
	server.Serve(lis)

}
