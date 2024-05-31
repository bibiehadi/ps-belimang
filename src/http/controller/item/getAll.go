package itemcontroller

import (
	"belimang/src/entities"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func (controller *itemController) GetAllItem(c echo.Context) error {
	merchantId := c.Param("merchantId")
	var params entities.MerchantItemQueryParams
	bindError := c.Bind(&params)

	fmt.Println(params)
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

	limitStr := c.QueryParam("limit")
	if limitStr != "" {
		limit, err := strconv.Atoi(limitStr)
		if err == nil && limit > 0 {
			params.Limit = limit
		} else {
			return c.JSON(http.StatusBadRequest, entities.ErrorResponse{
				Status:  false,
				Message: "Invalid limit parameter",
			})
		}
	} else {
		params.Limit = 5
	}

	offsetStr := c.QueryParam("offset")
	if offsetStr != "" {
		offset, err := strconv.Atoi(offsetStr)
		if err == nil && offset >= 0 {
			params.Offset = offset
		} else {
			return c.JSON(http.StatusBadRequest, entities.ErrorResponse{
				Status:  false,
				Message: "Invalid offset parameter",
			})
		}
	} else {
		params.Offset = 0
	}

	if err := controller.validator.Struct(params); err != nil {
		var validationErrors []string
		for _, err := range err.(validator.ValidationErrors) {
			validationErrors = append(validationErrors, fmt.Sprintf("%s is %s", err.Field(), err.Tag()))
		}
		return c.JSON(http.StatusBadRequest, entities.ErrorResponse{
			Status:  false,
			Message: validationErrors,
		})
	}
	params.MerchantId = merchantId
	items, meta, err := controller.ItemService.GetAll(params)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, entities.ErrorResponse{
			Status:  false,
			Message: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, entities.SuccessGetAllResponse{
		Message: "Success",
		Data:    items,
		Meta:    meta,
	})
}
