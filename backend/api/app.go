package api

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/hsingyingli/inkwave/api/handler"
	"github.com/hsingyingli/inkwave/api/middleware"
	"github.com/hsingyingli/inkwave/api/route"
	"github.com/hsingyingli/inkwave/pkg/db"
	"github.com/hsingyingli/inkwave/pkg/service"
	"github.com/hsingyingli/inkwave/pkg/util"
	"github.com/jackc/pgx/v5/pgxpool"
)

type App struct {
	ctx          context.Context
	app          *fiber.App
	cfg          util.Config
	dbConnection *pgxpool.Pool
}

func NewApp(ctx context.Context, cfg util.Config) (*App, error) {
	app := fiber.New()
	conn, err := pgxpool.New(ctx, cfg.DB_URL)
	if err != nil {
		return nil, err
	}

	dbRepository := db.New(conn)
	serviceManager := service.NewServices(dbRepository)
	handlerManager := handler.NewHandlers(serviceManager)
	middlewareManager := middleware.NewMiddlewares(serviceManager)

	route.RegisterRoutes(app, middlewareManager, handlerManager)

	return &App{
		ctx:          ctx,
		app:          app,
		cfg:          cfg,
		dbConnection: conn,
	}, nil
}

func (app *App) Shutdown() error {
	if app.dbConnection != nil {
		app.dbConnection.Close()
	}
	return nil
}

func (app *App) Listen(port string) error {
	return app.app.Listen(port)
}
