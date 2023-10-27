package products

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Router(r *gin.RouterGroup) {

	r.POST("/create", Auth(), CreateProduct)
	r.GET("/list/all", Auth(), GetAllProducts)
	r.GET("/list/:id", Auth(), GetProductByID)
	r.DELETE("/delete/:id", Auth(), DeleteProduct)
}

func Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("Authorization")
		fmt.Println(token)
		if token == "Bearer" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Token not found",
			})
			return
		}

		if token == "lucas" {
			ctx.Next()
			return
		}
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "Token not found",
		})
	}
}
