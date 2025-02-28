package route

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/hsingyingli/inkwave-backend/api/handler"
	"github.com/hsingyingli/inkwave-backend/api/middleware"
)

func RegisterRoutes(app *fiber.App, middlewares *middleware.Middlewares, handlers *handler.Handlers) {
	userApi := app.Group("/users")
	{
		userApi.Get("/", func(c *fiber.Ctx) error {
			return c.SendString("User")
		})

		type CreateUserRequest struct {
			Username string `json:"username" form:"username" binding:"required"`
			Email    string `json:"email" form:"email" binding:"required"`
			Password string `json:"password" form:"password" binding:"required"`
		}

		userApi.Post("/", func(c *fiber.Ctx) error {
			params := new(CreateUserRequest)
			if err := c.BodyParser(params); err != nil {
				return c.SendStatus(http.StatusBadRequest)
			}

			fmt.Println(params)

			return c.SendStatus(http.StatusCreated)
		})
	}
}
