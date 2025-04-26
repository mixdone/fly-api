package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	DB struct {
		Host    string `mapstructure:"host"`
		Port    string `mapstructure:"port"`
		User    string `mapstructure:"user"`
		DBName  string `mapstructure:"dbname"`
		SSLMode string `mapstructure:"sslmode"`
	} `mapstructure:"db"`

	Log struct {
		Level  string `mapstructure:"level"`
		Format string `mapstructure:"format"`
	} `mapstructure:"log"`

	Server struct {
		Host string `mapstructure:"host"`
		Port int    `mapstructure:"port"`
	} `mapstructure:"server"`
}

func LoadConfig() (*Config, error) {
	viper.SetConfigName("config")
	viper.AddConfigPath("./configs")
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("error reading config file, %s", err)
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, fmt.Errorf("unable to unmarshal config into struct, %s", err)
	}

	return &config, nil
}
