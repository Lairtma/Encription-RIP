package main

import (
	"RIP/internal/api"
	"log"
)

func main() {
	app, err := app.New()
	if err != nil {
		log.Println("app can not be created", err)
		return
	}
	app.Run()
}
