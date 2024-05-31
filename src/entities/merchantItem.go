package entities

import (
	"time"
)

// MerchantItem represents a product sold by a merchant
type MerchantItem struct {
	ID              string          `json:"id"`
	Name            string          `json:"name"`
	ProductCategory ProductCategory `json:"productCategory"`
	Price           int             `json:"price"`
	ImageURL        string          `json:"imageUrl"`
	MerchantID      uint            `json:"merchantId"`
	CreatedAt       time.Time       `json:"createdAt"`
	UpdatedAt       time.Time       `json:"updatedAt"`
	Merchant        *Merchant       `json:"merchant"` // Optional pointer to Merchant struct
}

type MerchantItemRequest struct {
	Name            string `json:"name" validate:"required,min=2,max=30"`
	ProductCategory string `json:"productCategory" validate:"required,oneof=Beverage Food Snack Condiments Additions"`
	Price           int    `json:"price" validate:"required,numeric,min=1"`
	ImageURL        string `json:"imageUrl" validate:"required,url"`
	MerchantID      string
}

type MerchantItemQueryParams struct {
	ItemId          string `json:"itemId"`
	Name            string `json:"name"`
	ProductCategory string `json:"productCategory"`
	MerchantId      string `json:"merchantId"`
	Limit           int    `json:"limit"`
	Offset          int    `json:"offset"`
	CreatedAt       string `json:"createdAt"`
}

type MerchantItemResponse struct {
	ItemId          string    `json:"itemId"`
	Name            string    `json:"name"`
	ProductCategory string    `json:"productCategory"`
	Price           int       `json:"price"`
	ImageUrl        string    `json:"imageUrl"`
	CreatedAt       time.Time `json:"createdAt"`
}
type MerchantItemMetaResponse struct {
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
	Total  int `json:"total"`
}

type CreateMerchantItemResponse struct {
	ItemId string `json:"itemId"`
}
