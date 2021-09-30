package main

import (
	"server/server"
)

const serverAdress = "192.168.0.104:8081"

func main() {
	server.Run(serverAdress)
}
