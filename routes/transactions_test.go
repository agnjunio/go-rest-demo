package routes

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/agnjunio/go-rest-demo/models"
	"github.com/go-playground/assert/v2"
)

func TestCreateTransaction(t *testing.T) {
	// Not enough balance
	transaction1 := models.Transaction{
		Account:       AccountID,
		OperationType: OpTypeID,
		Amount:        -100.50,
	}
	body := bytes.NewBuffer(transaction1.JSON())

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/transactions/", body)
	router.ServeHTTP(w, req)

	var response map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &response)

	assert.Equal(t, w.Code, http.StatusUnprocessableEntity)
	assert.Equal(t, response["message"], "Not enough balance.")

	// Transaction created
	var expectedAmount2 float32 = 100.50
	transaction2 := models.Transaction{
		Account:       AccountID,
		OperationType: OpTypeID,
		Amount:        expectedAmount2,
	}
	body = bytes.NewBuffer(transaction2.JSON())

	w = httptest.NewRecorder()
	req, _ = http.NewRequest("POST", "/transactions/", body)
	router.ServeHTTP(w, req)

	var createdTransaction models.Transaction
	json.Unmarshal(w.Body.Bytes(), &createdTransaction)

	assert.Equal(t, w.Code, http.StatusCreated)
	assert.Equal(t, createdTransaction.Account, AccountID)
	assert.Equal(t, createdTransaction.OperationType, OpTypeID)
	assert.Equal(t, createdTransaction.Amount, expectedAmount2)
}
