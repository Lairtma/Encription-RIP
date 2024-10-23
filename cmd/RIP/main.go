package main

import (
	_ "RIP/docs"
	"RIP/internal/api"
	"log"
)

func main() {
	app, err := api.New()
	if err != nil {
		log.Println("app can not be created", err)
		return
	}
	app.Run()
}
