package main

import (
	"fmt"
	"os"

	c "github.com/VladNF/otus-golang/hw12_13_14_15_calendar/internal/common"
	"github.com/spf13/viper"
)

type Config struct {
	Logger c.LoggerConf `mapstructure:"logger"`
	MQ     c.MQConf     `mapstructure:"mq"`
}

func NewConfigFromFile(file string) Config {
	config := &Config{}
	viper.SetConfigType("yaml")
	viper.SetConfigFile(file)
	if err := viper.ReadInConfig(); err == nil {
		if err = viper.Unmarshal(config); err != nil {
			fmt.Fprintln(os.Stderr, fmt.Errorf("config file: %w", err))
		}
	} else {
		fmt.Fprintln(os.Stderr, fmt.Errorf("config file: %w", err))
		fmt.Fprintln(os.Stderr, "trying to read from OS env vars")
		config.MQ = c.MQConfFromEnv()
		config.Logger = c.LoggerConfFromEnv()
	}
	return *config
}
