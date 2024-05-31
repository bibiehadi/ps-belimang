package authservice

import (
	"belimang/src/entities"
	"belimang/src/helpers"
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

func (service *authService) Login(loginRequest entities.LoginRequest) (string, entities.Auth, error) {
	userData, err := service.authRepository.FindByUsername(loginRequest.Username)

	if err != nil {
		return "", entities.Auth{}, errors.New("INVALID USERNAME OR PASSWORD")
	}

	if !helpers.CompareHashAndPassword(userData.Password, loginRequest.Password) {
		return "", entities.Auth{}, errors.New("INVALID USERNAME OR PASSWORD")
	}

	claims := &entities.CustomClaims{
		UserId:   userData.Id,
		Username: userData.Username,
		Role:     string(userData.Role),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 8).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))

	if err != nil {
		return "", entities.Auth{}, err
	}

	return tokenString, userData, err
}
