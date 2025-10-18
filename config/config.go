package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Database databaseConfig
	Server   server
}

type databaseConfig struct {
	User     string
	Password string
	DbName   string
	PortDb   string
	Host     string
}

type server struct {
	Port string
}

func LoadConfig() (Config, error) {
	if err := godotenv.Load(); err != nil {
		return Config{}, fmt.Errorf("error loading .env file: %w", err)
	}

	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	portDb := os.Getenv("DB_PORT")
	host := os.Getenv("DB_HOST")

	appPort := os.Getenv("APP_PORT")

	cfg := Config{
		Database: databaseConfig{
			User:     user,
			Password: password,
			DbName:   dbName,
			PortDb:   portDb,
			Host:     host,
		},
		Server: server{
			Port: appPort,
		},
	}
	return cfg, nil
}
