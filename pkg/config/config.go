package config

import (
	"os"
	"sync"
)

type Config struct {
	Port   string
	DBHost string
	DBPort string
	DBPass string
	DBUser string
	DBName string
}

var (
	config *Config
	once   sync.Once
)

func loadConfig() {
	config = &Config{
		Port:   os.Getenv("PORT"),
		DBHost: os.Getenv("DB_HOST"),
		DBPort: os.Getenv("DB_PORT"),
		DBPass: os.Getenv("DB_PASSWORD"),
		DBUser: os.Getenv("DB_USER"),
		DBName: os.Getenv("DB_NAME"),
	}
}

func GetConfig() *Config {
	if config == nil {
		once.Do(loadConfig)
	}
	return config
}
