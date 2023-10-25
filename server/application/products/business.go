package products

import (
	"os"
	"time"

	"github.com/Lucas-Linhar3s/TheBarber/database"
	"github.com/Lucas-Linhar3s/TheBarber/domain/products"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// CreateProduct creates a new product in the database.
func CreateProduct(ctx *gin.Context, req ProductReq) (createdID *uuid.UUID, err error) {
	API_URL := os.Getenv("API_URL")
	API_KEY := os.Getenv("API_KEY")

	client, err := database.NewClient(API_URL, API_KEY, nil, "public")
	if err != nil {
		return nil, err
	}

	var (
		repo = products.GetService(products.GetRepository(client))
	)

	data := products.Product{
		ID:        uuid.New(),
		AccountID: req.AccountID,
		Name:      req.Name,
		Price:     req.Price,
		CreatedAt: time.Now().Local(),
		CreatedAt: time.Now().Local(),
	}

	createdID, err = repo.CreateProduct(data)
	if err != nil {
		return nil, err
	}

	return
}

func GetAllProducts(ctx *gin.Context) (product *ProductResPag, err error) {
	API_URL := os.Getenv("API_URL")
	API_KEY := os.Getenv("API_KEY")

	client, err := database.NewClient(API_URL, API_KEY, nil, "public")
	if err != nil {
		return nil, err
	}

	var (
		repo = products.GetService(products.GetRepository(client))
	)

	result, err := repo.GetAllProducts()
	if err != nil {
		return nil, err
	}
	var products []ProductRes

	for _, p := range result {
		products = append(products, ProductRes{
			ID:        p.ID,
			AccountID: p.AccountID,
			Name:      p.Name,
			Price:     p.Price,
			CreatedAt: p.CreatedAt,
			UpdatedAt: p.UpdatedAt,
		})
	}

	product = &ProductResPag{
		Count:      len(result),
		ProductRes: products,
	}

	return

}

func GetProductByID(ctx *gin.Context, id uuid.UUID) (product *ProductRes, err error) {
	API_URL := os.Getenv("API_URL")
	API_KEY := os.Getenv("API_KEY")

	client, err := database.NewClient(API_URL, API_KEY, nil, "public")
	if err != nil {
		return nil, err
	}

	var (
		repo = products.GetService(products.GetRepository(client))
	)

	result, err := repo.GetProductByID(id)
	if err != nil {
		return nil, err
	}

	product = &ProductRes{
		ID:        result.ID,
		AccountID: result.AccountID,
		Name:      result.Name,
		Price:     result.Price,
		CreatedAt: result.CreatedAt,
		UpdatedAt: result.UpdatedAt,
	}

	return
}
