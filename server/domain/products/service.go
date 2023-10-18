package products

import (
	"github.com/Lucas-Linhar3s/TheBarber/database"
	"github.com/google/uuid"
)

type Service struct {
	pg IProduct
}

// NewService creates a new account service
func GetService(pg IProduct) *Service {
	return &Service{
		pg: pg,
	}
}

// GetRepository Responsavel por pegar se conectar com os metodos no banco de dados
func GetRepository(product *database.Client) IProduct {
	return newRepository(product)
}

// GetAllProducts retrieves all products from the service.
func (s *Service) GetAllProducts() (products []Product, err error) {
	return s.pg.GetAllProducts()
}

// GetProductByID retrieves a product by its ID.
func (s *Service) GetProductByID(id uuid.UUID) (product Product, err error) {
	return s.pg.GetProductByID(id)
}

// CreateProduct creates a new product in the repository.
func (s *Service) CreateProduct(product Product) (productID *uuid.UUID, err error) {
	return s.pg.CreateProduct(product)
}

// DeleteProduct deletes a product from the repository.
func (s *Service) DeleteProduct(id uuid.UUID) error {
	return s.pg.DeleteProduct(id)
}

// UpdateProduct updates a product in the repository.
func (s *Service) UpdateProduct(product Product) error {
	return s.pg.UpdateProduct(product)
}
