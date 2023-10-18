package products

import (
	"os"

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
	}

	createdID, err = repo.CreateProduct(data)
	if err != nil {
		return nil, err
	}

	return
}
