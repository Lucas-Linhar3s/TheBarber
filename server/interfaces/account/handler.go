package account

import (
	"github.com/Lucas-Linhar3s/TheBarber/application/account"
	"github.com/Lucas-Linhar3s/TheBarber/config"
	"github.com/gin-gonic/gin"
)

func CreatedAccount(ctx *gin.Context) {
	config.ResponseWithMessage(ctx, 201, "CreatedAccount")
}

func Login(ctx *gin.Context) {
	data, err := account.CreatedAccount(ctx)
	if err != nil {
		config.ResponseWithError(ctx, 500, err)
		return
	}

	config.ResponseWithMessageAndData(ctx, 200, "Login", data)

}
