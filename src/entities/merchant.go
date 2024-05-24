package entities

import (
	"errors"
	"regexp"
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
	Name             string    `json:"name"`
	MerchantCategory string    `json:"merchant_category"`
	ImageURL         string    `json:"image_url"`
	Location         Location  `json:"location"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}

func validateMerchant(merchant *MerchantRequest) error {
	// Validate Name
	if len(merchant.Name) < 2 || len(merchant.Name) > 30 {
		return errors.New("name must be between 2 and 30 characters")
	}

	// Validate MerchantCategory
	validCategories := []string{"SmallRestaurant", "MediumRestaurant", "LargeRestaurant", "MerchandiseRestaurant", "BoothKiosk", "ConvenienceStore"}
	isValidCategory := false
	for _, category := range validCategories {
		if merchant.MerchantCategory == category {
			isValidCategory = true
			break
		}
	}
	if !isValidCategory {
		return errors.New("merchant_category must be one of the valid categories")
	}

	// Validate ImageURL
	imageURLPattern := `^(http(s?):)([/|.|\w|\s|-])*\.(?:jpg|gif|png)$`
	matched, err := regexp.MatchString(imageURLPattern, merchant.ImageURL)
	if err != nil || !matched {
		return errors.New("image_url must be a valid image URL")
	}

	// Validate Location
	if merchant.Location.Lat == 0 || merchant.Location.Long == 0 {
		return errors.New("latitude and longitude cannot be zero")
	}

	return nil
}
