package orderRepository

import (
	"belimang/src/entities"
	"context"
	"fmt"
	"strconv"
	"strings"
)

func (r *orderRepository) FindAll(params entities.OrderQueryParams) ([]entities.GetOrderResponse, error) {
	query := `WITH limited_orders AS (
		SELECT id 
		FROM orders
		LIMIT $1
		OFFSET $2
	)
	SELECT orders.id AS orderId,
	m.id AS merchantId,
	m.name AS merchantName,
	m.merchant_category AS merchantCategory,
	m.image_url AS merchantImageUrl,
	m.latitude AS latitude,
	m.longitude AS longitude,
	m.created_at AS merchantCreatedAt,
	oi.item_id  AS itemId,
	mi.name AS itemName,
	mi.product_category AS itemCategory,
	mi.price AS price,
	oi.quantity AS quantity,
	mi.image_url AS itemImageUrl,
	mi.created_at AS itemCreatedAt
	FROM limited_orders
	JOIN orders ON limited_orders.id = orders.id 
	JOIN order_items oi  ON orders.id = oi.order_id 
	JOIN merchants m ON oi.merchant_id = m.id 
	JOIN merchant_items mi ON oi.item_id = mi.id 
	WHERE orders.status = true AND 1=1`

	if params.MerchantID != "" {
		query += fmt.Sprintf(" AND oi.merchant_id = '%s'", params.MerchantID)
	}

	if params.Name != "" {
		query += fmt.Sprintf(" AND (LOWER(m.name) LIKE '%%%s%%' OR LOWER(mi.name) LIKE '%%%s%%')", strings.ToLower(params.Name), strings.ToLower(params.Name))
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
	} else {
		query += " ORDER BY orderId, merchantId ASC"
	}

	rows, err := r.db.Query(context.Background(), query, strconv.Itoa(params.Limit), strconv.Itoa(params.Offset))
	fmt.Println(query)
	if err != nil {
		fmt.Println(err.Error())
		return []entities.GetOrderResponse{}, err
	}

	defer rows.Close()
	var Orders []entities.GetOrderResponse
	orderMap := make(map[string]*entities.GetOrderResponse)
	for rows.Next() {
		var orderId string
		var merchant entities.Merchant
		var item entities.OrderItemResponse
		err := rows.Scan(&orderId, &merchant.ID, &merchant.Name, &merchant.MerchantCategory, &merchant.ImageURL, &merchant.Latitude, &merchant.Longitude, &merchant.CreatedAt, &item.ItemId, &item.Name, &item.ProductCategory, &item.Price, &item.Quantity, &item.ImageUrl, &item.CreatedAt)
		if err != nil {
			return []entities.GetOrderResponse{}, err
		}

		orderResponse, orderExists := orderMap[orderId]
		if !orderExists {
			orderResponse = &entities.GetOrderResponse{
				OrderId: orderId,
				Orders: []entities.GetOrder{
					{
						Merchant: entities.MerchantResponse{
							MerchantId:       merchant.ID,
							Name:             merchant.Name,
							MerchantCategory: merchant.MerchantCategory,
							Location: entities.Location{
								Lat:  merchant.Latitude,
								Long: merchant.Longitude,
							},
							ImageURL:  merchant.ImageURL,
							CreatedAt: merchant.CreatedAt,
						},
						Items: []entities.OrderItemResponse{},
					},
				},
			}
			Orders = append(Orders, *orderResponse)
			orderMap[orderId] = &Orders[len(Orders)-1]
		}

		var tempMerchant *entities.GetOrder
		for i := range orderResponse.Orders {
			if orderResponse.Orders[i].Merchant.MerchantId == merchant.ID {
				tempMerchant = &orderResponse.Orders[i]
				break
			}
		}

		if tempMerchant == nil {
			newMerchant := entities.GetOrder{
				Merchant: entities.MerchantResponse{
					MerchantId:       merchant.ID,
					Name:             merchant.Name,
					MerchantCategory: merchant.MerchantCategory,
					ImageURL:         merchant.ImageURL,
					Location: entities.Location{
						Lat:  merchant.Latitude,
						Long: merchant.Longitude,
					},
					CreatedAt: merchant.CreatedAt,
				},
				Items: []entities.OrderItemResponse{{
					ItemId:          item.ItemId,
					Name:            item.Name,
					ProductCategory: item.ProductCategory,
					Price:           item.Price,
					Quantity:        item.Quantity,
					ImageUrl:        item.ImageUrl,
					CreatedAt:       item.CreatedAt,
				}},
			}
			orderMap[orderId].Orders = append(orderResponse.Orders, newMerchant)
			tempMerchant = &orderResponse.Orders[len(orderResponse.Orders)-1]

		}
		tempMerchant.Items = append(tempMerchant.Items, entities.OrderItemResponse{
			ItemId:          item.ItemId,
			Name:            item.Name,
			ProductCategory: item.ProductCategory,
			Price:           item.Price,
			Quantity:        item.Quantity,
			ImageUrl:        item.ImageUrl,
			CreatedAt:       item.CreatedAt,
		},
		)

	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("row iteration error: %v", err)
	}

	return Orders, nil
}
