package util

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Database : global database instance for use around the app
var Database *mongo.Database

// InitDB : Mongo Initializer
func InitDB(config Config) {
	client, err := mongo.NewClient(options.Client().ApplyURI(config.DB.URL))
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)

	if err != nil {
		log.Fatal(err)
	}

	Database = client.Database(config.DB.DBNAME)
}
