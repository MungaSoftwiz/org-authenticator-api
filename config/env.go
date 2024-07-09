package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type PostgreSQLConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

var Env = initConfig()

func initConfig() PostgreSQLConfig {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file:", err)
	}
	return envFetch()
}

func envFetch() PostgreSQLConfig {
	return PostgreSQLConfig{
		Host:     getEnv("PGHOST", "localhost"),
		Port:     getEnv("PGPORT", strconv.Itoa(5432)),
		User:     getEnv("PGUSER", "root"),
		Password: getEnv("PGPASSWORD", "password"),
		DBName:   getEnv("PGDATABASE", "org-db"),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok && value != "" {
		return value
	}
	return fallback
}
