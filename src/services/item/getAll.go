package itemservice

import (
	"belimang/src/entities"
	"fmt"
)

func (s *itemService) GetAll(params entities.MerchantItemQueryParams) ([]entities.MerchantItemResponse, entities.MerchantItemMetaResponse, error) {
	merchantItems, meta, err := s.itemRepository.GetAll(params)
	if err != nil {
		return []entities.MerchantItemResponse{}, entities.MerchantItemMetaResponse{}, err
	}
	fmt.Println(merchantItems)
	return merchantItems, meta, nil
}
