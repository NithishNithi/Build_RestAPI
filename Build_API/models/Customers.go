package models

type Customers struct {
	Customer_Id int                `json:"customer_id" bson:"customer_id"`
	Name        string             `json:"name,omitempty" bson:"name,omitempty"`
	Password    string             `json:"password" bson:"password"`
	Bank_Id     int                `json:"bank_id" bson:"bank_id"`
}
