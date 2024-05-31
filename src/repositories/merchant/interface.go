package merchantRepository

import (
	"belimang/src/entities"

	"github.com/jackc/pgx/v5/pgxpool"
)

type MerchantRepository interface {
	Create(merchant entities.Merchant) (entities.Merchant, error)
	FindAll(params entities.MerchantQueryParams) ([]entities.MerchantResponse, entities.MerchantMetaResponse, error)
	FindNearby(params entities.MerchantQueryParams) ([]entities.MerchantResponse, entities.MerchantMetaResponse, error)
	FindById(id string) (entities.Merchant, error)
}

type merchantRepository struct {
	db *pgxpool.Pool
}

func New(db *pgxpool.Pool) *merchantRepository {
	return &merchantRepository{db}
}
