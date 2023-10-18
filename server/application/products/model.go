package products

import "github.com/google/uuid"

// ProductReq struct
type ProductReq struct {
	AccountID uuid.UUID `json:"account_id", binding:"required"`
	Name      string    `json:"name", binding:"required"`
	Price     float64   `json:"price", binding:"required"`
}

// SuccessRes struct para retornar uma	respost personalizada personalizada
type SuccessRes struct {
	Data    string `json:"data"`
	Message string `json:"message"`
}
