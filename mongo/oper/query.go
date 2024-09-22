package oper

import (
	"context"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

func FindOne(collection *mongo.Collection) {
	res := collection.FindOne(context.TODO(), bson.D{
		{"forever", bson.D{
			{"$exists", true},
		}},
	})
	raw, err := res.Raw()
	if err != nil {
		log.Print("findOne err: ", err)
		if errors.Is(err, mongo.ErrNoDocuments) {
			return
		}
		panic(err)
	}
	raw, err = bson.MarshalExtJSON(raw, false, false)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(raw))
}

func QueryIn(collection *mongo.Collection) {
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
	//for _, result := range results {
	//	k, _ = bson.MarshalExtJSON(result, false, false)
	//	fmt.Println(string(k))
	//}
}

func Query(collection *mongo.Collection) {
	opt := &options.FindOptions{}
	opt.SetLimit(5). // 限制返回结果
				SetSkip(1). // 跳过结果

				SetProjection( // 限制返回列
			bson.D{
				{"size", 0},
			},
		)
	cur, err := collection.Find(context.TODO(), bson.D{
		{"qty", bson.D{
			{"$gt", 5},
		}},
	}, opt)
	if err != nil {
		panic(err)
	}
	for cur.Next(context.TODO()) {
		var elem map[string]interface{}
		err := cur.Decode(&elem)
		if err != nil {
			panic(err)
		}
		fmt.Println(elem)
	}
}

func Estimate(collection *mongo.Collection) {
	//filter := bson.D{{}}
	estCount, _ := collection.EstimatedDocumentCount(context.TODO())
	fmt.Println(estCount)

}
