package util

import (
	"encoding/json"
	"fmt"
	"os"
)

// Config : Template for JSON Config
type Config struct {
	DB   string `json:"DB"`
	URL  string `json:"URL"`
	PORT string `json:"PORT"`
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
