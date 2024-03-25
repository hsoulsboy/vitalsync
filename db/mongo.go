package db

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	dbName = "vitalsync"
	uri    = "mongodb://root:root@localhost:27017/"
)

func GetConnection() (*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI(uri)

	client, err := mongo.Connect(context.Background(), clientOptions)

	if err != nil {
		return nil, err
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		return nil, err
	}

	return client, nil
}

func GetCollection(connection *mongo.Client, collectionName string) *mongo.Collection {
	collection := connection.Database(dbName).Collection(collectionName)

	return collection
}

func CountAllDocuments(collection *mongo.Collection) (int64, error) {
	numOfDocs, err := collection.CountDocuments(context.Background(), bson.M{})
	if err != nil {
		return 0, err
	}

	return numOfDocs, nil
}
