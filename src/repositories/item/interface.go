package itemRepository

import (
	"belimang/src/entities"

	"github.com/jackc/pgx/v5/pgxpool"
)

type ItemRepository interface {
	Create(item entities.MerchantItemRequest) (entities.MerchantItem, error)
	GetAll(params entities.MerchantItemQueryParams) ([]entities.MerchantItemResponse, entities.MerchantItemMetaResponse, error)
	FindById(id string) (entities.MerchantItem, error)
}

type itemRepository struct {
	db *pgxpool.Pool
}

func New(db *pgxpool.Pool) *itemRepository {
	return &itemRepository{db}
}
