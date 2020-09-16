package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	s3simple "github.com/Riku32/s3-simple"
	"github.com/Riku32/waifu.pics/src/api"
	"github.com/Riku32/waifu.pics/src/api/router"
	"github.com/Riku32/waifu.pics/src/cmd/admin"
	"github.com/Riku32/waifu.pics/src/config"
	"github.com/Riku32/waifu.pics/src/database"
	"github.com/Riku32/waifu.pics/src/static"
)

func main() {
	dev := flag.Bool("dev", false, "dev mode")
	newuser := flag.Bool("newuser", false, "create an administrator")
	flag.Parse()
	static.Dev = *dev

	mydir, _ := os.Getwd()
	fmt.Println(mydir)
	// Different structure for different modes
	var cfg config.Config
	if static.Dev {
		cfg = config.LoadConfig("../config.yml")
	} else {
		// Docker volume in folder
		cfg = config.LoadConfig("./config/config.yml")
	}

	db := database.InitSQL(cfg)

	// Admin creation argument
	if *newuser {
		admin.CreateAdmin(db)
		return
	}

	// Initiate S3
	s3, err := s3simple.New(s3simple.Config{
		Region:   cfg.Storage.Region,
		Endpoint: cfg.Storage.Endpoint,
		Bucket:   cfg.Storage.Bucket,
		Credentials: s3simple.Credentials{
			Accesskey: cfg.Storage.Accesskey,
			Secretkey: cfg.Storage.Secretkey,
		},
	})

	if err != nil {
		log.Fatalln("Unable to start S3")
	}

	options := api.Options{
		Database: db,
		Config:   cfg,
		S3:       s3,
	}

	router.New(options)

	// Do not close program
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
}
