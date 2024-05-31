package orderController

import (
	"belimang/src/entities"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func (controller *orderController) Order(c echo.Context) error {
	var estimatedId entities.OrderRequest
	// userID, _ := utils.GetUserIDFromJWTClaims(c)
	if err := c.Bind(&estimatedId); err != nil {
		return c.JSON(http.StatusBadRequest, entities.ErrorResponse{
			Status:  false,
			Message: "Invalid request body",
		})
	}

	// Validasi input menggunakan validator
	if err := controller.validator.Struct(estimatedId); err != nil {
		var validationErrors []string
		for _, err := range err.(validator.ValidationErrors) {
			validationErrors = append(validationErrors, fmt.Sprintf("%s is %s", err.Field(), err.Tag()))
		}
		return c.JSON(http.StatusBadRequest, entities.ErrorResponse{
			Status:  false,
			Message: validationErrors,
		})
	}

	orderId, err := controller.OrderService.Order(estimatedId.OrderId)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, entities.ErrorResponse{
			Status:  false,
			Message: "Failed to update product",
		})
	}

	return c.JSON(http.StatusOK, entities.SuccessResponse{
		Message: "Product updated successfully",
		Data: entities.OrderResponse{
			OrderId: orderId,
		},
	})
}
