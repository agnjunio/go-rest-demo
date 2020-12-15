package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type OperationType struct {
	OpTypeID    primitive.ObjectID
	Description string
}
