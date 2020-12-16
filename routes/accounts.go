package routes

import (
	"context"
	"net/http"

	"github.com/agnjunio/go-rest-demo/database"
	"github.com/agnjunio/go-rest-demo/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func InitAccountRoutes(router *gin.RouterGroup) {
	router.POST("/", createAccount)
	router.GET("/:id", getAccount)
}

func getAccount(c *gin.Context) {
	id, _ := primitive.ObjectIDFromHex(c.Param("id"))
	client := c.MustGet("mongo").(*mongo.Client)
	collection := database.GetDB(client).Collection(database.AccountsCollection)

	account := models.Account{}
	err := collection.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&account)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Account not found.",
			"error":   err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, account)
	}
}

func createAccount(c *gin.Context) {
	var account models.Account

	if err := c.BindJSON(&account); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Request validation failed.",
			"error":   err.Error(),
		})
	} else {
		client := c.MustGet("mongo").(*mongo.Client)
		collection := database.GetDB(client).Collection(database.AccountsCollection)

		result, err := collection.InsertOne(context.TODO(), account)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Account creation failed.",
				"error":   err.Error(),
			})
		} else {
			account.AccountID = result.InsertedID.(primitive.ObjectID)
			c.JSON(http.StatusCreated, account)
		}
	}
}
