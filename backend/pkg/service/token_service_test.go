package service

import (
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCreateAuthToken(t *testing.T) {
	secretKey := "mysecretkey"
	service := NewServices(nil)
	userId := int64(12345)
	duration := time.Hour

	token, err := service.CreateAuthToken(userId, duration, secretKey)
	assert.NoError(t, err)
	assert.NotEmpty(t, token)
}

func TestValidateAuthToken(t *testing.T) {
	secretKey := "mysecretkey"
	service := NewServices(nil)
	userId := int64(12345)
	duration := time.Hour

	token, err := service.CreateAuthToken(userId, duration, secretKey)
	assert.NoError(t, err)
	assert.NotEmpty(t, token)

	claims, err := service.ValidateAuthToken(token, secretKey)
	assert.NoError(t, err)
	assert.NotNil(t, claims)
	assert.Equal(t, strconv.FormatInt(userId, 10), claims.Subject)
}

func TestValidateAuthToken_InvalidToken(t *testing.T) {
	secretKey := "mysecretkey"
	service := NewServices(nil)
	invalidToken := "invalidtoken"

	claims, err := service.ValidateAuthToken(invalidToken, secretKey)
	assert.Error(t, err)
	assert.Nil(t, claims)
}
