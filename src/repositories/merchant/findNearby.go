package merchantRepository

import (
	"belimang/src/entities"
	"belimang/src/helpers"
	"context"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func (r *merchantRepository) FindNearby(params entities.MerchantQueryParams) ([]entities.MerchantResponse, entities.MerchantMetaResponse, error) {
	query := "SELECT id, name, merchant_category, image_url, latitude, longitude, created_at FROM merchants WHERE 1=1"

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

	if params.CreatedAt != "" {
		if params.CreatedAt == "asc" || params.CreatedAt == "desc" {
			query += fmt.Sprintf(" ORDER BY created_at %s", params.CreatedAt)
		}
	}

	query += " LIMIT " + strconv.Itoa(params.Limit) + " OFFSET " + strconv.Itoa(params.Offset)
	rows, err := r.db.Query(context.Background(), query)

	if err != nil {
		fmt.Println(err.Error())
		return []entities.MerchantResponse{}, entities.MerchantMetaResponse{}, err
	}

	defer rows.Close()
	var Merchants []entities.MerchantResponse
	for rows.Next() {
		var merchant entities.Merchant
		err := rows.Scan(&merchant.ID, &merchant.Name, &merchant.MerchantCategory, &merchant.ImageURL, &merchant.Latitude, &merchant.Longitude, &merchant.CreatedAt)
		if err != nil {
			return []entities.MerchantResponse{}, entities.MerchantMetaResponse{}, err
		}
		mr := entities.MerchantResponse{
			MerchantId:       merchant.ID,
			Name:             merchant.Name,
			MerchantCategory: merchant.MerchantCategory,
			ImageURL:         merchant.ImageURL,
			Location: entities.Location{
				Lat:  merchant.Latitude,
				Long: merchant.Longitude,
			},
			Distance:  helpers.Haversine(params.Lat, params.Long, merchant.Latitude, merchant.Longitude),
			CreatedAt: merchant.CreatedAt,
		}
		Merchants = append(Merchants, mr)
	}

	sort.Slice(Merchants, func(i, j int) bool {
		return Merchants[i].Distance < Merchants[j].Distance
	})

	var metaQuery string = "SELECT COUNT(*) FROM merchants"
	metaRows, err := r.db.Query(context.Background(), metaQuery)
	if err != nil {
		fmt.Println(err.Error())
		return []entities.MerchantResponse{}, entities.MerchantMetaResponse{}, err
	}
	defer metaRows.Close()
	var meta entities.MerchantMetaResponse
	if metaRows.Next() {
		if err := metaRows.Scan(&meta.Total); err != nil {
			fmt.Println(err.Error())
			return nil, entities.MerchantMetaResponse{}, err
		}
	}
	metaRows.Scan(&meta.Total)
	meta.Limit = params.Limit
	meta.Offset = params.Offset

	return Merchants, meta, nil
}
