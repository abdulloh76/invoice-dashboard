package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func ConnectDB() {
	dsn := EnvVariables.DOCKER_POSTGRES_URI
	// dsn := EnvVariables.POSTGRES_URI

	d, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
