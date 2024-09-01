package oper

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func RunCommand(collection *mongo.Collection) {
	// todo fsyncLock, fsyncUnLock
	RunCommand0(collection, bson.D{
		{"isMaster", 1},
	})
	//fmt.Println("lock")
	//time.Sleep(time.Second * 5)
	//RunCommand0(collection, bson.D{
	//	{"fsyncUnLock", 1},
	//})
	//fmt.Println("unlock")
}

func RunCommand0(collection *mongo.Collection, command interface{}) {
	db := collection.Database()
	result := db.RunCommand(context.TODO(), command)
	err := result.Err()
	if err != nil {
		panic(err)
	}
	var out map[string]interface{}
	err = result.Decode(&out)
	if err != nil {
		panic(err)
	}
	fmt.Println(out)
}

func CreateIndex(collection *mongo.Collection) {
	name, err := collection.Indexes().CreateOne(context.TODO(), mongo.IndexModel{
		Keys: bson.D{{"item", 1}},
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(name)
}

// Watch 只有在分片下才能使用
func Watch(collection *mongo.Collection) {
	watch, err := collection.Watch(context.TODO(), mongo.Pipeline{
		bson.D{
			{"$match", bson.D{{
				"operationType", "insert",
			}}},
		},
	})
	if err != nil {
		panic(err)
		return
	}
	i := 0
	for watch.Next(context.TODO()) {
		var m map[string]interface{}
		err := watch.Decode(&m)
		if err != nil {
			return
		}
		fmt.Println(m)
		i++
		if i == 2 {
			break
		}
	}
	defer func() {
		err = watch.Close(context.TODO())
		if err != nil {
			panic(err)
		}
	}()
}
