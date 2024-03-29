package db

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	db "vitalsync/db/metadata"
)

const (
	dbName = "vitalsync"
	uri    = "mongodb://root:root@localhost:27017/"
)

// GetConnection opens a connection with a target Mongo database
// It returns a Mongo client for further operations
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

// GetCollection retrieves or creates a collection from a database
// It returns a Mongo collection for further operations
func GetCollection(connection *mongo.Client, collectionName string) *mongo.Collection {
	collection := connection.Database(dbName).Collection(collectionName)

	return collection
}

// CountAllDocuments counts the number of documents present in a Mongo collection
// It returns an integer counter of documents
func CountAllDocuments(collection *mongo.Collection) (int64, error) {
	numOfDocs, err := collection.CountDocuments(context.Background(), bson.M{})
	if err != nil {
		return 0, err
	}

	return numOfDocs, nil
}

func InsertDocument(collection *mongo.Collection, metaData db.MetaData) (*mongo.InsertOneResult, error) {

	result, err := collection.InsertOne(context.Background(), metaData)

	if err != nil {
		return nil, err
	}

	return result, nil
}
