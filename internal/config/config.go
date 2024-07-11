package config

import (
	"log"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	Server struct {
		Port string
		Host string
	}
	Postgres struct {
		Host     string
		Port     string
		Username string
		Password string
		DBName   string
		Pool     struct {
			Max uint
		}
	}
	SecretKey string
	JWT       struct {
		HoursToExpire int
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
