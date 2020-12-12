package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitAccountRoutes(router *gin.RouterGroup) {
	router.POST("/", createAccount)
	router.GET("/:id", getAccount)
}

func getAccount(c *gin.Context) {
	c.JSON(http.StatusOK, "pong")
}

func createAccount(c *gin.Context) {
	c.JSON(http.StatusOK, "pong")
}
