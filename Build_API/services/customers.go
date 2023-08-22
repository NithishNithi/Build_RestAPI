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

// func create customer record  --- CREATE
func CreateCustomerRecord(record models.Customers,collection *mongo.Collection) {
	_, err := collection.InsertOne(context.Background(), &record)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted 1 Customer_Record")
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
func UpdateCustomerPassword(Pass string,Collection *mongo.Collection) {
	id, _ := primitive.ObjectIDFromHex(Pass)
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"password": "132213&2234"}}
	result, err := Collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Modified count", result.ModifiedCount)
	fmt.Printf("Modified count: %d\n", result.ModifiedCount)
}

// Delete customer record ----- Delete