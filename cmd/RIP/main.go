package main

import (
	"log"

	"RIP/internal/api"
)

func main() {
	app, err := api.New()
	if err != nil {
		log.Println("app can not be created", err)
		return
	}
	app.Run()
}
