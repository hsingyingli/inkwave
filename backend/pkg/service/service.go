package service

import (
	"context"

	"github.com/hsingyingli/inkwave/pkg/db"
)

type ServiceManager struct {
	dbRepository *db.Queries
}

type Service interface {
	CreateUser(ctx context.Context, username string, email string, password string) error
}

func NewServices(dbRepository *db.Queries) *ServiceManager {
	return &ServiceManager{
		dbRepository: dbRepository,
	}
}

var _ Service = (*ServiceManager)(nil)
