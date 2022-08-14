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
	PORT         string
	DATABASE_URL string
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
}

func InitializeFromOS() {
	Configs.DATABASE_URL = os.Getenv("DATABASE_URL")
	Configs.PORT = os.Getenv("PORT")
}
