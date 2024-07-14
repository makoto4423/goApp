package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

	collection := client.Database("makoto").Collection("clash")

	// insert
	//docs := []interface{}{bson.D{
	//	{"item", "journal"},
	//	{"qty", 25},
	//	{"size", bson.D{
	//		{"h", 14},
	//		{"w", 21},
	//		{"uom", "cm"},
	//	}},
	//	{"status", "A"},
	//},
	//	bson.D{
	//		{"item", "notebook"},
	//		{"qty", 50},
	//		{"size", bson.D{
	//			{"h", 8.5},
	//			{"w", 11},
	//			{"uom", "in"},
	//		}},
	//		{"status", "A"},
	//	},
	//	bson.D{
	//		{"item", "paper"},
	//		{"qty", 100},
	//		{"size", bson.D{
	//			{"h", 8.5},
	//			{"w", 11},
	//			{"uom", "in"},
	//		}},
	//		{"status", "D"},
	//	},
	//	bson.D{
	//		{"item", "planner"},
	//		{"qty", 75},
	//		{"size", bson.D{
	//			{"h", 22.85},
	//			{"w", 30},
	//			{"uom", "cm"},
	//		}},
	//		{"status", "D"},
	//	},
	//	bson.D{
	//		{"item", "postcard"},
	//		{"qty", 45},
	//		{"size", bson.D{
	//			{"h", 10},
	//			{"w", 15.25},
	//			{"uom", "cm"},
	//		}},
	//		{"status", "A"},
	//	}}
	//result, err := collection.InsertMany(context.TODO(), docs)
	//if err != nil {
	//	panic(err)
	//}
	//for i, val := range result.InsertedIDs {
	//	fmt.Println(strconv.Itoa(i), val.(primitive.ObjectID).String())
	//}

	//// insertOne
	//docs := bson.D{
	//	{"forever", time.Now()},
	//}
	//res, err := collection.InsertOne(context.TODO(), docs)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Printf("Inserted a single document: %v\n", res.InsertedID)

	//res := collection.FindOne(context.TODO(), bson.D{
	//	{"forever", bson.D{
	//		{"$exists", true},
	//	}},
	//})
	//raw, err := res.Raw()
	//if err != nil {
	//	panic(err)
	//}
	//raw, err = bson.MarshalExtJSON(raw, false, false)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(string(raw))

	// in 查询
	userId, _ := primitive.ObjectIDFromHex("66937a0e9ae5b8749aaab2be")
	cur, err := collection.Find(context.TODO(), bson.D{
		{"_id", bson.D{
			{"$in", bson.A{
				userId,
			}},
		}},
	})
	if err != nil {
		panic(err)
	}
	defer func() {
		err = cur.Close(context.TODO())
		if err != nil {
			panic(err)
		}
	}()
	for cur.Next(context.TODO()) {
		var result map[string]interface{}
		err := cur.Decode(&result)
		if err != nil {
			panic(err)
		}
		fmt.Println(result)
	}

	// query
	//opt := &options.FindOptions{}
	//opt.SetLimit(5). // 限制返回结果
	//			SetSkip(1). // 跳过结果
	//
	//			SetProjection( // 限制返回列
	//		bson.D{
	//			{"size", 0},
	//		},
	//	)
	//cur, err := collection.Find(context.TODO(), bson.D{
	//	{"qty", bson.D{
	//		{"$gt", 5},
	//	}},
	//}, opt)
	//if err != nil {
	//	panic(err)
	//}
	//for cur.Next(context.TODO()) {
	//	var elem map[string]interface{}
	//	err := cur.Decode(&elem)
	//	if err != nil {
	//		panic(err)
	//	}
	//	fmt.Println(elem)
	//}

	//// Send a ping to confirm a successful connection
	//var result bson.M
	//if err := client.Database("admin").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Decode(&result); err != nil {
	//	panic(err)
	//}
	//collection := client.Database("makoto").Collection("orders")
	//
	//ans := collection.FindOne(context.TODO(), bson.D{{"state", "Texas"}})
	//
	//raw, _ := ans.Raw()
	//k, err := bson.MarshalExtJSON(raw, false, false)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(string(k))
	//
	//cursor, err := collection.Find(context.TODO(), bson.D{})
	//cursor.SetBatchSize(10)
	//fmt.Println("here")
	//if err != nil {
	//	panic(err)
	//}
	//var results []interface{}
	//if err = cursor.All(context.TODO(), &results); err != nil {
	//	panic(err)
	//}
	//for _, result := range results {
	//	k, _ = bson.MarshalExtJSON(result, false, false)
	//	fmt.Println(string(k))
	//}

	//package main
	//
	//import (
	//	"context"
	//	"fmt"
	//	"log"
	//
	//	"go.mongodb.org/mongo-driver/bson"
	//	"go.mongodb.org/mongo-driver/bson/primitive"
	//	"go.mongodb.org/mongo-driver/mongo"
	//	"go.mongodb.org/mongo-driver/mongo/options"
	//)
	//
	//type User struct {
	//	ID    primitive.ObjectID `bson:"_id"`
	//	Name  string             `bson:"name"`
	//	Age   int                `bson:"age"`
	//	Email string             `bson:"email"`
	//}
	//
	//func main() {
	//	// 连接 MongoDB
	//	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//	defer client.Disconnect(context.TODO())
	//
	//	// 获取数据库和集合
	//	collection := client.Database("your_database_name").Collection("users")
	//
	//	// 根据 ID 查询用户
	//	userID, _ := primitive.ObjectIDFromHex("5d3e5b1c9c9d440000c0ffee")
	//	var user User
	//	err = collection.FindOne(context.TODO(), bson.M{"_id": userID}).Decode(&user)
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//
	//	fmt.Println(user)
	//}
}
