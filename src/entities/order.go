package entities

type EstimateRequest struct {
	UserLocation Location `json:"userLocation" validate:"required"`
	Orders       []Order  `json:"orders" validate:"required"`
}

type Order struct {
	MerchantId      string      `json:"merchantId" validate:"required"`
	IsStartingPoint *bool       `json:"isStartingPoint" validate:"required"`
	Items           []OrderItem `json:"items" validate:"required"`
}

type OrderItem struct {
	ItemId   string `json:"itemId" validate:"required"`
	Quantity int    `json:"quantity" validate:"required"`
}

type EstimateResponse struct {
	TotalPrice           float64 `json:"totalPrice" validate:"required"`
	EstimateDeliveryTime float64 `json:"estimatedDeliveryTimeInMinutes" validate:"required"`
	EstimateId           string  `json:"calculatedEstimateId" validate:"required"`
}

// type Order struct {
// 	ID         int       `json:"id"`
// 	UserId     string    `json:"user_id"`
// 	Status     bool      `json:"status"`
// 	TotalPrice int       `json:"total_price"`
// 	MerchantId Merchant  `json:"merchant_id"`
// 	CreatedAt  time.Time `json:"createdAt"`
// 	UpdatedAt  time.Time `json:"updatedAt"`
// }

type OrderRequest struct {
	OrderId string `json:"calculatedEstimateId" validate:"required"`
}

type OrderResponse struct {
	OrderId string `json:"orderId"`
}

type OrderQueryParams struct {
	MerchantID       string `json:"merchantId"`
	Name             string `json:"name"`
	MerchantCategory string `json:"merchantCategory"`
	CreatedAt        string
	Limit            int
	Offset           int
}

type GetOrderResponse struct {
	OrderId string     `json:"orderId"`
	Orders  []GetOrder `json:"orders"`
}

type GetOrder struct {
	Merchant MerchantResponse    `json:"merchant"`
	Items    []OrderItemResponse `json:"items"`
}
