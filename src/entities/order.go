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

type OrderRequest struct {
	OrderId string `json:"calculatedEstimateId" validate:"required"`
}

type OrderResponse struct {
	OrderId string `json:"orderId"`
}
