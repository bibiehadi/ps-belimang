package entities

type ImageUploadResponse struct {
	URL string `json:"imageUrl" validate:"required"`
}
