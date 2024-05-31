package orderRepository

import (
	"belimang/src/entities"
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func (repository *orderRepository) Create(estimateRequest entities.EstimateRequest, estDeliveryTime, totalDistance, totalPrice, totalDeliveryFree float64, userId string) (string, error) {
	query := `INSERT INTO orders (user_id, est_delivery_time, total_distance, total_purchase, total_deliveryFee, total_order, status) values ($1, $2, $3, $4, $5, $6, $7) RETURNING id`

	tx, err := repository.db.BeginTx(context.Background(), pgx.TxOptions{})
	if err != nil {
		return "", fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback(context.Background())
	var orderId string
	if err = tx.QueryRow(context.Background(), query, userId, estDeliveryTime, totalDistance, totalPrice, totalDeliveryFree, totalPrice+totalDeliveryFree, false).Scan(&orderId); err != nil {
		return "", err
	}

	for _, merchant := range estimateRequest.Orders {
		var merchantId string
		if err = tx.QueryRow(context.Background(), "INSERT INTO ordered_merchants (order_id, merchant_id) VALUES ($1, $2) RETURNING id", orderId, merchant.MerchantId).Scan(&merchantId); err != nil {
			return "", err
		}

		for _, item := range merchant.Items {
			if _, err = tx.Exec(context.Background(), "INSERT INTO ordered_items (ordered_merchants_id, item_id, quantity) VALUES ($1, $2, $3)", merchantId, item.ItemId, item.Quantity); err != nil {
				return "", err
			}
		}
	}

	if errCommit := tx.Commit(context.Background()); errCommit != nil {
		return "", fmt.Errorf("failed to commit transaction: %w", errCommit)
	}

	return orderId, nil
}
