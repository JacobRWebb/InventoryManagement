package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
)

type Config struct {
	ApplicationConfig ApplicationConfig
	CertsConfig       CertsConfig
	ConsulConfig      ConsulConfig
	ServiceConfig     ServiceConfig
}

type ApplicationConfig struct {
	Port int `validate:"required"`
}

type CertsConfig struct {
	CertificateLocation string `validate:"required"`
	KeyLocation         string `validate:"required"`
}

type ConsulConfig struct {
	Address string `validate:"required,url"`
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
		ApplicationConfig: ApplicationConfig{
			Port: getEnvAsInt("PORT", 3333),
		},
		CertsConfig: CertsConfig{
			CertificateLocation: getEnv("CERT_CERTIFICATE_LOCATION", "certs/pub.pem"),
			KeyLocation:         getEnv("CERT_KEY_LOCATION", "certs/key.pem"),
		},
		ConsulConfig: ConsulConfig{
			Address: getEnv("CONSUL_ADDR", "http://localhost:8500"),
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

func getEnvAsInt(key string, defaultValue int) int {
	valueStr := getEnv(key, "")
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}
	return defaultValue
}
