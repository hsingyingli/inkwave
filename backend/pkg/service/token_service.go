package service

import (
	"errors"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type AuthClaim struct {
	jwt.RegisteredClaims
}

func (s *ServiceManager) CreateAuthToken(userId int64, duration time.Duration, secretKey string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, AuthClaim{
		jwt.RegisteredClaims{
			Issuer:    "inkwave",
			Subject:   strconv.FormatInt(userId, 10),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(duration)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	})

	tokenString, err := token.SignedString([]byte(secretKey))

	return tokenString, err
}

func (s *ServiceManager) ValidateAuthToken(token string, secretKey string) (*AuthClaim, error) {
	parsedToken, err := jwt.ParseWithClaims(token, &AuthClaim{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}

		return []byte(secretKey), nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := parsedToken.Claims.(*AuthClaim); ok && parsedToken.Valid {
		if claims.Issuer != "inkwave" {
			return nil, errors.New("invalid issuer")
		}
		return claims, nil
	}

	return nil, errors.New("invalid token")
}
