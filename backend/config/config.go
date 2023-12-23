package config

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
)

type Config struct {
	DBName       string `env:"DB_NAME"`
	Host         string `env:"DB_HOST"`
	Password     string `env:"DB_PASSWORD"`
	Port         string `env:"DB_PORT"`
	User         string `env:"DB_USER"`
	S3Region     string `env:"S3_REGION"`
	S3AccesKeyID string `env:"AWS_ACCESS_KEY_ID"`
	S3SecretKey  string `env:"AWS_SECRET_KEY"`
	S3BucketName string `env:"S3_BUCKET_NAME"`
	Environment  string `env:"ENVIRONMENT"`
}

// Load the config from the environment variables
func LoadConfig() (*Config, error) {
	config := Config{}
	err := LoadFromFile(&config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}

// Load the config from a file
func LoadFromEnv(config *Config) error {
	config.DBName = os.Getenv("DB_NAME")
	config.Host = os.Getenv("DB_HOST")
	config.Password = os.Getenv("DB_PASSWORD")
	config.Port = os.Getenv("DB_PORT")
	config.User = os.Getenv("DB_USER")
	config.S3Region = os.Getenv("S3_REGION")
	config.S3AccesKeyID = os.Getenv("AWS_ACCESS_KEY_ID")
	config.S3SecretKey = os.Getenv("AWS_SECRET_KEY")
	config.S3BucketName = os.Getenv("S3_BUCKET_NAME")
	config.Environment = os.Getenv("ENVIRONMENT")

	return ValidateConfig(config)
}

// Load the config from a file
func LoadFromFile(config *Config) error {
	err := godotenv.Load("config/config.env")
	if err != nil {
		log.Warn().Err(err).Msg("Error loading .env file. Using default config.")
	}

	return LoadFromEnv(config)
}

// Validate the config
func ValidateConfig(config *Config) error {
	if config.DBName == "" {
		return errors.New("DB_NAME is not set")
	}
	if config.Host == "" {
		return errors.New("DB_HOST is not set")
	}
	if config.Password == "" {
		return errors.New("DB_PASSWORD is not set")
	}
	if config.Port == "" {
		return errors.New("DB_PORT is not set")
	}
	if config.User == "" {
		return errors.New("DB_USER is not set")
	}
	if config.S3Region == "" {
		return errors.New("S3_REGION is not set")
	}
	if config.S3AccesKeyID == "" {
		return errors.New("AWS_ACCESS_KEY_ID is not set")
	}
	if config.S3SecretKey == "" {
		return errors.New("AWS_SECRET_KEY is not set")
	}
	if config.S3BucketName == "" {
		return errors.New("S3_BUCKET_NAME is not set")
	}
	if config.Environment == "" {
		return errors.New("ENVIRONMENT is not set")
	}

	return nil
}
