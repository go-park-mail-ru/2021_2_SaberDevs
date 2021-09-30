package main

import (
	"saverdevs.com/2021_2_Saberdevs/server"
)

const serverAdress = "192.168.0.104:8081"

func main() {
	server.Run(serverAdress)
}
