package api

import (
	s3simple "github.com/Riku32/s3-simple"
	"github.com/Waifu-pics/waifu-api/config"
	"github.com/Waifu-pics/waifu-api/database"
)

// Options : route object
type Options struct {
	Database database.Database
	Config   config.Config
	S3       *s3simple.Session
}
