package oper

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/writeconcern"
	"strconv"
	"time"
)

func Insert(collection *mongo.Collection) {
	docs := []interface{}{bson.D{
		{"item", "journal"},
		{"qty", 25},
		{"size", bson.D{
			{"h", 14},
			{"w", 21},
			{"uom", "cm"},
		}},
		{"status", "A"},
	},
		bson.D{
			{"item", "notebook"},
			{"qty", 50},
			{"size", bson.D{
				{"h", 8.5},
				{"w", 11},
				{"uom", "in"},
			}},
			{"status", "A"},
		},
		bson.D{
			{"item", "paper"},
			{"qty", 100},
			{"size", bson.D{
				{"h", 8.5},
				{"w", 11},
				{"uom", "in"},
			}},
			{"status", "D"},
		},
		bson.D{
			{"item", "planner"},
			{"qty", 75},
			{"size", bson.D{
				{"h", 22.85},
				{"w", 30},
				{"uom", "cm"},
			}},
			{"status", "D"},
		},
		bson.D{
			{"item", "postcard"},
			{"qty", 45},
			{"size", bson.D{
				{"h", 10},
				{"w", 15.25},
				{"uom", "cm"},
			}},
			{"status", "A"},
		}}
	result, err := collection.InsertMany(context.TODO(), docs)
	if err != nil {
		panic(err)
	}
	for i, val := range result.InsertedIDs {
		fmt.Println(strconv.Itoa(i), val.(primitive.ObjectID).String())
	}
}

func InsertOne(collection *mongo.Collection) {

	// insertOne
	docs := bson.D{
		{"forever", time.Now()},
	}
	collection.Database().WriteConcern().W = 2
	res, err := collection.InsertOne(context.TODO(), docs)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Inserted a single document: %v\n", res.InsertedID)
}

func Transaction(client *mongo.Client) {
	session, err := client.StartSession()
	if err != nil {
		panic(err)
	}
	defer func() {
		session.EndSession(context.TODO())
	}()
	err = session.StartTransaction()
	if err != nil {
		panic(err)
	}
	wc := writeconcern.Majority()
	txnOptions := options.Transaction().SetWriteConcern(wc)
	docs := []interface{}{
		bson.D{{"title", "The Bluest Eye"}, {"author", "Toni Morrison"}},
		bson.D{{"title", "Sula"}, {"author", "Toni Morrison"}},
		bson.D{{"title", "Song of Solomon"}, {"author", "Toni Morrison"}},
	}
	collection := client.Database("makoto").Collection("clash")
	_, err = session.WithTransaction(context.TODO(), func(ctx mongo.SessionContext) (interface{}, error) {
		result, err := collection.InsertMany(ctx, docs)
		return result, err
	}, txnOptions)
	//err = session.CommitTransaction(context.TODO())
	//if err != nil {
	//	return
	//}
}
