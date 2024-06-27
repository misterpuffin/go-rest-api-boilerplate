package util

import (
	"log"
	"strings"

	"github.com/spf13/viper"
)

var config *viper.Viper

type Config struct {
	Server struct {
		Port uint
		Host string
	}
	Postgres struct {
		Host     string
		Port     uint
		Username string
		Password string
		Pool     struct {
			Max uint
		}
	}
}

func LoadConfig(env string) (config Config, err error) {
	log.Print("This is the environment: ", env)

	viper.SetConfigName(env)
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config/")

	viper.SetEnvKeyReplacer(strings.NewReplacer(`.`, `_`))
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
