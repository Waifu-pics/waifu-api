package database

import (
	"context"
	"log"
	"time"
	"waifu.pics/util"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// InitDB : Mongo Initializer
func InitDB(config util.Config) *mongo.Database {
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

	return client.Database(config.DB.DBNAME)
}
