package main

import (
	"fmt"
	"os"

	c "github.com/VladNF/otus-golang/hw12_13_14_15_calendar/internal/common"
	"github.com/spf13/viper"
)

type Config struct {
	HTTP    c.HTTPConf    `mapstructure:"http"`
	GRPC    c.GRPCConf    `mapstructure:"grpc"`
	Logger  c.LoggerConf  `mapstructure:"logger"`
	Storage c.StorageConf `mapstructure:"storage"`
}

func NewConfig(file string) Config {
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
		config.HTTP = c.HTTPConfFromEnv()
		config.GRPC = c.GRPCConfFromEnv()
		config.Logger = c.LoggerConfFromEnv()
		config.Storage = c.StorageConfFromEnv()
	}
	fmt.Fprintf(os.Stderr, "Loaded config %v\n", *config)
	return *config
}
