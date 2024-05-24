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

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":  userData.Id,
		"role": userData.Role,
		"exp":  time.Now().Add(time.Hour * 8).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))

	if err != nil {
		return "", entities.Auth{}, err
	}

	return tokenString, userData, err
}
