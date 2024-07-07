package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const uri = "mongodb://localhost:27017"

func main() {
	// Use the SetServerAPIOptions() method to set the Stable API version to 1
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI).SetAppName("helloWorld")
	// Create a new client and connect to the server
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
	// Send a ping to confirm a successful connection
	var result bson.M
	if err := client.Database("admin").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Decode(&result); err != nil {
		panic(err)
	}
	collection := client.Database("makoto").Collection("orders")

	ans := collection.FindOne(context.TODO(), bson.D{{"state", "Texas"}})

	raw, _ := ans.Raw()
	k, err := bson.MarshalExtJSON(raw, false, false)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(k))

	cursor, err := collection.Find(context.TODO(), bson.D{})
	cursor.SetBatchSize(10)
	fmt.Println("here")
	if err != nil {
		panic(err)
	}
	var results []interface{}
	if err = cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}
	for _, result := range results {
		k, _ = bson.MarshalExtJSON(result, false, false)
		fmt.Println(string(k))
	}
}
