package config

import "github.com/spf13/viper"

type EnvVariables struct {
	PORT string
}

func LoadConfig(path string) (envs EnvVariables, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&envs)
	return
}
