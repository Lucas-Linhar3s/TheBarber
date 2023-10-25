package supabase

import (
	"encoding/json"

	"github.com/Lucas-Linhar3s/TheBarber/database"
	"github.com/Lucas-Linhar3s/TheBarber/infrastructure/products"
	"github.com/google/uuid"
)

// PGRepository struct
type PGRepository struct {
	DB *database.Client
}

// GetAllProducts returns all products from the PGRepository.
func (r *PGRepository) GetAllProducts() (products []products.Product, err error) {

	query, count, err := r.DB.From("product").Select("*", "exact", false).Execute()
	if err != nil {
		return products, err
	}

	if count == 0 {
		return products, err
	}

	if err = json.Unmarshal(query, &products); err != nil {
		return products, err
	}

	return

}

// GetProductByID retrieves a product from the PGRepository based on its ID.
func (r *PGRepository) GetProductByID(id uuid.UUID) (product products.Product, err error) {
	query, count, err := r.DB.From("product").Select("*", "exact", false).Eq("id", id.String()).Execute()
	if err != nil {
		return product, err
	}

	if count == 0 {
		return product, err
	}

	var products []products.Product

	if err = json.Unmarshal(query, &products); err != nil {
		return product, err
	}

	for _, v := range products {
		product = v
	}

	return
}

// CreateProduct creates a new product in the PGRepository.
func (r *PGRepository) CreateProduct(product products.Product) (productID *uuid.UUID, err error) {
	_, count, err := r.DB.From("product").Insert(product, true, "", "*", "exact").Execute()
	if err != nil {
		return productID, err
	}

	if count == 0 {
		return productID, err
	}

	productID = &product.ID

	return
}

// DeleteProduct deletes a product from the PGRepository.
func (r *PGRepository) DeleteProduct(id uuid.UUID) error {
	return nil
}

// UpdateProduct updates the specified product in the PGRepository.
func (r *PGRepository) UpdateProduct(product products.Product) error {
	return nil
}
