package utils

import (
	"os"
	"time"
)

type ConfigStruct struct {
	DATABASE_URL     string
	REDIS_URL        string
	CACHE_EXPIRATION time.Duration
}

func InitializeConfigs() *ConfigStruct {
	postgresDSN, ok := os.LookupEnv("DATABASE_URL")
	if !ok {
		panic("Need DATABASE_URL environment variable")
	}

	var cacheExpireDuration time.Duration = 600
	redisURL, ok := os.LookupEnv("REDIS_URL")
	if !ok {
		panic("Need REDIS_URL environment variable")
	}

	return &ConfigStruct{
		DATABASE_URL:     postgresDSN,
		REDIS_URL:        redisURL,
		CACHE_EXPIRATION: cacheExpireDuration,
	}
}
