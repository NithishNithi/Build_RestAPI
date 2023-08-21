package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Customers struct {
	ID primitive.ObjectID `json:"_id" bson:"_id"`
	Name string `json:"name,omitempty" bson:"name,omitempty"`
	Password string `json:"password" bson:"password"`
	Bank_Id int `json:"bank_id" bson:"bank_id"`
}