package merchantservice

import (
	"belimang/src/entities"
	merchantRepository "belimang/src/repositories/merchant"
)

type MerchantService interface {
	Create(request entities.MerchantRequest) (string, error)
	FindAll(params entities.MerchantQueryParams) ([]entities.MerchantResponse, entities.MerchantMetaResponse, error)
	FindNearby(params entities.MerchantQueryParams) ([]entities.GetOrder, entities.MerchantMetaResponse, error)
}

type merchantService struct {
	merchantRepository merchantRepository.MerchantRepository
}

func New(repository merchantRepository.MerchantRepository) *merchantService {
	return &merchantService{merchantRepository: repository}
}
