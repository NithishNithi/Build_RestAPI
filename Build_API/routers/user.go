package routers

import (
	"Build_API/controllers"

	"github.com/gorilla/mux"
)

func Routers() *mux.Router{
	router := mux.NewRouter()
	router.HandleFunc("/api/banking/customer",controllers.InsertCustomerRecord)

	return router
}