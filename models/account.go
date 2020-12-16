package models

import (
	"encoding/json"
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Account struct {
	AccountID      primitive.ObjectID `json:"account_id,omitempty" bson:"_id,omitempty"`
	DocumentNumber string             `json:"document_number" binding:"required" bson:"document_number"`
}

func (a Account) JSON() []byte {
	j, err := json.Marshal(a)
	if err != nil {
		fmt.Println(err)
	}

	return j
}
