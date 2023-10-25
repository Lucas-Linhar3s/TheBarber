package products

import (
	"github.com/Lucas-Linhar3s/TheBarber/database"
	"github.com/Lucas-Linhar3s/TheBarber/infrastructure/products"
	"github.com/Lucas-Linhar3s/TheBarber/infrastructure/products/supabase"
	"github.com/google/uuid"
)

type repository struct {
	repo *supabase.PGRepository
}

func newRepository(client *database.Client) *repository {
	return &repository{
		repo: &supabase.PGRepository{DB: client},
	}
}

// ConvertForRepo converts a model Product to a products.Product.
func (r *repository) ConvertForRepo(model Product) (products.Product, error) {
	return products.Product{
		ID:        model.ID,
		AccountID: model.AccountID,
		Name:      model.Name,
		Price:     model.Price,
		CreatedAt: model.CreatedAt,
		UpdatedAt: model.UpdatedAt,
	}, nil
}

// ConvertForDomain converts a product from the repository format to the domain format.
func (r *repository) ConvertForDomain(product products.Product) (Product, error) {
	return Product{
		ID:        product.ID,
		AccountID: product.AccountID,
		Name:      product.Name,
		Price:     product.Price,
		CreatedAt: product.CreatedAt,
		UpdatedAt: product.UpdatedAt,
	}, nil
}

// GetAllProducts retrieves all products from the repository.
func (r *repository) GetAllProducts() (products []Product, err error) {
	result, err := r.repo.GetAllProducts()
	if err != nil {
		return products, err
	}

	for _, p := range result {
		product, err := r.ConvertForDomain(p)
		if err != nil {
			return products, err
		}
		products = append(products, product)
	}

	return
}

// GetProductByID retrieves a product by its ID.
func (r *repository) GetProductByID(id uuid.UUID) (product Product, err error) {

	result, err := r.repo.GetProductByID(id)
	if err != nil {
		return product, err
	}

	product, err = r.ConvertForDomain(result)
	if err != nil {
		return product, err
	}
	return
}

// CreateProduct creates a new product in the repository.
func (r *repository) CreateProduct(product Product) (productID *uuid.UUID, err error) {
	data, err := r.ConvertForRepo(product)
	if err != nil {
		return productID, err
	}
	return r.repo.CreateProduct(data)
}

// DeleteProduct deletes a product from the repository.
func (r *repository) DeleteProduct(id uuid.UUID) error {
	return r.repo.DeleteProduct(id)
}

// UpdateProduct updates a product in the repository.
func (r *repository) UpdateProduct(product Product) error {
	data, err := r.ConvertForRepo(product)
	if err != nil {
		return err
	}
	return r.repo.UpdateProduct(data)
}
