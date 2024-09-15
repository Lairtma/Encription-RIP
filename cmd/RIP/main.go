package main

import (
	"log"

	"RIP/internal/api"
)

func main() {
	log.Println("Server start up")
	api.StartServer()
	log.Println("Server down")
}
