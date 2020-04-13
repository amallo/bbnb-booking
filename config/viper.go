package config

import (
	"log"

	"github.com/spf13/viper"
)

func WithViperConfig(envFile string) ReadConfigKeyFunc {
	viper.SetConfigFile(envFile)
	viper.SetConfigType("env")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error while reading config file %s", err)
	}
	return func(key string) string {
		value, ok := viper.Get(key).(string)
		if !ok {
			log.Fatalf("Invalid type assertion")
		}
		return value
	}
}

type ReadConfigKeyFunc = func(key string) string
