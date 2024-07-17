package config

import (
	"fmt"
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
)

type Config struct {
	ConsulConfig  ConsulConfig
	ServiceConfig ServiceConfig
}

type ConsulConfig struct {
	Address        string `validate:"required,url"`
	DeregisterTime string `validate:"required"`
	IntervalTime   string `validate:"required"`
}

type ServiceConfig struct {
	UserServiceName string `validate:"required"`
}

func LoadConfig() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		fmt.Println("No .env file found. Using environment variables.")
	} else {
		fmt.Println("Loaded .env file successfully.")
	}

	config := &Config{
		ConsulConfig: ConsulConfig{
			Address:        getEnv("CONSUL_ADDR", "http://localhost:8500"),
			DeregisterTime: getEnv("CONSUL_DEREGISTER_TIME", "10m"),
			IntervalTime:   getEnv("CONSUL_INTERVAL_TIME", "5m"),
		},
		ServiceConfig: ServiceConfig{
			UserServiceName: getEnv("USER_SERVICE_NAME", "User_Service"),
		},
	}

	v := validator.New()

	if err := v.Struct(config); err != nil {
		return nil, fmt.Errorf("config validation failed: %v", err)
	}

	return config, nil
}

func getEnv(key string, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultValue
}
