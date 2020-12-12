package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitTransactionRoutes(router *gin.RouterGroup) {
	router.POST("/", createTransaction)
}

func createTransaction(c *gin.Context) {
	c.JSON(http.StatusOK, "pong")
}
