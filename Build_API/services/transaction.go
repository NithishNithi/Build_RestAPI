package services

import (
	"Build_API/models"
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Check_Id
func Transaction_Check_Id(cusID int, Collection *mongo.Collection) bool {
	filter := bson.M{"customer_id": cusID}
	result, err := Collection.CountDocuments(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}

	return result > 0
}

// CreateTransactionRecord
func CreateTransactionRecord(record models.Transactions, Collection1 *mongo.Collection) {
	result, err := Collection1.InsertOne(context.Background(), &record)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted Transaction_Record ---->", result.InsertedID)
}

// UpdateTransactionAmount
var enter_Trans_amount int = 98000

func UpdateTransactionAmount(amountid string, Collection1 *mongo.Collection) {
	id, _ := primitive.ObjectIDFromHex(amountid)
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"amount": enter_Trans_amount}}

	_, err := Collection1.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Transaction Amount Updated")
}

// GetAllTransactionRecord

func GetAllTransactionRecord(Collection1 *mongo.Collection) []primitive.M {
	cur, err := Collection1.Find(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}

	var trans_rec []primitive.M

	for cur.Next(context.Background()) {
		var record bson.M
		err := cur.Decode(&record)
		if err != nil {
			log.Fatal(err)
		}
		trans_rec = append(trans_rec, record)

	}
	return trans_rec
}

func GetTransactionsByTimeRange(startTime, endTime time.Time, Collection1 *mongo.Collection) ([]*models.Transactions, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	filter := bson.M{
		"transaction_time": bson.M{
			"$gte": startTime,
			"$lte": endTime,
		},
	}

	options := options.Find()
	cur, err := Collection1.Find(ctx, filter, options)
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx) // Close the cursor when done

	var transactions []*models.Transactions
	for cur.Next(ctx) {
		transaction := &models.Transactions{}
		if err := cur.Decode(transaction); err != nil {
			return nil, err
		}
		transactions = append(transactions, transaction)
	}

	if err := cur.Err(); err != nil {
		return nil, err
	}

	return transactions, nil
}
func GetTransactionsSumByTimeRange(startTime, endTime time.Time, Collection1 *mongo.Collection) (float64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	pipeline := []bson.M{
		{
			"$match": bson.M{
				"transaction_time": bson.M{
					"$gte": startTime,
					"$lte": endTime,
				},
			},
		},
		{
			"$group": bson.M{
				"_id": nil,
				"totalAmount": bson.M{
					"$sum": "$amount", // Assuming the transaction amount field is named "amount"
				},
			},
		},
	}

	cur, err := Collection1.Aggregate(ctx, pipeline)
	if err != nil {
		return 0, err
	}
	defer cur.Close(ctx)

	var result struct {
		TotalAmount float64 `bson:"totalAmount"`
	}

	if cur.Next(ctx) {
		if err := cur.Decode(&result); err != nil {
			return 0, err
		}
	}

	return result.TotalAmount, nil
}
