package orderController

import (
	"belimang/src/entities"
	"fmt"
	"net/http"
	"reflect"
	"strconv"

	"github.com/labstack/echo/v4"
)

func (controller *orderController) FindAll(c echo.Context) error {
	params := entities.OrderQueryParams{}
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

	if merchantId := c.QueryParam("merchantId"); merchantId != "" {
		params.MerchantID = merchantId
	}

	if name := c.QueryParam("name"); name != "" {
		params.Name = name
	}

	if merchantCategory := c.QueryParam("merchantCategory"); merchantCategory != "" {
		params.MerchantCategory = merchantCategory
	}

	if createdAt := c.QueryParam("createdAt"); createdAt != "" {
		if createdAt != "asc" && createdAt != "desc" {
		} else {
			params.CreatedAt = createdAt
		}
	}

	orders, err := controller.OrderService.FindAll(params)
	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusInternalServerError, entities.ErrorResponse{
			Status:  false,
			Message: "Failed to fetch orders",
		})
	}

	if orders == nil || reflect.ValueOf(orders).IsNil() {
		return c.JSON(http.StatusOK, entities.SuccessGetAllResponse{
			Message: "success",
			Data:    []entities.GetOrderResponse{},
		})
	}

	return c.JSON(http.StatusOK, entities.SuccessGetAllResponse{
		Message: "success",
		Data:    orders,
	})
}
