package merchantRepository

import (
	"belimang/src/entities"
	"belimang/src/helpers"
	"context"
	"fmt"
	"strconv"
	"strings"
)

func (r *merchantRepository) FindNearby(params entities.MerchantQueryParams) ([]entities.MerchantNearbyResponse, entities.MerchantMetaResponse, error) {
	// query := `
	//    WITH limited_merchants AS (
	// 	SELECT id, name, merchant_category, image_url, latitude, longitude,
	// 	2 * 6371 * asin( |/( sin(($1*(pi()/180)-latitude*(pi()/180))/2::decimal)^2 + (sin(($2*(pi()/180)-longitude*(pi()/180))/2::decimal)^2) * cos(latitude*(pi()/180)) * cos(5*(pi()/180)) ) ) AS distance, created_at
	// 	FROM merchants
	// 	// WHERE distance <= 3.00
	// 	// ORDER BY distance ASC
	// 	LIMIT $3 OFFSET $4
	//    )
	//    SELECT m.id, m.name, m.merchant_category, m.image_url, m.latitude, m.longitude, distance, m.created_at,
	//    mi.id, mi.name, mi.product_category, mi.price, mi.image_url, mi.created_at
	//    from limited_merchants
	//    JOIN merchants AS  m ON limited_merchants.id = m.id
	//    JOIN merchant_items AS mi ON m.id = mi.merchant_id
	//    WHERE 1=1
	// `
	query := `
		SELECT id, name, merchant_category, image_url, latitude, longitude, 
		2 * 6371 * asin( |/( sin(($1*(pi()/180)-latitude*(pi()/180))/2::decimal)^2 + (sin(($2*(pi()/180)-longitude*(pi()/180))/2::decimal)^2) * cos(latitude*(pi()/180)) * cos(5*(pi()/180)) ) ) AS distance, created_at 
		FROM merchants
	    WHERE 1=1
	   `

	if params.MerchantID != "" {
		query += fmt.Sprintf(" AND id = '%s'", params.MerchantID)
	}

	if params.Name != "" {
		query += fmt.Sprintf(" AND LOWER(name) LIKE '%%%s%%'", strings.ToLower(params.Name))
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
			query += fmt.Sprintf(" AND merchant_category = '%s'", params.MerchantCategory)
		}
	}

	// if params.CreatedAt != "" {
	// 	if params.CreatedAt == "asc" || params.CreatedAt == "desc" {
	query += fmt.Sprintf(" ORDER BY distance %s", "asc")
	// 	}
	// }
	query += " LIMIT " + strconv.Itoa(params.Limit) + " OFFSET " + strconv.Itoa(params.Offset)

	rows, err := r.db.Query(context.Background(), query, params.Lat, params.Long)
	fmt.Println(query)
	if err != nil {
		fmt.Println(err)
		return []entities.MerchantNearbyResponse{}, entities.MerchantMetaResponse{}, err
	}

	defer rows.Close()
	var Merchants []entities.MerchantNearbyResponse
	// merchantMap := make(map[string]*entities.MerchantNearbyResponse)
	for rows.Next() {
		var merchant entities.Merchant
		var listItem []entities.MerchantItemResponse = []entities.MerchantItemResponse{}
		var distance float64
		err := rows.Scan(&merchant.ID, &merchant.Name, &merchant.MerchantCategory, &merchant.ImageURL, &merchant.Latitude, &merchant.Longitude, &distance, &merchant.CreatedAt)

		if err != nil {
			return []entities.MerchantNearbyResponse{}, entities.MerchantMetaResponse{}, err
		}

		fmt.Printf("User Lat : %f, long: %f ", params.Lat, params.Long)
		fmt.Printf("User Lat : %f, long: %f ", merchant.Latitude, merchant.Longitude)
		fmt.Println("")
		fmt.Println("haversine : ", helpers.Haversine(entities.Location{
			Lat:  params.Lat,
			Long: params.Long,
		}, entities.Location{
			Lat:  merchant.Latitude,
			Long: merchant.Longitude,
		}))

		itemQuery := `
		SELECT 
		id, name, product_category, price, image_url, created_at
		from merchant_items
		WHERE merchant_id = $1 
		ORDER BY created_at ASC
		`

		itemRows, err := r.db.Query(context.Background(), itemQuery, merchant.ID)

		// fmt.Println("Query: ", query)
		if err != nil {
			fmt.Println(err)
			return []entities.MerchantNearbyResponse{}, entities.MerchantMetaResponse{}, err
		}

		for itemRows.Next() {
			var item entities.MerchantItemResponse
			err := itemRows.Scan(&item.ItemId, &item.Name, &item.ProductCategory, &item.Price, &item.ImageUrl, &item.CreatedAt)

			if err != nil {
				return []entities.MerchantNearbyResponse{}, entities.MerchantMetaResponse{}, err
			}
			listItem = append(listItem, entities.MerchantItemResponse{
				ItemId:          item.ItemId,
				Name:            item.Name,
				ProductCategory: item.ProductCategory,
				Price:           item.Price,
				ImageUrl:        item.ImageUrl,
				CreatedAt:       item.CreatedAt,
			},
			)
		}
		fmt.Println(listItem)

		merchantResponse := &entities.MerchantNearbyResponse{
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
			Items: listItem,
		}
		Merchants = append(Merchants, *merchantResponse)

	}

	var metaQuery string = "SELECT COUNT(*) FROM merchants"
	metaRows, err := r.db.Query(context.Background(), metaQuery)
	if err != nil {
		return []entities.MerchantNearbyResponse{}, entities.MerchantMetaResponse{}, err
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
