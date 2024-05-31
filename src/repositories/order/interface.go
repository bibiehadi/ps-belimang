package orderRepository

import (
	"github.com/jackc/pgx/v5/pgxpool"
)

type OrderResorderRepository interface {
}

type orderRepository struct {
	db *pgxpool.Pool
}

func New(db *pgxpool.Pool) *orderRepository {
	return &orderRepository{db}
}
