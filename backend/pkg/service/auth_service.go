package service

import (
	"context"
	"strconv"

	"github.com/hsingyingli/inkwave/pkg/db"
	"github.com/hsingyingli/inkwave/pkg/util"
)

type UserAuthResult struct {
	Username     string
	Email        string
	UpdatedAt    string
	AccessToken  string
	RefreshToken string
}

func (s *ServiceManager) LoginUser(ctx context.Context, email string, password string) (UserAuthResult, error) {
	var result UserAuthResult

	user, err := s.dbRepository.GetUserByEmail(ctx, email)
	if err != nil {
		return result, err
	}
	result.Username = user.Username
	result.Email = user.Email
	result.UpdatedAt = user.UpdatedAt.Time.String()

	err = util.CheckPassword(user.Password, password)
	if err != nil {
		return result, err
	}

	result.AccessToken, err = s.CreateAuthToken(user.ID, s.cfg.ACCESS_TOKEN_DURATION, s.cfg.ACCESS_TOKEN_SECRET_KEY)
	if err != nil {
		return result, err
	}

	result.RefreshToken, err = s.CreateAuthToken(user.ID, s.cfg.REFRESH_TOKEN_DURATION, s.cfg.REFRESH_TOKEN_SECRET_KEY)
	if err != nil {
		return result, err
	}

	return result, nil
}

func (s *ServiceManager) CreateUser(ctx context.Context, username string, email string, password string) error {
	hashedPassword, err := util.HashPassword(password)
	if err != nil {
		return err
	}

	_, err = s.dbRepository.CreateUser(ctx, db.CreateUserParams{
		Username: username,
		Email:    email,
		Password: hashedPassword,
	})
	if err != nil {
		return err
	}

	return nil
}

func (s *ServiceManager) RenewToken(ctx context.Context, refreshToken string) (UserAuthResult, error) {
	var result UserAuthResult

	claim, err := s.ValidateAuthToken(refreshToken, s.cfg.REFRESH_TOKEN_SECRET_KEY)
	if err != nil {
		return result, err
	}

	userId, err := strconv.ParseInt(claim.Subject, 10, 64)
	if err != nil {
		return result, err
	}

	user, err := s.dbRepository.GetUser(ctx, userId)
	if err != nil {
		return result, err
	}

	accessToken, err := s.CreateAuthToken(user.ID, s.cfg.ACCESS_TOKEN_DURATION, s.cfg.ACCESS_TOKEN_SECRET_KEY)
	if err != nil {
		return result, err
	}

	result = UserAuthResult{
		Username:     user.Username,
		Email:        user.Email,
		UpdatedAt:    user.UpdatedAt.Time.String(),
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	return result, nil
}
