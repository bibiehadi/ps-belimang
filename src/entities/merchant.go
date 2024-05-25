package entities

import (
	"time"
)

type Merchant struct {
	ID               int       `json:"id"`
	Name             string    `json:"name"`
	MerchantCategory string    `json:"merchantCategory"`
	ImageURL         string    `json:"imageUrl"`
	Latitude         float64   `json:"latitude"`
	Longitude        float64   `json:"longitude"`
	CreatedAt        time.Time `json:"createdAt"`
	UpdatedAt        time.Time `json:"updatedAt"`
}

type Location struct {
	Lat  float64 `json:"lat"`
	Long float64 `json:"long"`
}

type MerchantRequest struct {
	Name             string   `json:"name"`
	MerchantCategory string   `json:"merchantCategory"`
	ImageURL         string   `json:"imageUrl"`
	Location         Location `json:"location"`
}

type MerchantQueryParams struct {
	MerchantID       string `json:"merchantId"`
	Name             string `json:"name"`
	MerchantCategory string `json:"merchantCategory"`
	CreatedAt        string
	Limit            int
	Offset           int
}

type MerchantResponse struct {
	MerchantId       int       `json:"merchantId"`
	Name             string    `json:"name"`
	MerchantCategory string    `json:"merchantCategory"`
	ImageURL         string    `json:"imageUrl"`
	Location         Location  `json:"location"`
	CreatedAt        time.Time `json:"createdAt"`
}

type MerchantMetaResponse struct {
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
	Total  int `json:"total"`
}
