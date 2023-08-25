package services

import (
	"Build_API/models"
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Customers1 struct {
	Customer_Id    int    `json:"customer_id" bson:"customer_id"`
	Name           string `json:"name,omitempty" bson:"name,omitempty"`
	Password       string `json:"password" bson:"password"`
	Bank_Id        int    `json:"bank_id" bson:"bank_id"`
	Balance_amount int    `json:"balance_amount" bson:"balance_amount"`
}
type Customers2 struct {
	Customer_Id    int    `json:"customer_id" bson:"customer_id"`
	Name           string `json:"name,omitempty" bson:"name,omitempty"`
	Password       string `json:"password" bson:"password"`
	Bank_Id        int    `json:"bank_id" bson:"bank_id"`
	Balance_amount int    `json:"balance_amount" bson:"balance_amount"`
}

func TransferMoney(from int, to int, amount int, Collection *mongo.Collection, Collection3 *mongo.Collection) (int, int) {

	filter1 := bson.M{"customer_id": from}
	var customer1 Customers1
	err1 := Collection.FindOne(context.Background(), filter1).Decode(&customer1)
	if err1 != nil {
		log.Fatal(err1)
	}

	filter2 := bson.M{"customer_id": to}
	var customer2 Customers2
	err2 := Collection.FindOne(context.Background(), filter2).Decode(&customer2)
	if err2 != nil {
		log.Fatal(err2)
	}

	if amount <= customer1.Balance_amount && amount > 0 {
		customer2.Balance_amount += amount
		customer1.Balance_amount -= amount
	} else {
		panic("Insufficient Funds ")
	}
	from_balance := customer1.Balance_amount
	to_balance := customer2.Balance_amount

	//  update balance in cus 1
	cus1_filter := bson.M{"customer_id": from}
	cus1_update := bson.M{"$set": bson.M{"balance_amount": from_balance}}
	_, cus1_err := Collection.UpdateOne(context.Background(), cus1_filter, cus1_update)
	if cus1_err != nil {
		log.Fatal(cus1_err)
	}

	//  update balance in cus 1
	cus2_filter := bson.M{"customer_id": to}
	cus2_update := bson.M{"$set": bson.M{"balance_amount": to_balance}}
	_, cus2_err := Collection.UpdateOne(context.Background(), cus2_filter, cus2_update)
	if cus1_err != nil {
		log.Fatal(cus2_err)
	}

	// Insert the Transaction record to MakeTrasaction table in database

	transactionRecord := models.MakeTransaction{}
	transactionRecord.Transaction_id = primitive.NewObjectID()
	transactionRecord.From_Id = from
	transactionRecord.To_Id = to
	transactionRecord.Amount = amount
	_, transErr := Collection3.InsertOne(context.Background(), transactionRecord)
	if transErr != nil {
		fmt.Println("Error while inserting into transactions Table", transErr)
	} else {
		fmt.Println("Insertion Successfully Done")
	}

	return from_balance, to_balance

}
