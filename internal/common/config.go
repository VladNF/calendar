package common

import "github.com/spf13/viper"

type LoggerConf struct {
	Level string `mapstructure:"level"`
	File  string `mapstructure:"file"`
}

type StorageConf struct {
	Kind string `mapstructure:"kind"`
}

type HTTPConf struct {
	Host string `mapstructure:"host"`
	Port string `mapstructure:"port"`
}

type GRPCConf struct {
	Host string `mapstructure:"host"`
	Port string `mapstructure:"port"`
}

type MQConf struct {
	URI      string `mapstructure:"uri"`
	Exchange string `mapstructure:"exchange"`
	Key      string `mapstructure:"key"`
	Queue    string `mapstructure:"queue"`
}

func LoggerConfFromEnv() LoggerConf {
	viper.SetEnvPrefix("LOGGER")
	viper.AutomaticEnv()
	return LoggerConf{
		Level: viper.GetString("level"),
		File:  viper.GetString("file"),
	}
}

func StorageConfFromEnv() StorageConf {
	viper.SetEnvPrefix("STORAGE")
	viper.AutomaticEnv()
	return StorageConf{
		Kind: viper.GetString("kind"),
	}
}

func GRPCConfFromEnv() GRPCConf {
	viper.SetEnvPrefix("GRPC")
	viper.AutomaticEnv()
	return GRPCConf{
		Host: viper.GetString("host"),
		Port: viper.GetString("port"),
	}
}

func HTTPConfFromEnv() HTTPConf {
	viper.SetEnvPrefix("HTTP")
	viper.AutomaticEnv()
	return HTTPConf{
		Host: viper.GetString("host"),
		Port: viper.GetString("port"),
	}
}

func MQConfFromEnv() MQConf {
	viper.SetEnvPrefix("MQ")
	viper.AutomaticEnv()
	return MQConf{
		URI:      viper.GetString("uri"),
		Exchange: viper.GetString("exchange"),
		Key:      viper.GetString("key"),
		Queue:    viper.GetString("queue"),
	}
}
