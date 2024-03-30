package db

import (
	"context"
	"fmt"

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

// InsertDocument adds a new document to a given collection
// It returns the entry ID
func InsertDocument(collection *mongo.Collection, metaData db.MetaData) (*mongo.InsertOneResult, error) {

	result, err := collection.InsertOne(context.Background(), metaData)

	if err != nil {
		return nil, err
	}

	return result, nil
}

// GetPatientBySSN retrieves a Patient given its Social Security Number (SSN)
// It returns a slice containing filtered Patient entries: []Patient
func GetPatientBySSN(collection *mongo.Collection, patientSSN string) (db.MetaData, error) {

	cursor, err := collection.Find(context.Background(), bson.D{{Key: "ssn", Value: patientSSN}})

	if err != nil {
		return nil, fmt.Errorf("error during document search: %v", err)
	}

	var results []db.Patient
	err = cursor.All(context.Background(), &results)

	if err != nil {
		return nil, fmt.Errorf("error during document decoding: %v", err)
	}

	return results, nil
}

// GetAllPatients retrieves all Patients from a given collection
// It returns a slice containing all Patient entries: []Patient
func GetAllPatients(collection *mongo.Collection) (db.MetaData, error) {

	cursor, err := collection.Find(context.Background(), bson.D{})

	if err != nil {
		return nil, fmt.Errorf("error during document search: %v", err)
	}

	var results []db.Patient
	err = cursor.All(context.Background(), &results)

	if err != nil {
		return nil, fmt.Errorf("error during document decoding: %v", err)
	}

	return results, nil
}

// GetAllPatients retrieves all Patients from a given collection
// It returns a slice containing all Patient entries: []Patient
func GetAllElectroCardiogramDevices(collection *mongo.Collection) (db.MetaData, error) {

	cursor, err := collection.Find(context.Background(), bson.D{})

	if err != nil {
		return nil, fmt.Errorf("error during document search: %v", err)
	}

	var results []db.ElectroCardiogramDevice
	err = cursor.All(context.Background(), &results)

	if err != nil {
		return nil, fmt.Errorf("error during document decoding: %v", err)
	}

	return results, nil
}
