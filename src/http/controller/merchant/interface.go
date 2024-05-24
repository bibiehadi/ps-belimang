package merchantController

import (
	merchantService "belimang/src/services/merchant"

	"github.com/go-playground/validator/v10"
)

type merchantController struct {
	merchantService.MerchantService
	validator *validator.Validate
}

func New(services merchantService.MerchantService) *merchantController {
	validate := validator.New()
	return &merchantController{MerchantService: services, validator: validate}
}
