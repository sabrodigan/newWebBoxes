package config

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
)

var client *mongo.Client

func init() {
	var err error
	client, err = mongo.Connect(context.TODO(), options.Client().ApplyURI(os.Getenv("MONGODB_URI")))
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}
}

func GetDatabaseCollection(dbName, collectionName string) *mongo.Collection {
	if client == nil {
		log.Fatalf("MongoDB client is not initialized")
	}
	db := client.Database(dbName)
	return db.Collection(collectionName)
}
