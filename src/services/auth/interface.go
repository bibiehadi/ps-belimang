package authservice

import (
	"belimang/src/entities"
	authRepository "belimang/src/repositories/auth"
)

type AuthService interface {
	Register(registerRequest entities.RegisterRequest, isAdmin bool) (entities.Auth, error)
	Login(loginRequest entities.LoginRequest) (string, entities.Auth, error)
}

type authService struct {
	authRepository authRepository.AuthRepository
}

func New(repository authRepository.AuthRepository) *authService {
	return &authService{authRepository: repository}
}
