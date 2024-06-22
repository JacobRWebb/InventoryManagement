package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	ConsulAddr           string `mapstructure:"CONSUL_ADDR"`
	ConsulDeregisterTime string `mapstructure:"CONSUL_DEREGISTER_TIME"`
	ConsulIntervalTime   string `mapstructure:"CONSUL_INTERVAL_TIME"`
	ServiceName          string `mapstructure:"SERVICE_NAME"`
}

func NewConfig() (config *Config, err error) {
	viper.AutomaticEnv()

	viper.SetDefault("CONSUL_ADDR", "http://localhost:8500")
	viper.SetDefault("CONSUL_DEREGISTER_TIME", "1m")
	viper.SetDefault("CONSUL_INTERVAL_TIME", "1m")
	viper.SetDefault("SERVICE_NAME", "Inventory_Management")

	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	return config, nil
}
