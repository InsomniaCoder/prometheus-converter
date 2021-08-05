package config

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type AppConfig struct {
	Server ServerConfig
	Debug  bool
	Target TargetConfig
}

type ServerConfig struct {
	Port int
}

type TargetConfig struct {
	URL string
}

var Config AppConfig

func init() {
	log.Println("initializaing config")
	LoadConfiguration(&Config)
}

func LoadConfiguration(appConfig *AppConfig) {
	viper.SetConfigFile(`config.yaml`)
	err := viper.ReadInConfig()

	// Read Environment Variables
	viper.AutomaticEnv()
	if err != nil {
		log.Fatalf("Cannot read configuration, %v \n", err)
		panic(err)
	}

	marshalErr := viper.Unmarshal(appConfig)

	if marshalErr != nil {
		log.Fatalf("Unable to decode into struct, %v \n", marshalErr)
		panic(marshalErr)
	}

	log.Printf("loaded configuration %+v", appConfig)
}
