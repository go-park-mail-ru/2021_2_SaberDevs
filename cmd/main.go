package main

import (
	"github.com/go-park-mail-ru/2021_2_SaberDevs/cmd/sybernews"
	"log"
)

func main() {
	serverAddress, err := server.Config()
	if err != nil {
		log.Fatal("Error starting server: ", err)
	}
	server.Run(serverAddress)
}
