package main

import (
	"log"

	"github.com/Lucas-Linhar3s/TheBarber/interfaces"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Carrega as variáveis de ambiente do arquivo .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Erro ao carregar o arquivo .env")
	}

	r := gin.New()
	interfaces.Router(r.Group("/v1"))

	if err := r.Run(":8888"); err != nil {
		log.Fatal("Erro ao iniciar o servidor")
	}
}
