package merchantservice

import (
	"belimang/src/entities"
	merchantRepository "belimang/src/repositories/merchant"
)

type MerchantService interface {
	Create(merchant entities.MerchantRequest) (entities.Merchant, error)
}

type merchantService struct {
	merchantRepository merchantRepository.MerchantRepository
}

func New(repository merchantRepository.MerchantRepository) *merchantService {
	return &merchantService{merchantRepository: repository}
}
