package main

import (
	"os"

	"github.com/abdulloh76/invoice-dashboard/domain"
	"github.com/abdulloh76/invoice-dashboard/handlers"
	"github.com/abdulloh76/invoice-dashboard/internal/config"
	"github.com/abdulloh76/invoice-dashboard/internal/middleware"
	"github.com/abdulloh76/invoice-dashboard/store"

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

	store.NewRedisCacheStore(config.Configs.REDIS_URL, 1, 10)

	postgreDB := store.NewPostgresDBStore(config.Configs.DATABASE_URL)
	domain := domain.NewInvoicesDomain(postgreDB)
	handler := handlers.NewGinAPIHandler(domain)

	handlers.RegisterHandlers(router, handler)

	router.Run(":" + config.Configs.PORT)
}
