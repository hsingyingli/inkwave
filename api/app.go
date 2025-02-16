package api

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/hsingyingli/inkwave-backend/api/routes"
	db "github.com/hsingyingli/inkwave-backend/db/querier"
	"github.com/jackc/pgx/v5"
)

type App struct {
	app *fiber.App
	db  *db.Queries
}

func NewApp() (*App, error) {
	ctx := context.Background()
	app := fiber.New()
	conn, err := pgx.Connect(ctx, "user=pqgotest dbname=pqgotest sslmode=verify-full")
	if err != nil {
		return nil, err
	}
	defer conn.Close(ctx)

	return &App{
		app: app,
		db:  db.New(conn),
	}, nil
}

func (app *App) Initialize() error {
	// create repository

	// create service

	// route
	routes.RegisterRoutes(app.app)
	return nil
}

func (app *App) Listen(port string) error {
	return app.app.Listen(port)
}
