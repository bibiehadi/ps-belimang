package orderController

import (
	orderService "belimang/src/services/order"

	"github.com/go-playground/validator/v10"
)

type orderController struct {
	orderService.OrderService
	validator *validator.Validate
}

func New(services orderService.OrderService) *orderController {
	validate := validator.New()
	return &orderController{OrderService: services, validator: validate}
}
