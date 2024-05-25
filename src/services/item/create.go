package itemservice

import "belimang/src/entities"

func (s *itemService) Create(item entities.MerchantItemRequest) (entities.MerchantItem, error) {
	merchantItem, err := s.itemRepository.Create(item)
	if err != nil {
		return entities.MerchantItem{}, err
	}
	return merchantItem, nil
}
