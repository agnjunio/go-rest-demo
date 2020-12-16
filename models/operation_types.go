package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type OperationType struct {
	OpTypeID    primitive.ObjectID `json:"operation_type_id,omitempty" bson:"_id,omitempty"`
	Description string             `json:"description,omitempty" bson:"description,omitempty"`
}
