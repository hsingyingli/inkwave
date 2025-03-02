package service

import (
	"context"

	"github.com/hsingyingli/inkwave/pkg/db"
)

func (s *ServiceManager) CreateUser(ctx context.Context, username string, email string, password string) error {
	_, err := s.dbRepository.CreateUser(ctx, db.CreateUserParams{
		Username: username,
		Email:    email,
		Password: password,
	})
	if err != nil {
		return err
	}

	return nil
}
