package handler

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

type CreateUserRequest struct {
	Username string `json:"username" validate:"required,min=3,max=32"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,max=64"`
}

func (h *HandlerManager) CreateUser(c *fiber.Ctx) error {
	var req CreateUserRequest
	if err := c.BodyParser(&req); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	if err := h.validate.Struct(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"errors": err.Error(),
		})
	}

	ctx := c.Context()
	if err := h.serviceManager.CreateUser(ctx, req.Username, req.Email, req.Password); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"errors": err.Error(),
		})
	}

	return c.SendStatus(fiber.StatusCreated)
}

type LoginUserRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,max=64"`
}

type LoginUserResponse struct {
	AccessToken string `json:"access_token"`
	Username    string `json:"username"`
	Email       string `json:"email"`
	UpdatedAt   string `json:"updated_at"`
}

func (h *HandlerManager) LoginUser(c *fiber.Ctx) error {
	var req LoginUserRequest
	if err := c.BodyParser(&req); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	if err := h.validate.Struct(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"errors": err.Error(),
		})
	}

	result, err := h.serviceManager.LoginUser(c.Context(), req.Email, req.Password)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{})
	}

	c.Cookie(&fiber.Cookie{
		Name:     "inkwave_refresh_token",
		Value:    result.RefreshToken,
		Expires:  time.Now().Add(h.cfg.REFRESH_TOKEN_DURATION),
		HTTPOnly: true,
	})

	return c.Status(fiber.StatusOK).JSON(LoginUserResponse{
		AccessToken: result.AccessToken,
		Username:    result.Username,
		Email:       result.Email,
		UpdatedAt:   result.UpdatedAt,
	})
}

type RenewAccessTokenResponse struct {
	AccessToken string `json:"access_token"`
	Username    string `json:"username"`
	Email       string `json:"email"`
	UpdatedAt   string `json:"updated_at"`
}

func (h *HandlerManager) RenewAccessToken(c *fiber.Ctx) error {
	refreshToken := c.Cookies("inkwave_refresh_token")

	if refreshToken == "" {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	result, err := h.serviceManager.RenewToken(c.Context(), refreshToken)
	if err != nil {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	return c.Status(fiber.StatusOK).JSON(RenewAccessTokenResponse{
		AccessToken: result.AccessToken,
		Username:    result.Username,
		Email:       result.Email,
		UpdatedAt:   result.UpdatedAt,
	})
}
