package routers

import (
	"Build_API/controllers"
	"github.com/gorilla/mux"
)

func Routers() *mux.Router {
	// Create a new Mux router instance
	router := mux.NewRouter()

	// Define routes and associate them with controller functions
	router.HandleFunc("/api/createcustomer", controllers.CreateCustomerRecord).Methods("POST")
	router.HandleFunc("/api/getallrecord", controllers.GetAllCustomerRecord).Methods("GET") 
	router.HandleFunc("/api/updatepassword/{id}", controllers.UpdatePassword).Methods("PUT")
	router.HandleFunc("/api/deletecustomer/{id}", controllers.DeleteCustomerRecord).Methods("DELETE") 
	return router
}
