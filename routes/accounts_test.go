package routes

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/agnjunio/go-rest-demo/database"
	"github.com/agnjunio/go-rest-demo/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var router *gin.Engine
var id = primitive.NewObjectIDFromTimestamp(time.Now())

func TestMain(m *testing.M) {
	// Mock mongodb somehow
	err := godotenv.Load("../.env")
	if err != nil {
		log.Println(".env file not found")
	}

	client, _ := database.Connect()
	router = InitRouter(client)
	collection := database.GetDB(client).Collection(Collection)

	// Ensure test account exists
	testAccount := models.Account{
		AccountID:      id,
		DocumentNumber: "12345321",
	}
	collection.InsertOne(context.TODO(), testAccount)

	os.Exit(m.Run())
}

func TestCreateAccount(t *testing.T) {
	documentNumber := "1234567890"
	body := bytes.NewBuffer(models.Account{DocumentNumber: documentNumber}.JSON())

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/accounts/", body)
	router.ServeHTTP(w, req)

	var createdAccount models.Account
	json.Unmarshal(w.Body.Bytes(), &createdAccount)

	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Equal(t, documentNumber, createdAccount.DocumentNumber)
}

func TestGetAccount(t *testing.T) {
	// Test Not Found
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/accounts/1", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)

	// Test Ok
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/accounts/"+id.Hex(), nil)
	router.ServeHTTP(w, req)

	var account models.Account
	json.Unmarshal(w.Body.Bytes(), &account)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, id, account.AccountID)
}
