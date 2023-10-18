package products

import "github.com/gin-gonic/gin"

func Router(r *gin.RouterGroup) {
	r.POST("/create", CreateProduct)
	r.GET("/list/all", GetAllProducts)
}
