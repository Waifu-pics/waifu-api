package util

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Mongo struct {
	Database *mongo.Database
	Context  context.Context
}

// DbDriver : Mongo Driver
func DbDriver(config Config) Mongo {
	client, err := mongo.NewClient(options.Client().ApplyURI(config.DB.URL))
	if err != nil {
		log.Fatal(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)

	if err != nil {
		log.Fatal(err)
	}

	database := client.Database(config.DB.DBNAME)

	return Mongo{database, ctx}
}
