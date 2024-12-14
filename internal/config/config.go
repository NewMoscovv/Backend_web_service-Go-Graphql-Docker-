package config

import (
	"os"
)

type Config struct {
	DBHost     string
	DBUser     string
	DBPassword string
	DBName     string
	ServerPort string
}

// LoadConfig считывает конфигурацию из переменных окружения
func LoadConfig() *Config {
	return &Config{
		DBHost:     getEnv("DB_HOST", "localhost"),
		DBUser:     getEnv("DB_USER", "postgress-user"),
		DBPassword: getEnv("DB_PASSWORD", "postgress-password"),
		DBName:     getEnv("DB_NAME", "pg_db"),
		ServerPort: getEnv("PORT", "8080"),
	}
}

// getEnv возвращает значение переменной окружения или значение по умолчанию
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
