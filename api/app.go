package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hsingyingli/inkwave-backend/api/routes"
)

type App struct {
	app *fiber.App
}

func NewApp() (*App, error) {
	app := fiber.New()

	return &App{
		app: app,
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
