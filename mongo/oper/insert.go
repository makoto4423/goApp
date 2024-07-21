package oper

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
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
	res, err := collection.InsertOne(context.TODO(), docs)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Inserted a single document: %v\n", res.InsertedID)
}
