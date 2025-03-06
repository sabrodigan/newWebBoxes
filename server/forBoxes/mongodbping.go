package forBoxes

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Ping() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	// this is to ensure that the login is not exposed
	// Get database URL from environment variables
	mongoURI := os.Getenv("MONGODB_URI")
	if mongoURI == "" {
		log.Fatalf("MONGODB_URI is not set in the environment variables")
	}

	// Set client options
	clientOptions := options.Client().ApplyURI(mongoURI)

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	startTime := time.Now()
	err = client.Ping(context.TODO(), nil)
	elapsedTime := time.Since(startTime)

	if err != nil {
		log.Fatal("Failed to ping MongoDB:", err)
	}

	fmt.Printf("Connected to MongoDB! Ping took %.4f seconds\n", elapsedTime.Seconds())
}
