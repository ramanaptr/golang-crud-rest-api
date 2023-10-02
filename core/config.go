package core

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Domain           string `mapstructure:"domain"`
	Port             string `mapstructure:"port"`
	ConnectionString string `mapstructure:"connection_string"`
	Secret           string `mapstructure:"secret"`
}

var AppConfig *Config

func LoadAppConfig() {
	log.Println("Loading Server Configurations...")
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.SetConfigType("json")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}
	err = viper.Unmarshal(&AppConfig)
	if err != nil {
		log.Fatal(err)
	}
}
