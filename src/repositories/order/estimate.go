package orderRepository

// func (repository *orderRepository) Create(estimateRequest entities.EstimateRequest, estDeliveryTime int, totalDistance, totalPrice, totalDeliveryFree float64) (string, error) {
// query := `INSERT INTO orders (id, user_id, est_delivery_time, total_distance, total_price, total_deliveryFee, status) values ($1, $2, $3, $4, $5, $6, $7) RETURNING id`

// tx, err := r.db.BeginTx(context.Background(), pgx.TxOptions{})
// if err != nil {
// 	return "", fmt.Errorf("failed to begin transaction: %w", err)
// }
// defer tx.Rollback(context.Background())

// _, errInsert := tx.Exec(context.Background(), query, estimateRequest)
// }
