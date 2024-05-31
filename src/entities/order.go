package entities

type EstimateRequest struct {
	UserLocation Location       `json:"userLocation" validate:"required"`
	Orders       []OrderRequest `json:"orders" validate:"required"`
}

type OrderRequest struct {
	MerchantId      string             `json:"merchantId" validate:"required"`
	IsStartingPoint *bool              `json:"isStartingPoint" validate:"required"`
	Items           []OrderItemRequest `json:"items" validate:"required"`
}

type OrderItemRequest struct {
	ItemId   string `json:"itemId" validate:"required"`
	Quantity int    `json:"quantity" validate:"required"`
}

type EstimateResponse struct {
	TotalPrice           float64 `json:"totalPrice" validate:"required"`
	EstimateDeliveryTime float64 `json:"estimatedDeliveryTimeInMinutes" validate:"required"`
	EstimateId           string  `json:"calculatedEstimateId" validate:"required"`
}
