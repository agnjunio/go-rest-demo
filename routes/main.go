package routes

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func InitRouter(mongo *mongo.Client) *gin.Engine {
	router := gin.Default()

	if mongo != nil {
		router.Use(func(c *gin.Context) {
			c.Set("mongo", mongo)
			c.Next()
		})
	}

	InitAccountRoutes(router.Group("/accounts"))
	InitTransactionRoutes(router.Group("/transactions"))

	return router
}
