package api

import (
	s3simple "github.com/Riku32/s3-simple"
	"github.com/Riku32/waifu.pics/src/config"
	"github.com/Riku32/waifu.pics/src/database"
)

// Options : route object
type Options struct {
	Database database.Database
	Config   config.Config
	S3       *s3simple.Session
}
