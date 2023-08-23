package services

import (
	"Build_API/models"
	"Build_API/password"
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)



// func create customer record  --- CREATE
func CreateCustomerRecord(record models.Customers, collection *mongo.Collection) {
	_, err := collection.InsertOne(context.Background(), &record)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted Customer_Record")
}

// GET ALL RECORD ---- READ
func GetAllCustomerRecord(Collection *mongo.Collection) []primitive.M {
	cur, err := Collection.Find(context.Background(), bson.D{{}})
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

// Update customer password ----- UPDATE
var enterPass string  = "passkey102"
func UpdateCustomerPassword(Pass string, Collection *mongo.Collection) {
	id, _ := primitive.ObjectIDFromHex(Pass)
	hashedPassword, _ := password.HashPassword(enterPass)
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"password": hashedPassword}}

	result, err := Collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Modified count", result.ModifiedCount)
}

// Delete customer record ----- Delete
func DeleteCustomerRecord(del string, Collection *mongo.Collection) {
	id, _ := primitive.ObjectIDFromHex(del)
	filter := bson.M{"_id": id}
	deletecus, err := Collection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Customer Record has been deleted", deletecus.DeletedCount)
}
