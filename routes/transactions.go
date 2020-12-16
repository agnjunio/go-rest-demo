package routes

import (
	"context"
	"net/http"
	"time"

	db "github.com/agnjunio/go-rest-demo/database"
	"github.com/agnjunio/go-rest-demo/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func InitTransactionRoutes(router *gin.RouterGroup) {
	router.POST("/", createTransaction)
}

func createTransaction(c *gin.Context) {
	var transaction models.Transaction

	if err := c.BindJSON(&transaction); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Request validation failed.",
			"error":   err.Error(),
		})

		return
	}

	client := c.MustGet("mongo").(*mongo.Client)
	accounts := db.GetDB(client).Collection(db.AccountsCollection)
	transactions := db.GetDB(client).Collection(db.TransactionsCollection)
	operationTypes := db.GetDB(client).Collection(db.OperationTypesCollection)

	// Validate amount is not zero
	if transaction.Amount == 0 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": "Amount cannot be zero.",
			"error":   nil,
		})

		return
	}

	// Validate account exists
	var account models.Account

	err := accounts.FindOne(context.TODO(), bson.M{"_id": transaction.Account}).Decode(&account)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Account not found.",
			"error":   err.Error(),
		})

		return
	}

	// Validate operation type
	var opType models.OperationType

	err = operationTypes.FindOne(context.TODO(), bson.M{"_id": transaction.OperationType}).Decode(&opType)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Operation type not found.",
			"error":   err.Error(),
		})

		return
	}

	// Validate account will have non-negative balance after this transaction
	cursor, _ := transactions.Find(context.TODO(), bson.M{"account": account.AccountID}, nil)

	var accTransactions []models.Transaction

	if err = cursor.All(context.TODO(), &accTransactions); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error retrieving account transactions.",
			"error":   err.Error(),
		})

		return
	}

	balance := transaction.Amount
	for _, t := range accTransactions {
		balance += t.Amount
	}

	if balance < 0 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": "Not enough balance.",
			"error":   nil,
		})

		return
	}

	// Set timestamp
	transaction.EventDate = primitive.Timestamp{
		T: uint32(time.Now().Unix()),
		I: 0,
	}

	result, err := transactions.InsertOne(context.TODO(), transaction)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Transaction creation failed.",
			"error":   err.Error(),
		})
	} else {
		transaction.TransactionID = result.InsertedID.(primitive.ObjectID)
		c.JSON(http.StatusCreated, transaction)
	}
}
