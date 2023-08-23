package models



type Transactions struct {
	
	Customer_Id int `json:"customer_id" bson:"customer_id"` 
	Amount     float64           `json:"amount" bson:"amount"`
}
