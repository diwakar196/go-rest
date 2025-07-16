package config

import (
	"os"
)

type Config struct {
	Port   string
	DBHost string
	DBPass string
	DBUser string
	DBName string
}

func LoadConfig() *Config {

	return &Config{
		Port:   os.Getenv("PORT"),
		DBHost: os.Getenv("DB_HOST"),
		DBPass: os.Getenv("DB_PASSWORD"),
		DBUser: os.Getenv("DB_USER"),
		DBName: os.Getenv("DB_NAME"),
	}
}
