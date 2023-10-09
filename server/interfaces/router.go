package interfaces

import (
	"github.com/Lucas-Linhar3s/TheBarber/interfaces/account"
	"github.com/gin-gonic/gin"
)

func Router(r *gin.RouterGroup) {
	// Account
	account.Router(r.Group("/account"))
}
