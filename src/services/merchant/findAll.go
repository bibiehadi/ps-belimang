package merchantservice

import "belimang/src/entities"

func (s *merchantService) FindAll(params entities.MerchantQueryParams) ([]entities.MerchantResponse, entities.MerchantMetaResponse, error) {
	merchants, meta, err := s.merchantRepository.FindAll(params)
	if err != nil {
		return nil, entities.MerchantMetaResponse{}, err
	}
	return merchants, meta, nil
}
