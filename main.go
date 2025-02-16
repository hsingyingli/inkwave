package main

import (
	"log"

	"github.com/hsingyingli/inkwave-backend/api"
)

func main() {
	app, error := api.NewApp()

	if error != nil {
		log.Fatal(error)
	}

	app.Initialize()

	error = app.Listen(":3000")
	if error != nil {
		log.Fatal(error)
	}
}
