package controllers

import (
	"Build_API/models"
	"Build_API/password"
	"Build_API/services"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// func create customer record  --- CREATE
func CreateCustomerRecord(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Insert Customer Data in Postman</h1>"))
	w.Header().Set("Content-Type", "application/json")
	var temp models.Customers
	_ = json.NewDecoder(r.Body).Decode(&temp)
	hashedPassword, err := password.HashPassword(temp.Password)
	if err != nil {
		log.Fatal(err)
	}
	temp.Password = hashedPassword
	services.CreateCustomerRecord(temp, Collection)
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
func UpdateCustomerPassword(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	w.Header().Set("Allow-Control-Alow-Methods", "POST")
	params := mux.Vars(r)
	services.UpdateCustomerPassword(params["id"], Collection)
	json.NewEncoder(w).Encode(params["id"])
}

// delete customer record ----- DELETE
func DeleteCustomerRecord(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	services.DeleteCustomerRecord(params["id"], Collection)
	json.NewEncoder(w).Encode(params["id"])
}
