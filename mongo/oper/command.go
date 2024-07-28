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
