package account

import (
	"os"
	"time"

	"github.com/Lucas-Linhar3s/TheBarber/config"
	"github.com/Lucas-Linhar3s/TheBarber/database"
	"github.com/Lucas-Linhar3s/TheBarber/domain/account"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// CreatedAccount is a function that creates an account
func CreatedAccount(ctx *gin.Context, req *Req) (createdId string, err error) {
	API_URL := os.Getenv("API_URL")
	API_KEY := os.Getenv("API_KEY")

	client, err := database.NewClient(API_URL, API_KEY, nil, "public")
	if err != nil {
		return "", err
	}

	var (
		repo = account.GetService(account.GetRepository(client))
	)

	hash, err := config.GenerateHash(req.Password)
	if err != nil {
		return "", err
	}

	data := account.Account{
		ID:          uuid.New(),
		Name:        req.Name,
		LastName:    req.LastName,
		Phone:       req.Phone,
		DateOfBirth: req.DateOfBirth,
		Photo:       req.Photo,
		IsBarber:    req.IsBarber,
		Email:       req.Email,
		Password:    hash,
		CreatedAt:   time.Now().Local(),
	}

	createdId, err = repo.CreateAccount(&data)
	if err != nil {
		return "", err
	}

	return
}
