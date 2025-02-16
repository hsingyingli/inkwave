package routes

import (
	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(app *fiber.App) {
	userApi := app.Group("/users")
	{
		userApi.Get("/", func(c *fiber.Ctx) error {
			return c.SendString("User")
		})
	}
}
