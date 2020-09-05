package api

import (
	"github.com/Riku32/waifu.pics/src/server/util/config"
	"github.com/Riku32/waifu.pics/src/server/util/database"
)

type API struct {
	Config   config.Config
	Database database.Database
}
