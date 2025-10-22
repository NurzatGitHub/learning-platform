package config

import (
	"os"
	"strconv"
	"time"
)

type Config struct {
	JWTSecret       string
	HTTPPort        string
	HTTPReadTimeout time.Duration
	// ... другие поля (DB DSN, Redis URL и т.д.)
}

func MustLoad() *Config {
	port := getEnv("HTTP_PORT", "8080")
	readTimeoutStr := getEnv("HTTP_READ_TIMEOUT", "10")

	readTimeout, err := strconv.Atoi(readTimeoutStr)
	if err != nil {
		panic("HTTP_READ_TIMEOUT must be a number")
	}

	return &Config{
		JWTSecret:       getEnv("JWT_SECRET", "your-default-secret"),
		HTTPPort:        port,
		HTTPReadTimeout: time.Duration(readTimeout) * time.Second,
		// ... инициализация других полей
	}
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}