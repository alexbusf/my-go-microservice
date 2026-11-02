package config

import (
	"os"
)

// AppConfig хранит настройки приложения
type AppConfig struct {
	Port     string
	DbConfig DbConfig
}

// DbConfig хранит настройки базы данных
type DbConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
}

// InitConfig читает переменные и возвращает готовую структуру
func InitConfig() AppConfig {
	return AppConfig{
		Port: getEnv("PORT", "8080"),
		DbConfig: DbConfig{
			// Для Docker используем имя сервиса 'db', а не localhost
			Host:     getEnv("DB_HOST", "db"), 
			Port:     getEnv("DB_PORT", "5432"),
			// Логин и база из нашего docker-compose
			User:     getEnv("DB_USER", "gouser"), 
			Password: getEnv("DB_PASSWORD", "secret"),
			Name:     getEnv("DB_NAME", "mydb"),
		},
	}
}

// Вспомогательная функция: берет переменную или значение по умолчанию
func getEnv(key string, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}