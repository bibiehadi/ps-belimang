package orderService

import (
	"belimang/src/entities"
	"belimang/src/helpers"
	"errors"
	"fmt"
	"sort"
)

func (service *orderService) Estimate(estimateRequest entities.EstimateRequest, userId string) (entities.EstimateResponse, error) {

	listLocation := make([]entities.Location, 0, len(estimateRequest.Orders)+1)
	listLocation = append(listLocation, entities.Location{
		Lat:  estimateRequest.UserLocation.Lat,
		Long: estimateRequest.UserLocation.Long,
	})
	var isStartingPoint int = 0
	sort.Slice(estimateRequest.Orders, func(i, j int) bool {
		if *estimateRequest.Orders[i].IsStartingPoint {
			isStartingPoint++
		}
		return *estimateRequest.Orders[i].IsStartingPoint
	})

	if isStartingPoint > 1 || isStartingPoint == 0 {
		return entities.EstimateResponse{}, errors.New("starting point must have 1")

	}

	var totalPrice float64 = 0.0

	for _, merch := range estimateRequest.Orders {
		merchant, err := service.merchantRepository.FindById(merch.MerchantId)
		if err != nil {
			return entities.EstimateResponse{}, err
		}
		merchLocation := entities.Location{Lat: merchant.Latitude, Long: merchant.Longitude}
		fmt.Println("user Location : ", estimateRequest.UserLocation)
		fmt.Println("merch Location : ", merchLocation)
		fmt.Println("Distance : ", helpers.Haversine(estimateRequest.UserLocation, merchLocation))
		// if helpers.Haversine(estimateRequest.UserLocation, merchLocation) > 3.00 {
		// 	return entities.EstimateResponse{}, errors.New("MERCHANT LOCATION MORE THAN 3 KM")
		// }

		for _, orderItem := range merch.Items {
			item, err := service.itemRepository.FindById(orderItem.ItemId)
			if err != nil {
				return entities.EstimateResponse{}, err
			}
			totalPrice += float64(orderItem.Quantity * item.Price)
		}
		listLocation = append(listLocation, merchLocation)
	}

	_, totalDistance := helpers.NearestNeighbor(listLocation)

	estDeliveryTime := totalDistance / (40.0 / 60.0)
	deliveryFee := totalDistance * 10000
	orderId, err := service.orderRepository.Create(
		estimateRequest, estDeliveryTime, totalDistance, totalPrice, deliveryFee, userId,
	)

	if err != nil {
		return entities.EstimateResponse{}, err
	}

	return entities.EstimateResponse{
		TotalPrice:           deliveryFee + totalPrice,
		EstimateDeliveryTime: estDeliveryTime,
		EstimateId:           orderId,
	}, nil
}
