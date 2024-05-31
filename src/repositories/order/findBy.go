package orderRepository

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func (r *orderRepository) FindById(id string) (string, error) {
	var orderId string
	var query string = "SELECT id FROM orders WHERE id = $1 LIMIT 1"
	err := r.db.QueryRow(context.Background(), query, id).Scan(&orderId)
	if err != nil {
		if err == pgx.ErrNoRows {
			return "", pgx.ErrNoRows
		}
	}
	return orderId, err
}
