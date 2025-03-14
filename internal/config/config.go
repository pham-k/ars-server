package config

import (
	"fmt"
	"github.com/spf13/viper"
	"server/internal/log"
)

type Config struct {
	HTTP struct {
		Host   string
		Port   int
		Domain string
	}
	RDB struct {
		Host                      string
		Port                      int
		Username                  string
		Password                  string
		Database                  string
		MaxOpenConnections        int
		MaxIdleConnections        int
		MaxIdleConnectionLifetime string
	}
	IDB struct {
		Host     string
		Port     int
		Username string
		Password string
		Database int
	}
	SMTP struct {
		Host     string
		Port     int
		Username string
		Password string
		Sender   string
	}
}

func NewConfig(configPath string, log log.Log) (Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(configPath)
	log.Info(fmt.Sprintf("Looking for config.toml from %v", configPath))

	var cfg Config

	if err := viper.ReadInConfig(); err != nil {
		log.Error("Fail to read config from file", "error", err)
		return cfg, err
	}

	if err := viper.Unmarshal(&cfg); err != nil {
		log.Error("Fail to unmarshal file", "error", err)
		return cfg, err
	}

	return cfg, nil
}
