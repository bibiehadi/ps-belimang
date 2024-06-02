package orderService

import "belimang/src/entities"

func (s *orderService) FindAll(params entities.OrderQueryParams) ([]entities.GetOrderResponse, error) {
	orders, err := s.orderRepository.FindAll(params)
	if err != nil {
		return []entities.GetOrderResponse{}, err
	}
	return orders, nil
}
