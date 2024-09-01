package main

import (
	client2 "goApp/mongo/client"
	"goApp/mongo/oper"
)

func main() {
	client := client2.GetClient()

	collection := client.Database("makoto").Collection("students2")

	oper.Watch(collection)
}
