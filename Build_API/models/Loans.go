package models

type Loans struct {
	Customer_Id int     `json:"customer_id" bson:"customer_id"`
	Loan_Amount float64 `json:"loan_amount" bson:"loan_amount"`
	Loan_Type        string  `json:"loan_type" bson:"loan_type"`
}
