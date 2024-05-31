package authController

import (
	"belimang/src/entities"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func (controller *authController) Login(c echo.Context) error {
	var loginRequest entities.LoginRequest
	err := c.Bind(&loginRequest)

	if err != nil {
		switch err.(type) {
		case validator.ValidationErrors:
			var errorMessages string
			for _, e := range err.(validator.ValidationErrors) {
				errorMessage := fmt.Sprintf("Error on field: %s, condition: %s", e.Field(), e.ActualTag())
				errorMessages = fmt.Sprintf(errorMessages + errorMessage)
			}
			c.JSON(
				http.StatusBadRequest,
				entities.ErrorResponse{
					Status:  false,
					Message: errorMessages,
				},
			)
			return nil
		case *json.UnmarshalTypeError:
			c.JSON(http.StatusBadRequest, entities.ErrorResponse{
				Status:  false,
				Message: err.Error(),
			})
			return nil

		default:
			if err == io.EOF {
				c.JSON(http.StatusBadRequest, entities.ErrorResponse{
					Status:  false,
					Message: "Request body is empty",
				})
				return nil
			}
			c.JSON(http.StatusBadRequest, entities.ErrorResponse{
				Status:  false,
				Message: err.Error(),
			})
			return nil
		}
	}

	if err := controller.validator.Struct(loginRequest); err != nil {
		var validationErrors []string
		for _, err := range err.(validator.ValidationErrors) {
			validationErrors = append(validationErrors, fmt.Sprintf("%s is %s", err.Field(), err.Tag()))
		}
		return c.JSON(http.StatusBadRequest, entities.ErrorResponse{
			Status:  false,
			Message: validationErrors,
		})
	}

	tokenString, _, err := controller.AuthService.Login(loginRequest)
	if err != nil {
		if err.Error() == "INVALID USERNAME OR PASSWORD" {
			return c.JSON(http.StatusNotFound, entities.ErrorResponse{
				Status:  false,
				Message: err.Error(),
			})

		}
		return c.JSON(http.StatusInternalServerError, entities.ErrorResponse{
			Status:  false,
			Message: "Internal server error",
		})
	}

	return c.JSON(http.StatusOK, entities.SuccessResponse{
		Message: "User login successfully",
		Data: entities.AuthResponse{
			Token: tokenString,
		},
	})
}
