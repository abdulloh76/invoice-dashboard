package utils

import (
	"time"

	"github.com/spf13/viper"
)

type ConfigStruct struct {
	DATABASE_URL     string
	REDIS_URL        string
	CACHE_EXPIRATION time.Duration
	GRPC_PORT        string
}

func LoadConfig(configPath, configName, configType string) *ConfigStruct {
	viper.AddConfigPath(configPath)
	viper.SetConfigName(configName)
	viper.SetConfigType(configType)

	err := viper.ReadInConfig()
	if err != nil {
		panic("Need environment variables file")
	}

	var cacheExpireDuration time.Duration = 600

	return &ConfigStruct{
		DATABASE_URL:     viper.GetString("DATABASE_URL"),
		REDIS_URL:        viper.GetString("REDIS_URL"),
		GRPC_PORT:        viper.GetString("GRPC_PORT"),
		CACHE_EXPIRATION: cacheExpireDuration,
	}
}
