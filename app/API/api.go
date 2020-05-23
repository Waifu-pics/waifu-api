package api

import (
	"go.mongodb.org/mongo-driver/mongo"
	"waifu.pics/util/config"
)

type API struct {
	Config   config.Config
	Database *mongo.Database
}
