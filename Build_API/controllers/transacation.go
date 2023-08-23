package controllers

import (
	"Build_API/models"
	"Build_API/services"
	"encoding/json"
	"fmt"
	"net/http"
)

func CreateTransaction(w http.ResponseWriter, r *http.Request){
	var temp models.Transactions
	_ = json.NewDecoder(r.Body).Decode(&temp)
	cusID:=temp.Customer_Id
	check:=services.Check_Id(cusID,Collection)

	if !check {
		fmt.Println("Create Customer Record to Create Transaction")
	} else {
		services.CreateTransactionRecord(temp,Collection1)
	}

}


