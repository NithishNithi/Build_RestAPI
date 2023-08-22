package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Transactions struct {
	ID         primitive.ObjectID `json:"id" bson:"_id"`
	CustomerID primitive.ObjectID `json:"customer_id" bson:"customer_id"` // Foreign key
	Amount     float64           `json:"amount" bson:"amount"`
}
