package merchantController

import (
	"belimang/src/entities"
	"net/http"
	"reflect"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

func (controller *merchantController) FindNearby(c echo.Context) error {
	latlong := strings.Split(c.Param("latlong"), ",")
	if len(latlong) != 2 {
		return c.JSON(http.StatusBadRequest, entities.ErrorResponse{
			Status:  false,
			Message: "Invalide coordinates",
		})
	}
	lat, _ := strconv.ParseFloat(latlong[0], 64)
	long, _ := strconv.ParseFloat(latlong[1], 64)

	params := entities.MerchantQueryParams{}
	params.Lat = lat
	params.Long = long
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

	merchants, meta, err := controller.MerchantService.FindNearby(params)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, entities.ErrorResponse{
			Status:  false,
			Message: "Failed to fetch merchants",
		})
	}

	if merchants == nil || reflect.ValueOf(merchants).IsNil() {
		return c.JSON(http.StatusOK, entities.SuccessGetAllResponse{
			// Message: "success",
			Data: []entities.Merchant{},
			Meta: entities.MerchantMetaResponse{},
		})
	}

	return c.JSON(http.StatusOK, entities.SuccessGetAllResponse{
		// Message: "success",
		Data: merchants,
		Meta: meta,
	})
}
