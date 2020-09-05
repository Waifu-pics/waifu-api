package api

import (
	"github.com/Riku32/waifu.pics/src/util/config"
	"github.com/Riku32/waifu.pics/src/util/database"
)

type API struct {
	Config   config.Config
	Database database.Database
}
