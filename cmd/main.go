package main

import (
	"os"

	"github.com/abdulloh76/invoice-dashboard/handlers"
	"github.com/abdulloh76/invoice-dashboard/internal/config"
	"github.com/abdulloh76/invoice-dashboard/middleware"

	"github.com/gin-gonic/gin"
)

func init() {
	if os.Getenv("PORT") == "" {
		config.LoadConfig("./config", "dev", "yml")
	} else {
		// * inside heroku
		config.InitializeFromOS()
	}
}

func main() {
	router := gin.Default()
	router.Use(middleware.CORSMiddleware())

	handlers.RegisterHandlers(router)

	router.Run(":" + config.Configs.PORT)
}
