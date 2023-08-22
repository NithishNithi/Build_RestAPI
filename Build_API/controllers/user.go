package controllers

import (
	constants "Build_API/Constants"
	"Build_API/models"

	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Collection *mongo.Collection

func init() {
	// client option
	clientOption := options.Client().ApplyURI(constants.ConnectionString)
	// connect to mongo
	client, err := mongo.Connect(context.TODO(), clientOption)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("MongoDB Connected Success fully")
	// collection Instance
	Collection = client.Database(constants.DBName).Collection(constants.Collection_name)
	fmt.Println("Collection Instance Ready")
}

// func create customer record
func CreateCustomerRecord(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Welcome", "to create section")
	var temp models.Customers
	_ = json.NewDecoder(r.Body).Decode(&temp)
	CreateCustomerRecord1(temp)
	// json.NewEncoder(w).Encode(temp)
}

func CreateCustomerRecord1(record models.Customers) {
	_, err := Collection.InsertOne(context.Background(), &record)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted 1 Customer_Record")
}

// upadte customer password

func UpdatePassword(w http.ResponseWriter, r *http.Request){
	params:=mux.Vars(r)
	UpdateCustomerPassoword(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

func UpdateCustomerPassoword(Pass string){
	id,_:=primitive.ObjectIDFromHex(Pass)
	filter:=bson.M{"_id":id}
	update:=bson.M{"$set":bson.M{"password":"passkey"}}
	result,err:=Collection.UpdateOne(context.Background(),filter,update)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Modified count",result.ModifiedCount)
}
