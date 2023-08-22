package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Customers struct {
	ID          primitive.ObjectID `json:"id" bson:"_id"`
	Customer_Id int                `json:"customer_id" bson:"customer_id"`
	Name        string             `json:"name,omitempty" bson:"name,omitempty"`
	Password    string             `json:"password" bson:"password"`
	Bank_Id     int                `json:"bank_id" bson:"bank_id"`
}
