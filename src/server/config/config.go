package config

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

// Config : Template for Yaml Config
type Config struct {
	Port string `yaml:"port"`
	Web  struct {
		Cdn string `yaml:"cdn"`
		Jwt string `yaml:"jwt"`
	} `yaml:"web"`
	Database struct {
		URL string `yaml:"url"`
	} `yaml:"database"`
	Storage struct {
		Endpoint  string `yaml:"endpoint"`
		Accesskey string `yaml:"accesskey"`
		Secretkey string `yaml:"secretkey"`
		Region    string `yaml:"region"`
		Bucket    string `yaml:"bucket"`
	} `yaml:"storage"`
	Endpoints struct {
		Sfw  []string `yaml:"sfw" json:"sfw"`
		Nsfw []string `yaml:"nsfw" json:"nsfw"`
	} `yaml:"endpoints"`
}

// LoadConfig : Load config from yaml
func LoadConfig(file string) Config {
	var config Config

	if !fileExists(file) {
		log.Println("Config not found!")
		os.Exit(0)
	}

	configFile, err := os.Open(file)
	defer configFile.Close()
	if err != nil {
		fmt.Println(err.Error())
	}
	err = yaml.NewDecoder(configFile).Decode(&config)
	if err != nil {
		log.Println("Invalid config")
		os.Exit(0)
	}

	return config
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
