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


func Check_Id(cusID int, Collection *mongo.Collection)(bool) {
	filter := bson.M{"customer_id": cusID}
	result, err := Collection.CountDocuments(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}

	if result > 0 {
		return true
	}
	return false
}


func CreateTransactionRecord(record models.Transactions, Collection1 *mongo.Collection) {
	result, err := Collection1.InsertOne(context.Background(), &record)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted Customer_Record ---->",result.InsertedID)
}
var enteramount int = 98000
func UpdateTransactionAmount(amountid string, Collection1 *mongo.Collection){
	id,_:=primitive.ObjectIDFromHex(amountid)
	filter := bson.M{"_id":id}
	update:=bson.M{"$set":bson.M{"amount":enteramount}}
	_,err:=Collection1.UpdateOne(context.Background(),filter,update)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Transacction Amount Updated")

}