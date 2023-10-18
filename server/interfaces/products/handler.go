package products

import (
	"github.com/Lucas-Linhar3s/TheBarber/application/products"
	"github.com/Lucas-Linhar3s/TheBarber/config"
	"github.com/gin-gonic/gin"
)

// CreateProduct godoc
// @Summary Create a new product
// @Description Create a new product
// @Tags Product
// @Accept json
// @Produce json
// @Param product body products.ProductReq true "Create new Products"
// @Success 201 {object} products.SuccessRes "Created"
// @Router /product/create [post]
func CreateProduct(ctx *gin.Context) {
	var req *products.ProductReq

	if err := ctx.ShouldBindJSON(&req); err != nil {
		config.ResponseWithError(ctx, 400, err)
		return
	}

	createdId, err := products.CreateProduct(ctx, *req)
	if err != nil {
		config.ResponseWithError(ctx, 400, err)
		return
	}

	config.ResponseWithMessageAndData(ctx, 201, "Success on created product", createdId)
}

// GetAllProducts godoc
// @Summary Get all products
// @Description Get all products
// @Tags Product
// @Accept json
// @Produce json
// @Success 200 {object} products.ProductResPag "ListAllProducts"
// @Router /product/list/all [get]
func GetAllProducts(ctx *gin.Context) {
	products, err := products.GetAllProducts(ctx)
	if err != nil {
		config.ResponseWithError(ctx, 400, err)
		return
	}

	config.Response(ctx, 200, products)
}
