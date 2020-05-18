package api

import (
	"go.mongodb.org/mongo-driver/mongo"
	"waifu.pics/util"
)

type API struct {
	Config util.Config
	Database *mongo.Database
}