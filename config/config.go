package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	Storage StorageConfig `json:"storage"`
}

type StorageConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	Database string
}

func GetConfig() *Config {
	getEnv()
	instance := &Config{}
	instance.Storage = StorageConfig{
		Username: os.Getenv("POSTGRES_USER"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		Host:     os.Getenv("POSTGRES_HOST"),
		Port:     os.Getenv("POSTGRES_PORT"),
		Database: os.Getenv("POSTGRES_DB"),
	}
	return instance
}

func getEnv() {
	err := godotenv.Load("../.env.docker")
	if err != nil {
		log.Fatal("Error loading .env file")
		return
	}
}
