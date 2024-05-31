package entities

import "time"

type Order struct {
	ID         int       `json:"id"`
	UserId     string    `json:"user_id"`
	Status     bool      `json:"status"`
	TotalPrice int       `json:"total_price"`
	MerchantId Merchant  `json:"merchant_id"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
}

type OrderResponse struct {
	OrderId string `json:"orderId"`
}

type UserLocation struct {
	Lat  float64 `json:"lat"`
	Long float64 `json:"long"`
}
