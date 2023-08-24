package models

import "time"

type Transactions struct {
	Customer_Id      int     `json:"customer_id" bson:"customer_id"`
	Amount           float64 `json:"amount" bson:"amount"`
	Transaction_Time time.Time  `json:"transaction_time" bson:"transaction_time"`
	Transaction_Type string  `json:"transaction_type" bson:"transaction_type"`
}
