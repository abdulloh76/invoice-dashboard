package main

import (
	"invoice-dashboard/internal/invoice"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	//run database
	// configs.ConnectDB()

	//routes
	invoice.RegisterHandlers(router)

	router.Run("localhost:8080")
}
