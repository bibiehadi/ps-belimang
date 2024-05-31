package orderRepository

import (
	"belimang/src/entities"

	"github.com/jackc/pgx/v5/pgxpool"
)

type OrderResorderRepository interface {
	Create(estimateRequest entities.EstimateRequest, estDeliveryTime, totalDistance, totalPrice, totalDeliveryFree float64, userId string) (string, error)
	FindById(id string) (string, error)
	Order(id string) (string, error)
}

type orderRepository struct {
	db *pgxpool.Pool
}

func New(db *pgxpool.Pool) *orderRepository {
	return &orderRepository{db}
}
