package controllers

import (
	"Build_API/models"
	"Build_API/services"

	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// func create customer record  --- CREATE
func CreateCustomerRecord(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var temp models.Customers
	temp.ID = primitive.NewObjectID()
	_ = json.NewDecoder(r.Body).Decode(&temp)
	services.CreateCustomerRecord(temp,Collection)
	json.NewEncoder(w).Encode(temp)
}


// GET ALL RECORD ---- READ
func GetAllCustomerRecord(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	allrecord := services.GetAllCustomerRecord(Collection)
	json.NewEncoder(w).Encode(allrecord)

}

// update customer password ----- UPDATE
func UpdatePassword(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	w.Header().Set("Allow-Control-Alow-Methods", "POST")
	params := mux.Vars(r)
	services.UpdateCustomerPassword(params["id"],Collection)
	json.NewEncoder(w).Encode(params["id"])
}

// delete customer record ----- DELETE
func DeleteCustomerRecord(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	DeleteCustomerRecord1(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

func DeleteCustomerRecord1(del string) {
	id, _ := primitive.ObjectIDFromHex(del)
	filter := bson.M{"_id": id}
	deletecus, err := Collection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Customer Record has been deleted", deletecus.DeletedCount)
}
