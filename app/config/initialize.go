package config

import (
	"os"

	"github.com/joho/godotenv"
)

func LoadConfigurations() (Config, error) {
	err := godotenv.Load(".env")
	if err != nil {
		return Config{}, err
	}

	return Config{
		Database: DatabaseConfigurations{
			User:           os.Getenv("DB_USER"),
			Password:       os.Getenv("DB_PASS"),
			Host:           os.Getenv("DB_HOST"),
			Name:           os.Getenv("DB_NAME"),
			SSLMode:        os.Getenv("DB_SSL"),
			MigrationsPath: os.Getenv("DB_MIG"),
		},
	}, nil
}
