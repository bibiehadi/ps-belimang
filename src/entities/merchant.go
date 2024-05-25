package entities

import (
	"time"
)

// Merchant represents a vendor selling items
type Merchant struct {
	ID        string         `json:"id"`
	Name      string         `json:"name"`
	Category  string         `json:"merchant_category"` // Can be an enum type if applicable
	ImageURL  string         `json:"image_url"`
	Latitude  float64        `json:"latitude"`
	Longitude float64        `json:"longitude"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	Items     []MerchantItem `json:"items,omitempty"` // Optional slice of MerchantItem (one-to-many relationship)
}
