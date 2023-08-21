package controllers

import (
	constants "Build_API/Constants"
	"Build_API/models"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Collection *mongo.Collection

	func init(){
		// client option
		clientOption := options.Client().ApplyURI(constants.ConnectionString)

		// connect to mongo
		client,err:=mongo.Connect(context.TODO(),clientOption)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("MongoDB Connected Success fully")

		// collection Instance
		Collection = client.Database(constants.DBName).Collection(constants.Collection_name)
		fmt.Println("Collection Instance Ready")
	}

	func InsertCustomerRecord1(record models.Customers){
		inserted,err:=Collection.InsertOne(context.Background(),record)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Inserted 1 Customer_Record",inserted.InsertedID)
	}

	func InsertCustomerRecord(w http.ResponseWriter, r *http.Request){
		var record = models.Customers
		json.NewDecoder(r.Body).Decode(&record)
		InsertCustomerRecord1(record)
		json.NewEncoder(w).Encode(record)
	}


	func UpdateCustomerPassowrd(Pass string){
		id,_:=primitive.ObjectIDFromHex(Pass)
		filter:=bson.M{"_id":id}
		update:=bson.M{"$set":bson.M{"password":"hi123"}}
		result,err:=Collection.UpdateOne(context.Background(),filter,update)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Modified count",result.ModifiedCount)

	}
