package account

import (
	"github.com/Lucas-Linhar3s/TheBarber/application/account"
	"github.com/Lucas-Linhar3s/TheBarber/config"
	"github.com/gin-gonic/gin"
)

// CreatedAccount godoc
// @Summary Create a new account
// @Description Create a new account
// @Tags CreateAccount
// @Accept json
// @Produce json
// @Param account body account.Req true "Params to create an account"
// @Success 201 "Created"
// @Router  /account/create [post]
func CreatedAccount(ctx *gin.Context) {
	var req *account.Req

	if err := ctx.ShouldBindJSON(&req); err != nil {
		config.ResponseWithError(ctx, 400, err)
		return
	}

	createdId, err := account.CreatedAccount(ctx, req)
	if err != nil {
		config.ResponseWithError(ctx, 400, err)
		return
	}

	config.ResponseWithMessageAndData(ctx, 201, "Success", createdId)
}

func Login(ctx *gin.Context) {
	config.ResponseWithMessage(ctx, 200, "Login")
}
