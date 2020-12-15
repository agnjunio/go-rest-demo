package routes

import "github.com/gin-gonic/gin"

func InitRouter() *gin.Engine {
	router := gin.Default()
	InitAccountRoutes(router.Group("/accounts"))
	InitTransactionRoutes(router.Group("/transactions"))

	return router
}
