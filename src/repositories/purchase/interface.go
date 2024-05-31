package purchase

import (
	"belimang/src/entities"

	"github.com/jackc/pgx/v5/pgxpool"
)

type PurchaseRepository interface {
	Create(order entities.Order, orderItem entities.OrderItem) (entities.Order, entities.OrderItem, error)
	// FindAll(params entities.MerchantQueryParams) ([]entities.MerchantResponse, entities.MerchantMetaResponse, error)
	// UsernameIsExist(username string) bool
	// FindByUsername(username string) (entities.Merchant, error)
}

type purchaseRepository struct {
	db *pgxpool.Pool
}

func New(db *pgxpool.Pool) *purchaseRepository {
	return &purchaseRepository{db}
}
