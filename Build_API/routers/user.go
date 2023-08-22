package routers

import (
	"Build_API/controllers"

	"github.com/gorilla/mux"
)

func Routers() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/api/customers", controllers.CreateCustomerRecord).Methods("POST")
	router.HandleFunc("api/updateCus{id}",controllers.UpdatePassword)

	return router
}
