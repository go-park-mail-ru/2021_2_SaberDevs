package main

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"

	wrapper "github.com/go-park-mail-ru/2021_2_SaberDevs/internal"
	app "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/article/article_app"
	arepo "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/article/repository"
	ser "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/article/server/serve"
	ausecase "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/article/usecase"
	srepo "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/session/repository"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
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
	server := grpc.NewServer(grpc_middleware.ChainUnaryServer(grpc_zap.UnaryServerInterceptor(zapLogger)))
	db, err := ser.DbConnect()
	if err != nil {
		fmt.Println(err)
	}

	tarantoolConn, err := ser.TarantoolConnect()
	if err != nil {
		fmt.Println(err)
	}

	defer ser.DbClose(db)

	rawJSON := []byte(`{
		"level": "debug",
		"encoding": "json",
		"outputPaths": ["stdout", "/tmp/logs"],
		"errorOutputPaths": ["stderr"],
		"initialFields": {"foo": "bar"},
		"encoderConfig": {
		  "messageKey": "message",
		  "levelKey": "level",
		  "levelEncoder": "lowercase"
		}
	  }`)

	var cfg zap.Config
	if err := json.Unmarshal(rawJSON, &cfg); err != nil {
		panic(err)
	}
	logger, err := cfg.Build()
	if err != nil {
		panic(err)
	}
	defer logger.Sync()
	log := wrapper.NewMyLogger(logger)
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
		logger.Fatal(err.Error())
	}()
	fmt.Println("starting server at :8079")
	server.Serve(lis)

}
