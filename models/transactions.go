package models

import (
	"encoding/json"
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Transaction struct {
	TransactionID primitive.ObjectID  `json:"transaction_id,omitempty" bson:"_id,omitempty"`
	Account       primitive.ObjectID  `json:"account_id,omitempty" bson:"account,omitempty" binding:"required"`
	OperationType primitive.ObjectID  `json:"operation_type_id,omitempty" bson:"operation_type,omitempty" binding:"required"`
	Amount        float32             `json:"amount,omitempty" bson:"amount,omitempty" binding:"required"`
	EventDate     primitive.Timestamp `json:"event_date,omitempty" bson:"event_date,omitempty"`
}

func (t Transaction) JSON() []byte {
	j, err := json.Marshal(t)
	if err != nil {
		fmt.Println(err)
	}

	return j
}
