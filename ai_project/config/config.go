package config

import (
	"errors"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	MongoURI     string
	DatabaseName string
}

// LoadConfig loads environment variables from a .env file (if present)
// and returns the configuration struct.
func LoadConfig() (*Config, error) {
	// Load .env file silently (no error if not found).
	_ = godotenv.Load()

	mongoURI := os.Getenv("MONGODB_URI")
	dbName := os.Getenv("DATABASE_NAME")

	if mongoURI == "" {
		return nil, errors.New("MONGODB_URI is required but not set")
	}
	if dbName == "" {
		return nil, errors.New("DATABASE_NAME is required but not set")
	}

	cfg := &Config{
		MongoURI:     mongoURI,
		DatabaseName: dbName,
	}

	log.Printf("Configuration loaded: %+v\n", cfg)
	return cfg, nil
}
