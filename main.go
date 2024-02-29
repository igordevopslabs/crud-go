package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/igordevopslabs/crud-go/src/configuration/logger"
	"github.com/igordevopslabs/crud-go/src/controller/routes"
	"github.com/joho/godotenv"
)

func main() {
	logger.Info("app iniciando")
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error ao carregar o arquivo .env")
	}
	fmt.Println(os.Getenv("TEST"))

	//aqui poderiamos usar o gin.New() tabm, mas o gin.Default() tras mais informações de debug
	router := gin.Default()

	routes.InitRoutes(&router.RouterGroup)

	if err := router.Run(":5000"); err != nil {
		log.Fatal(err)
	}
}
