package handler

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/hsingyingli/inkwave/pkg/service"
)

type HandlerManager struct {
	serviceManager *service.ServiceManager
	validate       *validator.Validate
}

type Handler interface {
	CreateUser(ctx *fiber.Ctx) error
}

func NewHandlers(serviceManager *service.ServiceManager) *HandlerManager {
	validate := validator.New()
	return &HandlerManager{
		serviceManager: serviceManager,
		validate:       validate,
	}
}

var _ Handler = (*HandlerManager)(nil)
