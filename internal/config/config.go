package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DATABASE_URL string
}

func Load() Config {
	_ = godotenv.Load()
	cfg := Config{
		DATABASE_URL: getEnv("DATABASE_URL"),
	}

	return cfg
}

func getEnv(key string) string {
	val, ok := os.LookupEnv(key)
	if !ok {
		log.Fatalf("cant find env key: %s", key)
	}

	return val
}
