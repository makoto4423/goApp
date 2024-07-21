package oper

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func Update(collection *mongo.Collection) {

	result, err := collection.UpdateMany(context.TODO(), bson.D{
		{"qty", bson.D{
			{"$mod", bson.A{10, 5}},
		}},
	}, bson.D{
		{"$set", bson.D{
			{"country", "USA"},
		}},
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(result.ModifiedCount)
}
