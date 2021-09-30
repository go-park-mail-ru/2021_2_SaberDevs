package main

import (
	"server/server"
)

func main() {
	const serverAdress = "192.168.0.104:8081"
	server.Run(serverAdress)
}
