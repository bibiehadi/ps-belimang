package merchantservice

import "belimang/src/entities"

func (s *merchantService) FindNearby(params entities.MerchantQueryParams) ([]entities.MerchantResponse, entities.MerchantMetaResponse, error) {
	merchants, meta, err := s.merchantRepository.FindNearby(params)
	if err != nil {
		return nil, entities.MerchantMetaResponse{}, err
	}
	return merchants, meta, nil
}
