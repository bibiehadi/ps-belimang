package itemRepository

import (
	"belimang/src/entities"
	"context"
	"fmt"
	"strconv"
	"strings"
)

func (r *itemRepository) GetAll(params entities.MerchantItemQueryParams) ([]entities.MerchantItemResponse, entities.MerchantItemMetaResponse, error) {
	var query string = "SELECT id, name, product_category, price, image_url, created_at FROM merchant_items"
	conditions := "merchant_id = '" + params.MerchantId + "' AND"

	// Filter by ID
	if params.ItemId != "" {
		conditions += " id = '" + params.ItemId + "' AND"
	}
	if params.Name != "" {
		conditions += " name = '" + params.Name + "' AND"
	}
	if params.ProductCategory != "" {
		conditions += " product_category = '" + params.ProductCategory + "' AND"
	}
	if conditions != "" {
		conditions = " WHERE " + strings.TrimSuffix(conditions, " AND")
	}
	query += conditions
	var orderBy []string
	if params.CreatedAt != "" {
		orderBy = append(orderBy, "created_at "+params.CreatedAt)
	}
	if len(orderBy) > 0 {
		query += " ORDER BY " + strings.Join(orderBy, ", ")
	} else {
		query += " ORDER BY created_at DESC"
	}

	query += " LIMIT " + strconv.Itoa(params.Limit) + " OFFSET " + strconv.Itoa(params.Offset)
	rows, err := r.db.Query(context.Background(), query)

	fmt.Println(query)

	if err != nil {
		fmt.Println(err.Error())
		return []entities.MerchantItemResponse{}, entities.MerchantItemMetaResponse{}, err
	}
	defer rows.Close()
	var Items []entities.MerchantItemResponse
	for rows.Next() {
		var item entities.MerchantItemResponse
		err := rows.Scan(&item.ItemId, &item.Name, &item.ProductCategory, &item.Price, &item.ImageUrl, &item.CreatedAt)
		if err != nil {
			return []entities.MerchantItemResponse{}, entities.MerchantItemMetaResponse{}, err
		}
		Items = append(Items, item)
	}
	var metaQuery string = "SELECT COUNT(*) FROM merchant_items"
	metaQuery += conditions
	metaRows, err := r.db.Query(context.Background(), query)
	if err != nil {
		fmt.Println(err.Error())
		return []entities.MerchantItemResponse{}, entities.MerchantItemMetaResponse{}, err
	}
	defer metaRows.Close()
	var meta entities.MerchantItemMetaResponse
	metaRows.Scan(&meta.Total)
	meta.Limit = params.Limit
	meta.Offset = params.Offset
	fmt.Println(Items)
	return Items, meta, nil

}
