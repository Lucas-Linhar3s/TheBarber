package main

import (
	"log"

	"github.com/Lucas-Linhar3s/TheBarber/docs"
	"github.com/Lucas-Linhar3s/TheBarber/interfaces"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title the barber API
// @version 1.0
// @description This is the barber.

// @contact.name Linhares

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

// @host localhost:8080
// @BasePath /v1
func main() {
	docs.SwaggerInfo.BasePath = "/v1"
	// Carrega as variáveis de ambiente do arquivo .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Erro ao carregar o arquivo .env")
	}

	r := gin.Default()

	api := r.Group("api")
	api.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	interfaces.Router(r.Group("/v1"))
	r.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found " + " : " + c.Request.URL.String()})
	})

	if err := r.Run(); err != nil {
		log.Fatal(err)
	} // listen and serve on 0.0.0.0:8080

	if err := r.Run(":8888"); err != nil {
		log.Fatal("Erro ao iniciar o servidor")
	}
}
