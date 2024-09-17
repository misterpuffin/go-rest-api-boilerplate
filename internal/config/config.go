package config

import (
	"log"
	"os"
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
	SecretKey string `mapstructure:"secret_key"`
	JWT       struct {
		HoursToExpire int `mapstructure:"hours_to_expire"`
	} `mapstructure:"jwt"`
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

var testConfig *Config

func LoadTestConfig() *Config {
	if testConfig != nil {
		return testConfig
	}

	var err error
	path := "./config/"
	for i := 0; i < 10; i++ {

		viper.AddConfigPath(path)
		viper.SetConfigName("test")
		viper.SetConfigType("yaml")

		viper.SetEnvKeyReplacer(strings.NewReplacer(`.`, `_`))
		viper.AutomaticEnv()

		err = viper.ReadInConfig()
		if err != nil {
			if strings.Contains(err.Error(), "Not Found") {
				path = "../" + path
				continue
			}
		} else {
			break
		}
	}
	var config Config

	err = viper.Unmarshal(&config)
	if err != nil {
		log.Fatalf("Could not read config: %v", err)
		os.Exit(1)
	}
	testConfig = &config
	return testConfig
}
