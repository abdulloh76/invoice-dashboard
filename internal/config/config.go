package config

import (
	"log"
	"os"

	"github.com/spf13/viper"
)

var (
	Configs ConfigStruct
)

type ConfigStruct struct {
	PORT           string
	DATABASE_URL   string
	REDIS_PASSWORD string
	REDIS_ADDRESS  string
}

func GetEnvConfigs() *ConfigStruct {
	return &Configs
}

func LoadConfig(configPath, configName, configType string) {
	viper.AddConfigPath(configPath)
	viper.SetConfigName(configName)
	viper.SetConfigType(configType)

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	Configs.PORT = viper.GetString("PORT")
	Configs.DATABASE_URL = viper.GetString("DATABASE_URL")
	Configs.REDIS_PASSWORD = viper.GetString("REDIS_PASSWORD")
	Configs.REDIS_ADDRESS = viper.GetString("REDIS_ADDRESS")
}

func InitializeFromOS() {
	Configs.PORT = os.Getenv("PORT")
	Configs.DATABASE_URL = os.Getenv("DATABASE_URL")
	Configs.REDIS_PASSWORD = os.Getenv("REDIS_PASSWORD")
	Configs.REDIS_ADDRESS = os.Getenv("REDIS_ADDRESS")
}
