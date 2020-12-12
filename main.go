package main

import (
	"go-rest-demo/routes"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println(".env file not found")
	}

	router := initRouter()

	err = router.Run()
	if err != nil {
		log.Fatal("Failed to start server")
	}
}

func initRouter() *gin.Engine {
	router := gin.Default()
	routes.InitAccountRoutes(router.Group("/accounts"))
	routes.InitTransactionRoutes(router.Group("/transactions"))

	return router
}
