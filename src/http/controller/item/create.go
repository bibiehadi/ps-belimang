package itemcontroller

import (
	"belimang/src/entities"
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"io"
	"net/http"
)

func (controller *itemController) CreateItem(c echo.Context) error {
	merchantId := c.Param("merchantId")
	var itemRequest entities.MerchantItemRequest
	bindError := c.Bind(&itemRequest)

	fmt.Println(itemRequest)
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

	if err := controller.validator.Struct(itemRequest); err != nil {
		var validationErrors []string
		for _, err := range err.(validator.ValidationErrors) {
			validationErrors = append(validationErrors, fmt.Sprintf("%s is %s", err.Field(), err.Tag()))
		}
		return c.JSON(http.StatusBadRequest, entities.ErrorResponse{
			Status:  false,
			Message: validationErrors,
		})
	}

	itemRequest.MerchantID = merchantId
	item, err := controller.ItemService.Create(itemRequest)
	if err != nil {
		return c.JSON(http.StatusBadRequest, entities.ErrorResponse{
			Status:  false,
			Message: err.Error(),
		})
	}
	return c.JSON(http.StatusCreated, entities.SuccessResponse{
		Message: "Item Created successfull",
		Data: entities.CreateMerchantItemResponse{
			ItemId: item.ID,
		},
	})

}
