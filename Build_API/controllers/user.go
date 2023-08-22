package controllers

import (
	
	"Build_API/userinput"
	"context"
	"fmt"
	"log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Collection *mongo.Collection

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
	fmt.Println("Collection Instance Ready")
}
