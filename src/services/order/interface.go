package orderService

import (
	"belimang/src/entities"
	itemRepository "belimang/src/repositories/item"
	merchantRepository "belimang/src/repositories/merchant"
)

type OrderService interface {
	Estimate(estimateRequest entities.EstimateRequest) (entities.EstimateResponse, error)
}

type orderService struct {
	merchantRepository merchantRepository.MerchantRepository
	itemRepository     itemRepository.ItemRepository
}

func New(merchantRepository merchantRepository.MerchantRepository, itemRepository itemRepository.ItemRepository) *orderService {
	return &orderService{merchantRepository, itemRepository}
}
