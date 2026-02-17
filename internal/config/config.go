package config

import (
	"os"
	"strconv"
	"time"
)

// Config holds application configuration
type Config struct {
	ServerPort       string
	JWTSecret        string
	JWTTokenDuration time.Duration
}

// Load loads configuration from environment variables with defaults
func Load() *Config {
	port := getEnv("PORT", "8080")
	jwtSecret := getEnv("JWT_SECRET", "your-secret-key-change-in-production")
	jwtDuration := getEnvAsDuration("JWT_DURATION", 24*time.Hour)

	return &Config{
		ServerPort:       port,
		JWTSecret:        jwtSecret,
		JWTTokenDuration: jwtDuration,
	}
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

func getEnvAsDuration(key string, defaultValue time.Duration) time.Duration {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	hours, err := strconv.Atoi(value)
	if err != nil {
		return defaultValue
	}
	return time.Duration(hours) * time.Hour
}
