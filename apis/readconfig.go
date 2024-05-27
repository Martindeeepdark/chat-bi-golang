package apis

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
)

type Config struct {
	API struct {
		AppID     string `yaml:"appid"`
		APIKey    string `yaml:"apiKey"`
		APISecret string `yaml:"apiSecret"`
		HostURL   string `yaml:"hostUrl"`
	} `yaml:"api"`
}

func ReadConfig() *Config {
	configFile, err := ioutil.ReadFile("configs/config.yml")
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	var config Config
	err = yaml.Unmarshal(configFile, &config)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	return &config
}
