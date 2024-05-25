package itemcontroller

import (
	itemService "belimang/src/services/item"
	"github.com/go-playground/validator/v10"
)

type itemController struct {
	itemService.ItemService
	validator *validator.Validate
}

func New(services itemService.ItemService) *itemController {
	validate := validator.New()
	return &itemController{ItemService: services, validator: validate}
}
