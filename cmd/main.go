package main

import (
	"os"
	"time"

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
	var cacheExpireDuration time.Duration = 600

	router := gin.Default()
	router.Use(middleware.CORSMiddleware())

	cache := store.NewRedisCacheStore(config.Configs.REDIS_URL, cacheExpireDuration)
	postgreDB := store.NewPostgresDBStore(config.Configs.DATABASE_URL)
	domain := domain.NewInvoicesDomain(postgreDB, cache)
	handler := handlers.NewGinAPIHandler(domain)

	handlers.RegisterHandlers(router, handler)

	router.Run(":" + config.Configs.PORT)
}
