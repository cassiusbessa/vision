package service

import (
	"fmt"

	"github.com/cassiusbessa/vision-social-media/domain/service/errors"
	"github.com/golang-jwt/jwt/v5"
)

type TokenService struct {
	secretKey string
}

func NewTokenService() *TokenService {
	return &TokenService{
		secretKey: "your-secure",
	}
}

func (s *TokenService) verifyToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.secretKey), nil
	})
	if err != nil {
		return nil, errors.NewUnauthorized("Invalid token")
	}
	return token, nil
}

func (s *TokenService) GetPayload(tokenString string) (string, error) {

	token, err := s.verifyToken(tokenString)
	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", errors.NewUnauthorized("Invalid token")
	}
	fmt.Println(claims)
	return claims["accountId"].(string), nil
}
