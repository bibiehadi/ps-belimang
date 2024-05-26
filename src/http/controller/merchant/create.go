package merchantController

import (
	"belimang/src/entities"
	"belimang/src/helpers"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"regexp"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func (controller *merchantController) Create(c echo.Context) error {
	var merchantRequest entities.MerchantRequest
	bindError := c.Bind(&merchantRequest)

	if bindError != nil {
		switch bindError.(type) {
		case validator.ValidationErrors:
			var errorMessages string
			for _, e := range bindError.(validator.ValidationErrors) {
				errorMessage := fmt.Sprintf("Error on field: %s, condition: %s", e.Field(), e.ActualTag())
				errorMessages = fmt.Sprintf(errorMessages + errorMessage)
			}
			return c.JSON(
				http.StatusBadRequest,
				entities.ErrorResponse{
					Status:  false,
					Message: errorMessages,
				},
			)

		case *json.UnmarshalTypeError:
			return c.JSON(http.StatusBadRequest, entities.ErrorResponse{
				Status:  false,
				Message: bindError.Error(),
			})

		default:
			if bindError == io.EOF {
				return c.JSON(http.StatusBadRequest, entities.ErrorResponse{
					Status:  false,
					Message: "Request body is empty",
				})

			}
			return c.JSON(http.StatusBadRequest, entities.ErrorResponse{
				Status:  false,
				Message: bindError.Error(),
			})

		}
	}

	err := validateMerchant(merchantRequest)
	if err != nil {
		return c.JSON(http.StatusBadRequest, entities.ErrorResponse{
			Status:  false,
			Message: err.Error(),
		})
	}

	if err := controller.validator.Struct(merchantRequest); err != nil {
		var validationErrors []string
		for _, err := range err.(validator.ValidationErrors) {
			validationErrors = append(validationErrors, fmt.Sprintf("%s is %s", err.Field(), err.Tag()))
		}
		return c.JSON(http.StatusBadRequest, entities.ErrorResponse{
			Status:  false,
			Message: validationErrors,
		})
	}

	if !helpers.ValidateUrl(merchantRequest.ImageURL) {
		c.JSON(http.StatusBadRequest, entities.ErrorResponse{
			Status:  false,
			Message: "URL FORMAT IS NOT VALID",
		})
		return nil
	}

	merchant, err := controller.MerchantService.Create(merchantRequest)
	if err != nil {
		return c.JSON(http.StatusBadRequest, entities.ErrorResponse{
			Status:  false,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, entities.SuccessResponse{
		Message: "Merchant registered successfull",
		Data:    merchant,
	})
}

func validateMerchant(merchant entities.MerchantRequest) error {
	// Validate Name
	if len(merchant.Name) < 2 || len(merchant.Name) > 30 {
		return errors.New("name must be between 2 and 30 characters")
	}

	// Validate MerchantCategory
	validCategories := []string{"SmallRestaurant", "MediumRestaurant", "LargeRestaurant", "MerchandiseRestaurant", "BoothKiosk", "ConvenienceStore"}
	isValidCategory := false
	for _, category := range validCategories {
		if merchant.MerchantCategory == category {
			isValidCategory = true
			break
		}
	}
	if !isValidCategory {
		return errors.New("merchantCategory must be one of the valid categories")
	}

	// Validate ImageURL
	matched, err := regexp.MatchString(`https?:\/\/(www\.)?[-a-zA-Z0-9@:%._\+~#=]{1,256}\.[a-zA-Z0-9()]{1,6}\b([-a-zA-Z0-9()@:%_\+.~#?&//=]*)`, merchant.ImageURL)
	if err != nil || !matched {
		return errors.New("imageUrl must be a valid image URL")
	}

	// Validate Location
	if merchant.Location.Lat == 0 || merchant.Location.Long == 0 {
		return errors.New("latitude and longitude cannot be zero")
	}

	return nil
}
