package config

import (
	"os"
)

type Config struct {
	Password string
}

func LoadConfig() *Config {
	return &Config{
		Password: os.Getenv("TODO_PASSWORD"),
	}
}