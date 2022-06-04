package main

import (
	"invoice-dashboard/config"
	"invoice-dashboard/internal/invoice"

	"github.com/gin-gonic/gin"
)

func init() {
	// load env vars to global variable first
	config.LoadConfig(".")
	config.ConnectDB()
}

func main() {
	router := gin.Default()

	invoice.RegisterHandlers(router)

	router.Run(":" + config.EnvVariables.PORT)
}
