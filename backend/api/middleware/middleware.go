package middleware

import "github.com/hsingyingli/inkwave/pkg/service"

type Middlewares struct{}

func NewMiddlewares(service *service.ServiceManager) *Middlewares {
	return &Middlewares{}
}
