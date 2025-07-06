package config

import (
	"fmt"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	ce "server/internal/core/error"
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

func NewConfig(configPath string) (Config, ce.CoreError) {
	logger := zap.L()
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(configPath)
	logger.Info(fmt.Sprintf("find config.toml in path %v", configPath))

	var cfg Config

	if viperErr := viper.ReadInConfig(); viperErr != nil {
		err := ce.New("read config from file", ce.WithErr(viperErr))
		err.AddOp("create config")
		return cfg, err
	}

	if viperErr := viper.Unmarshal(&cfg); viperErr != nil {
		err := ce.New("unmarshal config file", ce.WithErr(viperErr))
		err.AddOp("create config")
		return cfg, err
	}

	return cfg, nil
}
