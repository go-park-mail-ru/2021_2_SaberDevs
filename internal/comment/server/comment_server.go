package main

import (
	"fmt"
	server "github.com/go-park-mail-ru/2021_2_SaberDevs/cmd/sybernews"
	app "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/comment/comment_app"
	crepo "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/comment/repository"
	cusecase "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/comment/usecase"
	srepo "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/session/repository"
	urepo "github.com/go-park-mail-ru/2021_2_SaberDevs/internal/user/repository"
	"github.com/jmoiron/sqlx"
	"github.com/tarantool/go-tarantool"
	"google.golang.org/grpc"
	"net"
)

func TarantoolConnect() (*tarantool.Connection, error) {
	user, pass, addr, err := server.TarantoolConfig()
	if err != nil {
		return nil, err
	}

	opts := tarantool.Opts{User: user, Pass: pass}
	conn, err := tarantool.Connect(addr, opts)
	if err != nil {
		return nil, err
	}

	_, err = conn.Ping()
	if err != nil {
		return nil, err
	}
	return conn, nil
}

func DbConnect() (*sqlx.DB, error) {
	connStr, err := server.DbConfig()
	if err != nil {
		return nil, err
	}
	db, err := sqlx.Open("postgres", connStr)
	if err != nil {
		return db, err
	}
	err = db.Ping()
	if err != nil {
		return db, err
	}
	return db, err
}

func DbClose(db *sqlx.DB) error {
	err := db.Close()
	if err != nil {
		return err
	}
	return err
}

func main() {
	lis, err := net.Listen("tcp", ":8077")
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
	commentsRepo := crepo.NewCommentRepository(db)

	commentUsecase := cusecase.NewCommentUsecase(userRepo, sessionRepo, commentsRepo)

	app.RegisterCommentDeliveryServer(server, NewCommentManager(commentUsecase))

	fmt.Println("starting comment server at :8077")
	server.Serve(lis)
}