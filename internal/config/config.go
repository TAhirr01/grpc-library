package config

import (
	"log"
	"os"
)

type Config struct {
	Database DatabaseConfig
	Server   ServerConfig
}

type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	DSN      string
}

type ServerConfig struct {
	Port string
}

func Load() *Config {
	return &Config{
		Database: DatabaseConfig{
			Host:     getEnv("DB_HOST", "127.0.0.1"),
			Port:     getEnv("DB_PORT", "3306"),
			User:     getEnv("DB_USER", "root"),
			Password: getEnv("DB_PASSWORD", "root"),
			DBName:   getEnv("DB_NAME", "grpc-library"),
			DSN:      getEnv("DB_DSN", "root:root@tcp(127.0.0.1:3306)/grpc-library?charset=utf8mb4&parseTime=True&loc=Local"),
		},
		Server: ServerConfig{
			Port: getEnv("SERVER_PORT", "50051"),
		},
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}