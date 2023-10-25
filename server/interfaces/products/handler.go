package products

import (
	"github.com/Lucas-Linhar3s/TheBarber/application/products"
	"github.com/Lucas-Linhar3s/TheBarber/config"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// CreateProduct godoc
// @Summary Create a new product
// @Description Create a new product
// @Tags Product
// @Accept json
// @Produce json
// @Param product body products.ProductReq true "Create new Products"
// @Success 201 {object} products.SuccessRes "Created"
// @Security ApiKeyAuth
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
// @Security ApiKeyAuth
func GetAllProducts(ctx *gin.Context) {
	products, err := products.GetAllProducts(ctx)
	if err != nil {
		config.ResponseWithError(ctx, 400, err)
		return
	}

	config.Response(ctx, 200, products)
}

// GetProductByID godoc
// @Summary Get one products by ID
// @Description Get one products by ID
// @Tags Product
// @Accept json
// @Produce json
// @Param id path string true "Product ID"
// @Success 200 {object} products.ProductRes "ListOneProducts"
// @Router /product/list/{id} [get]
// @Security ApiKeyAuth
func GetProductByID(ctx *gin.Context) {

	id := ctx.Param("id")

	uuid, err := uuid.Parse(id)
	if err != nil {
		config.ResponseWithError(ctx, 400, err)
		return
	}

	product, err := products.GetProductByID(ctx, uuid)
	if err != nil {
		config.ResponseWithError(ctx, 400, err)
		return
	}

	config.ResponseWithMessageAndData(ctx, 200, "List one product", product)

}
