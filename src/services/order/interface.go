package orderService

import (
	"belimang/src/entities"
	itemRepository "belimang/src/repositories/item"
	merchantRepository "belimang/src/repositories/merchant"
	orderRepository "belimang/src/repositories/order"
)

type OrderService interface {
	Estimate(estimateRequest entities.EstimateRequest, userId string) (entities.EstimateResponse, error)
	Order(estimateId string) (string, error)
}

type orderService struct {
	merchantRepository merchantRepository.MerchantRepository
	itemRepository     itemRepository.ItemRepository
	orderRepository    orderRepository.OrderResorderRepository
}

func New(merchantRepository merchantRepository.MerchantRepository, itemRepository itemRepository.ItemRepository, orderRepository orderRepository.OrderResorderRepository) *orderService {
	return &orderService{merchantRepository, itemRepository, orderRepository}
}
