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

// check id from Customer collection , wheather the given id already exist in customers
func Loan_Check_Id(cusID int, Collection *mongo.Collection) bool {
	filter := bson.M{"customer_id": cusID}
	result, err := Collection.CountDocuments(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}

	return result > 0
}

// Create Loans Record
func CreateLoanRecord(record models.Loans, Collection1 *mongo.Collection) {
	result, err := Collection1.InsertOne(context.Background(), &record)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted Transaction_Record ---->", result.InsertedID)
}

// Update Loan Amount
var enter_Loan_amount int = 4000
func UpdateLoanAmount(amountid string, Collection2 *mongo.Collection) {
	id, _ := primitive.ObjectIDFromHex(amountid)
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"loan_amount": enter_Loan_amount}}

	_, err := Collection2.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Loan Amount Updated")
}

// Get All Update Record 

func GetAllLoanRecord(Collection2 *mongo.Collection) []primitive.M {
	cur, err := Collection2.Find(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}
	var cus_rec []primitive.M
	for cur.Next(context.Background()) {
		var record bson.M
		err := cur.Decode(&record)
		if err != nil {
			log.Fatal(err)
		}
		cus_rec = append(cus_rec, record)
	}
	fmt.Println("Done")
	defer cur.Close(context.Background())
	return cus_rec
}

// Delete Loan Record

func DeleteLoanRecord(del string, Collection2 *mongo.Collection) {
	id, _ := primitive.ObjectIDFromHex(del)
	filter := bson.M{"_id": id}
	deletecus, err := Collection2.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Loan Record has been deleted", deletecus.DeletedCount)
}