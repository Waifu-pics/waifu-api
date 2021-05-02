package main

import (
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"

	s3simple "github.com/Riku32/s3-simple"
	"github.com/Waifu-pics/waifu-api/api"
	"github.com/Waifu-pics/waifu-api/api/router"
	"github.com/Waifu-pics/waifu-api/cmd/admin"
	"github.com/Waifu-pics/waifu-api/config"
	"github.com/Waifu-pics/waifu-api/database"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	newuser := flag.Bool("newuser", false, "create an administrator")
	flag.Parse()

	conf := config.LoadConfig()
	db := database.InitSQL(conf)

	// Admin creation argument
	if *newuser {
		admin.CreateAdmin(db)
		return
	}

	// Initiate S3
	s3, err := s3simple.New(s3simple.Config{
		Region:   conf.Storage.Region,
		Endpoint: conf.Storage.Endpoint,
		Bucket:   conf.Storage.Bucket,
		Credentials: s3simple.Credentials{
			Accesskey: conf.Storage.Accesskey,
			Secretkey: conf.Storage.Secretkey,
		},
	})

	if err != nil {
		log.Fatalln("Unable to start S3")
	}

	options := api.Options{
		Database: db,
		Config:   conf,
		S3:       s3,
	}

	router.New(options)

	// Do not close program
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
}
