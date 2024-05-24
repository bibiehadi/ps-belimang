package authController

import (
	authService "belimang/src/services/auth"

	"github.com/go-playground/validator/v10"
)

type authController struct {
	authService.AuthService
	validator *validator.Validate
}

func New(services authService.AuthService) *authController {
	validate := validator.New()
	return &authController{AuthService: services, validator: validate}
}
