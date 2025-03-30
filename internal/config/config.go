package config

import (
	"fmt"
	"github.com/spf13/viper"
	"time"
)

type Config struct {
	Environment string `yaml:"environment"`

	Servers struct {
		HTTP struct {
			Host    string        `yaml:"host"`
			Port    int           `yaml:"port"`
			Timeout time.Duration `yaml:"timeout"`
		} `yaml:"http"`
	} `yaml:"servers"`

	Connections struct {
		Postgres struct {
			Username   string `yaml:"username"`
			Password   string `yaml:"password"`
			Host       string `yaml:"host"`
			Port       int    `yaml:"port"`
			DBName     string `yaml:"dbName"`
			SSLEnabled string `yaml:"sslEnabled"`
		} `yaml:"postgres"`
		Redis struct {
			Host string `yaml:"host"`
			Port int    `yaml:"port"`
		} `yaml:"redis"`
		Services struct {
			Users   string        `yaml:"users"`
			TimeOut time.Duration `yaml:"timeout"`
		} `yaml:"services"`
	} `yaml:"connections"`
	Secrets struct {
		Public string `yaml:"public"`
	} `yaml:"secrets"`
}

func LoadConfig(configPath string) (*Config, error) {
	viper.SetConfigFile(configPath)
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config file: %w", err)
	}

	return &config, nil
}
