package merchantservice

import (
	"belimang/src/entities"
	merchantRepository "belimang/src/repositories/merchant"
)

type MerchantService interface {
	Create(merchant entities.MerchantRequest) (entities.Merchant, error)
	FindAll(params entities.MerchantQueryParams) ([]entities.MerchantResponse, entities.MerchantMetaResponse, error)
	FindNearby(params entities.MerchantQueryParams) ([]entities.MerchantResponse, entities.MerchantMetaResponse, error)
}

type merchantService struct {
	merchantRepository merchantRepository.MerchantRepository
}

func New(repository merchantRepository.MerchantRepository) *merchantService {
	return &merchantService{merchantRepository: repository}
}
