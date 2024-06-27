package config

import (
	"github.com/spf13/viper"
	"log"
)

var config *viper.Viper

func Init(env string) {
	var err error
	log.Print("This is the environment: ", env)
	config = viper.New()
	config.SetConfigType("yaml")
	config.SetConfigName(env)
	config.AddConfigPath("./config/")
	err = config.ReadInConfig()
	if err != nil {
		log.Fatalf("Error on parsing configuration file: %s", err)
	}
}

func GetConfig() *viper.Viper {
	return config
}
