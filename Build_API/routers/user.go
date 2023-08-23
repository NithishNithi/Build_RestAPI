package routers

import (
	"Build_API/controllers"
	"github.com/gorilla/mux"
)

func Routers() *mux.Router {
	// Create a new Mux router instance
	router := mux.NewRouter()

	// Define routes and associate them with controller functions

// customers ------>
	router.HandleFunc("/customers",controllers.Customers)
	router.HandleFunc("/customers/createcustomer", controllers.CreateCustomerRecord).Methods("POST")
	router.HandleFunc("/customers/getallrecord", controllers.GetAllCustomerRecord).Methods("GET") 
	router.HandleFunc("/customers/updatepassword/{id}", controllers.UpdatePassword).Methods("PUT")
	router.HandleFunc("/customers/deletecustomer/{id}", controllers.DeleteCustomerRecord).Methods("DELETE") 

// Transactions
	router.HandleFunc("/transactions/createtransaction", controllers.CreateTransaction).Methods("POST")

	return router
}
