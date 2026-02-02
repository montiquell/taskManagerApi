package config

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	DATABASE_URL string

	ADDR         string
	READTIMEOUT  time.Duration
	WRITETIMEOUT time.Duration
	IDLETIMEOUT  time.Duration
}

func Load() Config {
	_ = godotenv.Load()

	cfg := Config{
		DATABASE_URL: getEnvString("DATABASE_URL"),

		ADDR:         getEnvString("ADDR"),
		READTIMEOUT:  time.Duration(getEnvInt("READTIMEOUT")) * time.Second,
		WRITETIMEOUT: time.Duration(getEnvInt("WRITETIMEOUT")) * time.Second,
		IDLETIMEOUT:  time.Duration(getEnvInt("IDLETIMEOUT")) * time.Second,
	}

	return cfg
}

func getEnvString(key string) string {
	val, ok := os.LookupEnv(key)
	if !ok {
		log.Fatalf("cant find env key: %s", key)
	}

	return val
}

func getEnvInt(key string) int {
	val := getEnvString(key)
	i, err := strconv.Atoi(val)
	if err != nil {
		log.Fatalf("cant convert env key to int: %s", key)
	}

	return i
}
