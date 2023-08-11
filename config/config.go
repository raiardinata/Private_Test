package config

import (
	"github.com/tkanos/gonfig"
)

// Define database connection parameters
type Configuration struct {
	DB_USERNAME string
	DB_PASSWORD string
	DB_PORT     string
	DB_HOST     string
	DB_NAME     string
}

func GetConfig() Configuration {
	conf := Configuration{}
	gonfig.GetConf("/home/raiardinata/ra_Server/go/ra-lab/Private_Test/config/config.json", &conf)

	return conf
}
