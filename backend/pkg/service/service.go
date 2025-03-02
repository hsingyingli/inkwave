package service

import (
	"context"
	"time"

	"github.com/hsingyingli/inkwave/pkg/db"
	"github.com/hsingyingli/inkwave/pkg/util"
)

type ServiceManager struct {
	cfg          *util.Config
	dbRepository *db.Queries
}

type Service interface {
	CreateUser(ctx context.Context, username string, email string, password string) error
	CreateAuthToken(userId int64, duration time.Duration, secretKey string) (string, error)
	ValidateAuthToken(token string, secretKey string) (*AuthClaim, error)
}

func NewServices(cfg *util.Config, dbRepository *db.Queries) *ServiceManager {
	return &ServiceManager{
		cfg:          cfg,
		dbRepository: dbRepository,
	}
}

var _ Service = (*ServiceManager)(nil)
