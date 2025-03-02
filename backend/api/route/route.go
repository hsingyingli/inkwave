package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hsingyingli/inkwave/api/handler"
	"github.com/hsingyingli/inkwave/api/middleware"
)

func RegisterRoutes(app *fiber.App, middlewares *middleware.Middlewares, handlerManager *handler.HandlerManager) {
	authApi := app.Group("/auth")
	{
		authApi.Post("/register", handlerManager.CreateUser)
		authApi.Post("/login", handlerManager.LoginUser)
		authApi.Post("/renew", handlerManager.RenewAccessToken)
	}
}
