package main

import (
	"context"
	"log"

	"github.com/hsingyingli/inkwave-backend/api"
	"github.com/hsingyingli/inkwave-backend/pkg/util"
)

func main() {
	ctx := context.Background()
	cfg, error := util.LoadEnv()

	if error != nil {
		log.Fatal(error)
	}

	app, error := api.NewApp(ctx, cfg)

	if error != nil {
		log.Fatal(error)
	}

	app.Initialize()

	error = app.Listen(":3000")
	if error != nil {
		log.Fatal(error)
	}
}
