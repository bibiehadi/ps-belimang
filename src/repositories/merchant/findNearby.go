package merchantRepository

import (
	"belimang/src/entities"
	"context"
	"fmt"
	"strconv"
	"strings"
)

func (r *merchantRepository) FindNearby(params entities.MerchantQueryParams) ([]entities.GetOrder, entities.MerchantMetaResponse, error) {
	query := `
	   WITH limited_merchants AS (
		SELECT id, name, merchant_category, image_url, latitude, longitude, 
		2 * 6371 * asin( |/( sin(($1*(pi()/180)-latitude*(pi()/180))/2::decimal)^2 + (sin(($2*(pi()/180)-longitude*(pi()/180))/2::decimal)^2) * cos(latitude*(pi()/180)) * cos(5*(pi()/180)) ) ) AS distance, created_at 
		FROM merchants
		LIMIT $3 OFFSET $4
	   )
	   SELECT m.id, m.name, m.merchant_category, m.image_url, m.latitude, m.longitude, distance, m.created_at,
	   mi.id, mi.name, mi.product_category, mi.price, mi.image_url, mi.created_at
	   from limited_merchants
	   JOIN merchants AS  m ON limited_merchants.id = m.id
	   JOIN merchant_items AS mi ON m.id = mi.merchant_id
	   WHERE 1=1 
	`

	if params.MerchantID != "" {
		query += fmt.Sprintf(" AND id = '%s'", params.MerchantID)
	}

	if params.Name != "" {
		query += fmt.Sprintf(" AND LOWER(m.name) LIKE '%%%s%%'", strings.ToLower(params.Name))
	}

	if params.MerchantCategory != "" {
		validCategories := map[string]bool{
			"SmallRestaurant":       true,
			"MediumRestaurant":      true,
			"LargeRestaurant":       true,
			"MerchandiseRestaurant": true,
			"BoothKiosk":            true,
			"ConvenienceStore":      true,
		}
		if validCategories[params.MerchantCategory] {
			query += fmt.Sprintf(" AND m.merchant_category = '%s'", params.MerchantCategory)
		}
	}

	// if params.CreatedAt != "" {
	// 	if params.CreatedAt == "asc" || params.CreatedAt == "desc" {
	query += fmt.Sprintf(" ORDER BY distance %s", params.CreatedAt)
	// 	}
	// }

	rows, err := r.db.Query(context.Background(), query, params.Lat, params.Long, strconv.Itoa(params.Limit), strconv.Itoa(params.Offset))
	if err != nil {
		return []entities.GetOrder{}, entities.MerchantMetaResponse{}, err
	}

	defer rows.Close()
	var Merchants []entities.GetOrder
	merchantMap := make(map[int]*entities.GetOrder)
	for rows.Next() {
		var merchant entities.Merchant
		var item entities.MerchantItemResponse
		var distance float64
		err := rows.Scan(&merchant.ID, &merchant.Name, &merchant.MerchantCategory, &merchant.ImageURL, &merchant.Latitude, &merchant.Longitude, &distance, &merchant.CreatedAt, &item.ItemId, &item.Name, &item.ProductCategory, &item.Price, &item.ImageUrl, &item.CreatedAt)

		if err != nil {
			return []entities.GetOrder{}, entities.MerchantMetaResponse{}, err
		}

		_, merchantExist := merchantMap[merchant.ID]
		if !merchantExist {
			merchantResponse := &entities.GetOrder{
				Merchant: entities.MerchantResponse{
					MerchantId:       merchant.ID,
					Name:             merchant.Name,
					MerchantCategory: merchant.MerchantCategory,
					ImageURL:         merchant.ImageURL,
					Location: entities.Location{
						Lat:  merchant.Latitude,
						Long: merchant.Longitude,
					},
					Distance:  distance,
					CreatedAt: merchant.CreatedAt,
				},
				Items: []entities.OrderItemResponse{},
			}
			Merchants = append(Merchants, *merchantResponse)
			merchantMap[merchant.ID] = &Merchants[len(Merchants)-1]
		}
		merchantMap[merchant.ID].Items = append(merchantMap[merchant.ID].Items, entities.OrderItemResponse{
			ItemId:          item.ItemId,
			Name:            item.Name,
			ProductCategory: item.ProductCategory,
			Price:           item.Price,
			ImageUrl:        item.ImageUrl,
			CreatedAt:       item.CreatedAt,
		},
		)
	}

	var metaQuery string = "SELECT COUNT(*) FROM merchants"
	metaRows, err := r.db.Query(context.Background(), metaQuery)
	if err != nil {
		return []entities.GetOrder{}, entities.MerchantMetaResponse{}, err
	}
	defer metaRows.Close()
	var meta entities.MerchantMetaResponse
	if metaRows.Next() {
		if err := metaRows.Scan(&meta.Total); err != nil {
			return nil, entities.MerchantMetaResponse{}, err
		}
	}
	metaRows.Scan(&meta.Total)
	meta.Limit = params.Limit
	meta.Offset = params.Offset

	return Merchants, meta, nil
}
