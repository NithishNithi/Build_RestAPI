package routers

import (
	"Build_API/controllers"

	"github.com/gin-gonic/gin"
)

// func Routers() *mux.Router {
// 	// Create a new Mux router instance
// 	router := mux.NewRouter()
// 	// Define routes and associate them with controller functions

// // customers ------>
// 	router.HandleFunc("/customers",controllers.Customers)
// 	router.HandleFunc("/customers/createcustomer", controllers.CreateCustomerRecord).Methods("POST")
// 	router.HandleFunc("/customers/getallrecord", controllers.GetAllCustomerRecord).Methods("GET")
// 	router.HandleFunc("/customers/updatepassword/{id}", controllers.UpdateCustomerPassword).Methods("PUT")
// 	router.HandleFunc("/customers/deletecustomer/{id}", controllers.DeleteCustomerRecord).Methods("DELETE")

// // Transactions ------>
// 	router.HandleFunc("/transactions",controllers.Transaction)
// 	router.HandleFunc("/transactions/createtransaction", controllers.CreateTransaction).Methods("POST")
// 	router.HandleFunc("/transactions/updatetransactionamount/{id}",controllers.UpdateTransactionAmount).Methods("PUT")
// 	return router
// }

func SetUpRouters(r *gin.Engine){
	
	r.POST("/customers/create", controllers.CreateCustomerRecord)
	r.GET("/customers/getall", controllers.GetAllCustomerRecord)
	r.PUT("/customers/update/:id", controllers.UpdateCustomerPassword)
	r.DELETE("/customers/delete/:id", controllers.DeleteCustomerRecord)

	// Transactions ------>
	
	r.POST("/transactions/create", controllers.CreateTransaction)
	r.PUT("/transactions/update/:id", controllers.UpdateTransactionAmount)
	r.GET("/transactions/getall",controllers.GetAllTransaction)
	r.POST("/transactions/timerange", controllers.GetTransactionsByTimeRange)
	r.POST("/transactions/sumbytimerange", controllers.GetTransactionsSumByTimeRange)

	// Loans ------>
	r.POST("/loans/create", controllers.CreateLoanRecord)
	r.PUT("/loans/update/:id", controllers.UpdateLoanRecord)
	r.GET("/loans/getall", controllers.GetAllLoanRecord)
	r.DELETE("/loans/update/:id", controllers.DeleteLoanRecord)
}
