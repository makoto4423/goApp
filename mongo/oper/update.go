package oper

import (
	"context"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

func UpdateMany(collection *mongo.Collection) {

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

func UpdateOne(collection *mongo.Collection) {
	filter := bson.D{
		{"country", "USA"},
	}
	update := bson.D{
		{"$set", bson.D{
			{"skin", "white"},
		}},
		{"$inc", bson.D{
			{"age", 1},
		}},
		{"$set", bson.D{
			{"qty", bson.A{
				60, 70, 90,
			}},
		}},
	}
	res, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			log.Println("No documents found")
			return
		}
		panic(err)
	}
	fmt.Println(res.ModifiedCount)
}

// UpdateIfNotExists { $replaceRoot: { newRoot:
//
//	   { $mergeObjects: [ { quiz1: 0, quiz2: 0, test1: 0, test2: 0 }, "$$ROOT" ] }
//	} },
//	{ $set: { modified: "$$NOW"}  }
func UpdateIfNotExists(collection *mongo.Collection) {
	many, err := collection.UpdateMany(context.TODO(), bson.D{}, bson.A{
		bson.D{
			{"$replaceRoot", bson.D{
				{"newRoot", bson.D{
					{"$mergeObjects", bson.A{
						bson.D{
							{"quiz1", 0},
							{"quiz2", 0},
							{"test1", 0},
						},
						"$$ROOT",
					}},
				}},
			}},
		},
	})
	if err != nil {
		return
	}
	fmt.Println(many)
}
