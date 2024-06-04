package entities

import (
	"time"
)

type Merchant struct {
	ID               string  `json:"id"`
	Name             string  `json:"name"`
	MerchantCategory string  `json:"merchantCategory"`
	ImageURL         string  `json:"imageUrl"`
	Latitude         float64 `json:"latitude"`
	Longitude        float64 `json:"longitude"`
	Range            float64
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
	Lat              float64
	Long             float64
	CreatedAt        string
	Limit            int
	Offset           int
}

type MerchantPostResponse struct {
	MerchantId string `json:"merchantId"`
}

type MerchantResponse struct {
	MerchantId       string    `json:"merchantId"`
	Name             string    `json:"name"`
	MerchantCategory string    `json:"merchantCategory"`
	ImageURL         string    `json:"imageUrl"`
	Location         Location  `json:"location"`
	Distance         float64   `json:"distance"`
	CreatedAt        time.Time `json:"createdAt"`
}

type MerchantMetaResponse struct {
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
	Total  int `json:"total"`
}
