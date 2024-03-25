package main

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const uri = "mongodb://root:admin@localhost:27017/testa"

func main() {

	clientOptions := options.Client().ApplyURI(uri)

	client, err := mongo.Connect(context.Background(), clientOptions)

	if err != nil {
		fmt.Println(err)
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Connected to Mongodb")

	collection := client.Database("vitalsync").Collection("sensor_data")

	numOfDocs, err := collection.CountDocuments(context.Background(), bson.M{})
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(numOfDocs)
}
