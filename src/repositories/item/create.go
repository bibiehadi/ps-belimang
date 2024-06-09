package itemRepository

import (
	"belimang/src/entities"
	"context"

	"github.com/google/uuid"
)

func (r itemRepository) Create(item entities.MerchantItemRequest) (entities.MerchantItem, error) {
	var id = uuid.NewString()
	var query string = "INSERT INTO  merchant_items (id,name,product_category,price,image_url,merchant_id) values ($1,$2,$3,$4,$5,$6) RETURNING id"
	var itemId string
	err := r.db.QueryRow(context.Background(), query, id, item.Name, item.ProductCategory, item.Price, item.ImageURL, item.MerchantID).Scan(
		&itemId,
	)
	if err != nil {
		return entities.MerchantItem{}, err
	}
	medicalRecord := entities.MerchantItem{
		ID: itemId,
	}
	return medicalRecord, err
}
