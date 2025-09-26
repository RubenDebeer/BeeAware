package config

import (
	"os"
)

type Config struct {
	MongoURI string
	DBName   string
}

func Load() Config {
	cfg := Config{
		MongoURI: getenv("MONGO_URI", "mongodb://localhost:27017"),
		DBName:   getenv("MONGO_DB", "beeaware_dev"),
	}
	return cfg
}

func getenv(k, def string) string {
	if v := os.Getenv(k); v != "" {
		return v
	}
	return def
}
