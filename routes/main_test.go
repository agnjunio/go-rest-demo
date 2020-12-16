package routes

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/agnjunio/go-rest-demo/database"
	"github.com/agnjunio/go-rest-demo/models"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var router *gin.Engine
var AccountID, _ = primitive.ObjectIDFromHex("111111111111111111111111")
var OpTypeID, _ = primitive.ObjectIDFromHex("111111111111111111111111")

func TestMain(m *testing.M) {
	// Warm up
	// Mock mongodb somehow
	err := godotenv.Load("../.env")
	if err != nil {
		log.Println(".env file not found")
	}

	client, _ := database.Connect()
	router = InitRouter(client)
	accColl := database.GetDB(client).Collection(database.AccountsCollection)
	trColl := database.GetDB(client).Collection(database.TransactionsCollection)
	opTypeColl := database.GetDB(client).Collection(database.OperationTypesCollection)
	upsertOpts := options.Update().SetUpsert(true)

	// Ensure test account exists
	testAccount := models.Account{
		AccountID:      AccountID,
		DocumentNumber: "TestAccount",
	}
	accColl.UpdateOne(context.TODO(), bson.M{"_id": AccountID}, bson.M{"$set": testAccount}, upsertOpts)

	// Ensure test operation type exists
	testOpType := models.OperationType{
		OpTypeID:    OpTypeID,
		Description: "TestOperationType",
	}
	opTypeColl.UpdateOne(context.TODO(), bson.M{"_id": OpTypeID}, bson.M{"$set": testOpType}, upsertOpts)

	// Delete all transactions from previous tests
	trColl.DeleteMany(context.TODO(), bson.M{"account": AccountID})

	exitCode := m.Run()

	// Tear-down
	accColl.DeleteOne(context.TODO(), bson.M{"_id": AccountID})
	accColl.DeleteMany(context.TODO(), bson.M{"document_number": "TestCreateAccount"})
	opTypeColl.DeleteOne(context.TODO(), bson.M{"_id": OpTypeID})
	trColl.DeleteMany(context.TODO(), bson.M{"account": AccountID})

	os.Exit(exitCode)
}
