package config

import (
	"log"

	"github.com/spf13/viper"
)

var (
	Configs ConfigStruct
)

type ConfigStruct struct {
	PORT         string
	POSTGRES_URI string
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
	Configs.POSTGRES_URI = viper.GetString("POSTGRES_URI")
}
