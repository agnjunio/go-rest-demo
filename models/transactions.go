package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Transaction struct {
	TransactionID primitive.ObjectID
	Account       primitive.ObjectID
	OperationType primitive.ObjectID
	Amount        int
	EventDate     primitive.Timestamp
}
