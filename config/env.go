package config

import "os"

type Config struct {
	DatabaseURI string
}

func GetEnvConfig() Config {

	conf := Config{
		DatabaseURI: os.Getenv("MONGO_URL"),
	}

	return conf
}
