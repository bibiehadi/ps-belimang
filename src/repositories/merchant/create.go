package merchantRepository

import (
	"belimang/src/entities"
	"context"
)

func (r *merchantRepository) Create(merchant entities.Merchant) (entities.Merchant, error) {
	var query string = "INSERT INTO merchants (name, merchant_category, image_url, latitude, longitude) values ($1,$2,$3,$4,$5) RETURNING id"
	var merchantId int
	err := r.db.QueryRow(context.Background(), query, merchant.Name, merchant.MerchantCategory, merchant.ImageURL, merchant.Latitude, merchant.Longitude).Scan(
		&merchantId,
	)
	if err != nil {
		return entities.Merchant{}, err
	}
	merchant.ID = merchantId
	return merchant, err
}
