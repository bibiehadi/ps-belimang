package authservice

import (
	"belimang/src/entities"
	"belimang/src/helpers"
	"errors"
)

func (service *authService) Register(registerRequest entities.RegisterRequest, isAdmin bool) (entities.Auth, error) {

	if service.authRepository.UsernameIsExist(registerRequest.Username) {
		return entities.Auth{}, errors.New("USERNAME ALREADY EXIST")
	}

	hashPassword, hashErr := helpers.HashPassword(registerRequest.Password)
	if hashErr != nil {
		return entities.Auth{}, hashErr
	}

	var role entities.Role
	if isAdmin {
		role = entities.Admin
	} else {
		role = entities.User
	}

	authData := entities.Auth{
		Username: registerRequest.Username,
		Email:    registerRequest.Email,
		Role:     role,
		Password: hashPassword,
	}

	return service.authRepository.Create(authData)
}
