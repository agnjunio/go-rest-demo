package routes

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/agnjunio/go-rest-demo/models"
	"github.com/go-playground/assert/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestCreateAccount(t *testing.T) {
	documentNumber := "TestCreateAccount"
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
	req, _ := http.NewRequest("GET", "/accounts/"+primitive.NewObjectID().Hex(), nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)

	// Test Ok
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/accounts/"+AccountID.Hex(), nil)
	router.ServeHTTP(w, req)

	var account models.Account
	json.Unmarshal(w.Body.Bytes(), &account)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, AccountID, account.AccountID)
}
