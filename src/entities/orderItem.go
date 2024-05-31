package entities

import "time"

type OrderItem struct {
	ID        int       `json:"id"`
	OrderId   Order     `json:"order_id"`
	ItemId    int       `json:"itemId"`
	ItemName  string    `lson:"itemName"`
	Quantity  int       `json:"quantity"`
	Price     int       `json:"price"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
