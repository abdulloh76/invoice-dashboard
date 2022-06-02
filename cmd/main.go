package main

import (
	"invoice-dashboard/config"
	"invoice-dashboard/internal/invoice"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	config, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	router := gin.Default()

	//run database
	// configs.ConnectDB()

	//routes
	invoice.RegisterHandlers(router)

	router.Run(":" + config.PORT)
}
