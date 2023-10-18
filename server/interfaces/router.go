package interfaces

import (
	"github.com/Lucas-Linhar3s/TheBarber/interfaces/account"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Router(r *gin.RouterGroup) {
	// Account
	account.Router(r.Group("/account"))
	// Swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
