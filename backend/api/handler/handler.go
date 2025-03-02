package handler

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/hsingyingli/inkwave/pkg/service"
	"github.com/hsingyingli/inkwave/pkg/util"
)

type HandlerManager struct {
	serviceManager *service.ServiceManager
	validate       *validator.Validate
	cfg            *util.Config
}

type Handler interface {
	CreateUser(ctx *fiber.Ctx) error
	LoginUser(ctx *fiber.Ctx) error
}

func NewHandlers(cfg *util.Config, serviceManager *service.ServiceManager) *HandlerManager {
	validate := validator.New()
	return &HandlerManager{
		serviceManager: serviceManager,
		validate:       validate,
		cfg:            cfg,
	}
}

var _ Handler = (*HandlerManager)(nil)
