package entities

type SuccessResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type SuccessGetAllResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Meta    interface{} `json:"meta"`
}

type ErrorResponse struct {
	Status  bool        `json:"status"`
	Message interface{} `json:"message"`
}
