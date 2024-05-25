package entities

import (
	"time"
)

// MerchantItem represents a product sold by a merchant
type MerchantItem struct {
	ID              string          `json:"id"`
	Name            string          `json:"name"`
	ProductCategory ProductCategory `json:"product_category"`
	Price           float64         `json:"price"`
	ImageURL        string          `json:"image_url"`
	MerchantID      uint            `json:"merchant_id"`
	CreatedAt       time.Time       `json:"created_at"`
	UpdatedAt       time.Time       `json:"updated_at"`
	Merchant        *Merchant       `json:"merchant"` // Optional pointer to Merchant struct
}

type MerchantItemRequest struct {
	Name            string `json:"name" validate:"required,min=2,max=30"`
	ProductCategory string `json:"product_category" validate:"required,oneof=Beverage Food Snack Condiments Additions"`
	Price           string `json:"price" validate:"required,numeric,min=1"`
	ImageURL        string `json:"image_url" validate:"required,url"`
	MerchantID      string
}

type MerchantItemQueryParams struct {
	ItemId          string `json:"itemId"`
	Name            string `json:"name"`
	ProductCategory string `json:"productCategory" validate:"oneof=Beverage Food Snack Condiments Additions"`
	Limit           int    `json:"limit"`
	Offset          int    `json:"offset"`
	CreatedAt       string `json:"createdAt"`
}

type MerchantItemResponse struct {
	ItemId          string `json:"itemId"`
	Name            string `json:"name"`
	ProductCategory string `json:"productCategory"`
	Price           int    `json:"price"`
	ImageUrl        string `json:"imageUrl"`
	CreatedAt       string `json:"createdAt"`
}

type CreateMerchantItemResponse struct {
	ItemId string `json:"itemId"`
}
