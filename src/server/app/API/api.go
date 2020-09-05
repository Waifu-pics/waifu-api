package api

import (
	"github.com/Riku32/waifu.pics/util/config"
	"github.com/Riku32/waifu.pics/util/database"
)

type API struct {
	Config   config.Config
	Database database.Database
}
