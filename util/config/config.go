package config

import (
	"encoding/json"
	"fmt"
	"os"
)

// Config : Template for JSON Config
type Config struct {
	URL  string `json:"URL"`
	PORT string `json:"PORT"`
	DB   struct {
		URL    string `json:"URL"`
		DBNAME string `json:"DBNAME"`
	} `json:"DB"`
	S3 struct {
		ENDPOINT  string `json:"ENDPOINT"`
		ACCESSKEY string `json:"ACCESSKEY"`
		SECRETKEY string `json:"SECRETKEY"`
		REGION    string `json:"REGION"`
		BUCKET    string `json:"BUCKET"`
	} `json:"S3"`
	ENDPOINTS []string `json:"ENDPOINTS"`
}

// LoadConfig : Load config from external json
func LoadConfig(file string) Config {
	var config Config

	if !fileExists(file) {
		fmt.Println("CONFIG NOT FOUND!")
		os.Exit(0)
	}

	configFile, err := os.Open(file)
	defer configFile.Close()
	if err != nil {
		fmt.Println(err.Error())
	}
	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&config)

	return config
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
