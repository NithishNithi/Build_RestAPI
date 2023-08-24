package controllers

import (
	"Build_API/userinput"
	"context"
	"fmt"
	"log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Collection, Collection1, Collection2 *mongo.Collection

func init() {
	// client option
	clientOption := options.Client().ApplyURI(userinput.ConnectionString)
	// connect to mongo
	client, err := mongo.Connect(context.TODO(), clientOption)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("MongoDB Connected Success fully")
	// collection Instance
	Collection = client.Database(userinput.DBName).Collection(userinput.Collection_name)
	Collection1 = client.Database(userinput.DBName).Collection(userinput.Transaction_CollectionName)
	Collection2 = client.Database(userinput.DBName).Collection(userinput.Loans_CollectionName)
	fmt.Println("Collection Instance Ready")
}

// func Customers() {
// 	w.Write([]byte("<h1 style>welcome to the Bank</h1>"))
// 	w.Write([]byte("<a href=http://localhost:4000/createcustomer><p1>Create_Customer</p1></a><br>"))
// 	w.Write([]byte("<a href=http://localhost:4000/getallrecord><p1>Get_All_Customer_Record</p1></a><br>"))
// 	w.Write([]byte("<a href=http://localhost:4000/updatepassword><p1>Update_Password</p1></a><br>"))
// 	w.Write([]byte("<a href=http://localhost:4000/deletecustomer><p1>Delete_Customer</p1></a><br>"))
// }

// func Transaction(w http.ResponseWriter, _ *http.Request){
// 	w.Write([]byte("<h1>Insert Customer Data in Postman</h1>"))
// }
