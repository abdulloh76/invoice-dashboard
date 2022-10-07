package main

import (
	"os"

	"github.com/abdulloh76/invoice-dashboard/pkg/domain"
	"github.com/abdulloh76/invoice-dashboard/pkg/handlers"
	"github.com/abdulloh76/invoice-dashboard/pkg/store"
	"github.com/abdulloh76/invoice-dashboard/pkg/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Use(utils.CORSMiddleware())

	configs := utils.InitializeConfigs()
	cache := store.NewRedisCacheStore(configs.REDIS_URL, configs.CACHE_EXPIRATION)
	postgreDB := store.NewPostgresDBStore(configs.DATABASE_URL)
	domain := domain.NewInvoicesDomain(postgreDB, cache)
	handler := handlers.NewGinAPIHandler(domain)

	handlers.RegisterHandlers(router, handler)

	router.Run(":" + os.Getenv("PORT"))
}
