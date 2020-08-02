package api

import (
	"waifu.pics/util/config"
	"waifu.pics/util/database"
)

type API struct {
	Config   config.Config
	Database database.Database
}
