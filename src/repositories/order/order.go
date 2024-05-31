package orderRepository

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func (r *orderRepository) Order(id string) (string, error) {
	var query string = "UPDATE orders SET status = true WHERE id = $1"
	_, err := r.db.Exec(context.Background(), query, id)
	if err != nil {
		fmt.Println(err)
		if err == pgx.ErrNoRows {
			return "", pgx.ErrNoRows
		}
	}
	return id, err
}
