package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type config struct {
	Http struct {
		ListenAddress string `yaml:"listen_address"`
	} `yaml:"http"`
	MySQL struct {
		DataSourceName string `yaml:"data_source"`
	} `yaml:"mysql"`
}

var cfg = config{}

func init() {

	data, err := ioutil.ReadFile("config/config.yaml")
	if err != nil {
		log.Fatalln(err)
	}

	err = yaml.Unmarshal(data, &cfg)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	log.Println("[config]", "config loaded", cfg)
}

func Config() *config {
	return &cfg
}
