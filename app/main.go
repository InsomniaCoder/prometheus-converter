package main

import (
	"github.com/insomniacoder/prometheus-converter/config"
	log "github.com/sirupsen/logrus"

	_ "github.com/insomniacoder/prometheus-converter/api"
)

func init() {
	log.Println("initializaing main")
	log.Printf("loaded config: %+v \n", config.Config)
	if config.Config.Debug {
		log.Println("Running in DEBUG mode")
	}
}

func main() {
}
