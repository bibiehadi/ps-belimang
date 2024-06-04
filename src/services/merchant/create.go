package merchantservice

import "belimang/src/entities"

func (s *merchantService) Create(request entities.MerchantRequest) (string, error) {

	merchant := entities.Merchant{
		Name:             request.Name,
		MerchantCategory: request.MerchantCategory,
		ImageURL:         request.ImageURL,
		Latitude:         request.Location.Lat,
		Longitude:        request.Location.Long,
	}

	return s.merchantRepository.Create(merchant)
}
