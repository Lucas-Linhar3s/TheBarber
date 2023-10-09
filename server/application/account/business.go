package account

import (
	"os"

	"github.com/Lucas-Linhar3s/TheBarber/config"
	"github.com/Lucas-Linhar3s/TheBarber/database"
	"github.com/Lucas-Linhar3s/TheBarber/domain/account"
	"github.com/gin-gonic/gin"
)

func CreatedAccount(ctx *gin.Context) (result interface{}, err error) {
	API_URL := os.Getenv("API_URL")
	API_KEY := os.Getenv("API_KEY")

	client, err := database.NewClient(API_URL, API_KEY, nil)
	if err != nil {
		config.ResponseWithError(ctx, 500, err)
		return
	}

	var (
		repo = account.GetService(account.GetRepository(client))
		req  Req
	)

	data := account.Account{
		Email:    "account",
		Password: req.Password,
	}

	result, err = repo.Login(data)
	if err != nil {
		config.ResponseWithError(ctx, 500, err)
		return
	}
	return

}
