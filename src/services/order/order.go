package orderService

func (service *orderService) Order(estimateId string) (string, error) {

	_, err := service.orderRepository.FindById(estimateId)
	if err != nil {
		return "", err
	}

	orderId, err := service.orderRepository.Order(estimateId)

	if err != nil {
		return "", err
	}

	return orderId, nil
}
