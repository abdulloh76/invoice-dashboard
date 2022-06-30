package main

import (
	"invoice-dashboard/internal/config"
	"invoice-dashboard/internal/entity"
	"invoice-dashboard/internal/invoice"
	"invoice-dashboard/internal/middleware"
	"invoice-dashboard/pkg/db"

	"github.com/gin-gonic/gin"
)

func init() {
	config.LoadConfig("./config", "dev", "yml")

	db.ConnectDB()
	database := db.GetDB()
	database.AutoMigrate(&entity.Address{})
	database.AutoMigrate(&entity.Invoice{})
	database.AutoMigrate(&entity.Item{})
}

func main() {
	router := gin.Default()
	router.Use(middleware.CORSMiddleware())

	invoice.RegisterHandlers(router)

	router.Run(":" + config.Configs.PORT)
}
