package main

import (
	"context"
	"log"

	"github.com/hsingyingli/inkwave/api"
	"github.com/hsingyingli/inkwave/pkg/util"
)

func main() {
	ctx := context.Background()
	cfg, error := util.LoadEnv()

	if error != nil {
		log.Fatal(error)
	}

	app, error := api.NewApp(ctx, cfg)

	defer app.Shutdown()

	if error != nil {
		log.Fatal(error)
	}

	error = app.Listen(":3000")
	if error != nil {
		log.Fatal(error)
	}
}
