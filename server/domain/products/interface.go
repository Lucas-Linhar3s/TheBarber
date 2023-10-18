package products

import (
	"github.com/Lucas-Linhar3s/TheBarber/infrastructure/products"
	"github.com/google/uuid"
)

// IProduct Interface de Product
type IProduct interface {
	GetAllProducts() (products []Product, err error)
	GetProductByID(id uuid.UUID) (product Product, err error)
	CreateProduct(product Product) (productID *uuid.UUID, err error)
	DeleteProduct(id uuid.UUID) error
	UpdateProduct(product Product) error
	ConvertForRepo(model Product) (products.Product, error)
	ConvertForDomain(product products.Product) (Product, error)
}
