package repository

import (
	"github.com/sabrodigan/newWebBoxes/server/config"
	"log"
)

type Repository struct {
	UserRepository IMongoRepository
}

// GetRepository initializes the Repository
func GetRepository() *Repository {

	dbName, err := config.GetEnvProperty("databaseName")
	if err != nil {
		log.Fatalf("Error fetching database name: %v", err)
	}

	users, err := config.GetEnvProperty("users")
	if err != nil {
		log.Fatalf("ALARM! fetching users collection: %v", err)
	}

	if dbName == "" || users == "" {
		log.Fatalf("ALARM! database name or users collection is empty")
	}

	// Pass the dbName and collection to the GetMongoRepository function
	return &Repository{
		UserRepository: GetMongoRepository(dbName, users),
	}
}
