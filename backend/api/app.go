package api

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/hsingyingli/inkwave-backend/api/handler"
	"github.com/hsingyingli/inkwave-backend/api/middleware"
	"github.com/hsingyingli/inkwave-backend/api/route"
	"github.com/hsingyingli/inkwave-backend/pkg/db"
	"github.com/hsingyingli/inkwave-backend/pkg/util"
	"github.com/jackc/pgx/v5"
)

type App struct {
	app *fiber.App
	db  *db.Queries
	cfg util.Config
}

func NewApp(ctx context.Context, cfg util.Config) (*App, error) {
	app := fiber.New()
	conn, err := pgx.Connect(ctx, cfg.DB_URL)
	if err != nil {
		return nil, err
	}
	defer conn.Close(ctx)

	return &App{
		app: app,
		db:  db.New(conn),
		cfg: cfg,
	}, nil
}

func (app *App) Initialize() {
	// create repository

	// create service

	// create handler
	//
	handlers := handler.NewHandlers()

	middlewares := middleware.NewMiddlewares()
	// route
	route.RegisterRoutes(app.app, middlewares, handlers)
}

func (app *App) Listen(port string) error {
	return app.app.Listen(port)
}
