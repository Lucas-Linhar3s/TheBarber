package products

import (
	"time"

	"github.com/google/uuid"
)

// ProductReq struct
type ProductReq struct {
	AccountID uuid.UUID `json:"account_id" binding:"required"`
	Name      string    `json:"name" binding:"required"`
	Price     float64   `json:"price" binding:"required"`
}

type ProductRes struct {
	ID        uuid.UUID `json:"id,omitempty"`
	AccountID uuid.UUID `json:"account_id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Price     float64   `json:"price,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

type ProductResPag struct {
	ProductRes []ProductRes `json:"products,omitempty"`
	Count      int          `json:"count,omitempty"`
}

// SuccessRes struct para retornar uma	respost personalizada personalizada
type SuccessRes struct {
	Data    string `json:"data"`
	Message string `json:"message"`
}
