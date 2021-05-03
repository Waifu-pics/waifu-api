package config

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type Config struct {
	Port        string
	Web         Web
	DatabaseUrl string
	Storage     Storage
	Endpoints   Endpoints
	Domain      string
	Frontend    string
}

type Web struct {
	Cdn string
	Jwt string
}

type Storage struct {
	Endpoint  string
	Accesskey string
	Secretkey string
	Region    string
	Bucket    string
}

type Endpoints struct {
	Sfw  []string `json:"sfw"`
	Nsfw []string `json:"nsfw"`
}

// LoadConfig : Load config from env
func LoadConfig() Config {
	return Config{
		Port:        getEnv("PORT"),
		DatabaseUrl: getEnv("DATABASE_URL"),
		Web: Web{
			Cdn: getEnv("CDN_URL"),
			Jwt: getEnv("JWT_KEY"),
		},
		Storage: Storage{
			Endpoint:  getEnv("S3_ENDPOINT"),
			Accesskey: getEnv("S3_ACCESS_KEY"),
			Secretkey: getEnv("S3_SECRET_KEY"),
			Region:    getEnv("S3_REGION"),
			Bucket:    getEnv("S3_BUCKET"),
		},
		Endpoints: Endpoints{
			Sfw:  strings.Split(getEnv("ENDPOINTS_SFW"), ","),
			Nsfw: strings.Split(getEnv("ENDPOINTS_NSFW"), ","),
		},
		Domain:   getEnv("DOMAIN"),
		Frontend: getEnv("FRONTEND_URL"),
	}
}

func getEnv(key string) string {
	value, set := os.LookupEnv(key)
	if !set {
		log.Fatalln(fmt.Sprintf("Config variable %s was missing", key))
	}
	return value
}
