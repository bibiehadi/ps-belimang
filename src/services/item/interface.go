package itemservice

import (
	"belimang/src/entities"
	itemRepository "belimang/src/repositories/item"
)

type ItemService interface {
	Create(item entities.MerchantItemRequest) (entities.MerchantItem, error)
	GetAll(params entities.MerchantItemQueryParams) ([]entities.MerchantItemResponse, entities.MerchantItemMetaResponse, error)
}

type itemService struct {
	itemRepository itemRepository.ItemRepository
}

func New(repository itemRepository.ItemRepository) *itemService {
	return &itemService{itemRepository: repository}
}
