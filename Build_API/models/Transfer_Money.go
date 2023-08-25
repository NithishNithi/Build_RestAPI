package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type MakeTransaction struct{
	Transaction_id primitive.ObjectID `json:"_id" bson:"_id"`
	From_Id int `json:"fromid" bson:"fromid"`
	To_Id int `json:"toid" bson:"toid"`
	Amount int `json:"amount" bson:"amount"`
}