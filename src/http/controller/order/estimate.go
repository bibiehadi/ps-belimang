package orderController

import (
	"belimang/src/entities"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func (controller *orderController) Estimate(c echo.Context) error {
	var estimateRequest entities.EstimateRequest
	bindError := c.Bind(&estimateRequest)

	if bindError != nil {
		return c.JSON(http.StatusBadRequest, entities.ErrorResponse{
			Status:  false,
			Message: bindError.Error(),
		})
	}

	if err := controller.validator.Struct(estimateRequest); err != nil {
		var validationErrors []string
		for _, err := range err.(validator.ValidationErrors) {
			validationErrors = append(validationErrors, fmt.Sprintf("%s is %s", err.Field(), err.Tag()))
		}
		return c.JSON(http.StatusBadRequest, entities.ErrorResponse{
			Status:  false,
			Message: validationErrors,
		})
	}

	var count int = 0
	for _, element := range estimateRequest.Orders {

		if err := controller.validator.Struct(element); err != nil {
			var validationErrors []string
			for _, err := range err.(validator.ValidationErrors) {
				validationErrors = append(validationErrors, fmt.Sprintf("%s is %s", err.Field(), err.Tag()))
			}
			return c.JSON(http.StatusBadRequest, entities.ErrorResponse{
				Status:  false,
				Message: validationErrors,
			})
		}

		if *element.IsStartingPoint {
			count++
		}

		if count > 1 {
			return c.JSON(http.StatusBadRequest, entities.ErrorResponse{
				Status:  false,
				Message: "Starting Point more than one",
			})
		}

		for _, item := range element.Items {
			if err := controller.validator.Struct(item); err != nil {
				var validationErrors []string
				for _, err := range err.(validator.ValidationErrors) {
					validationErrors = append(validationErrors, fmt.Sprintf("%s is %s", err.Field(), err.Tag()))
				}
				return c.JSON(http.StatusBadRequest, entities.ErrorResponse{
					Status:  false,
					Message: validationErrors,
				})
			}
		}
	}

	order, err := controller.OrderService.Estimate(estimateRequest)
	if err != nil {
		return c.JSON(http.StatusBadRequest, entities.ErrorResponse{
			Status:  false,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, entities.SuccessResponse{
		Message: "Order successfull",
		Data:    order,
	})
}
