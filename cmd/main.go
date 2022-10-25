package main

import (
	"github.com/abdulloh76/invoice-dashboard/pkg/domain"
	"github.com/abdulloh76/invoice-dashboard/pkg/handlers"
	"github.com/abdulloh76/invoice-dashboard/pkg/infrastructure"
	"github.com/abdulloh76/invoice-dashboard/pkg/store"
	"github.com/abdulloh76/invoice-dashboard/pkg/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	configs := utils.LoadConfig("config", "dev", "yml")

	router := gin.Default()
	router.Use(utils.CORSMiddleware())

	cache := store.NewRedisCacheStore(configs.REDIS_URL, configs.CACHE_EXPIRATION)
	grpcTarget := configs.GRPC_HOST + ":" + configs.GRPC_PORT
	userGrpcClient := infrastructure.NewUserGrpcClient(grpcTarget)
	postgresDB := store.NewPostgresDBStore(configs.DATABASE_URL)
	domain := domain.NewInvoicesDomain(postgresDB, cache)
	handler := handlers.NewGinAPIHandler(domain, userGrpcClient)

	handlers.RegisterHandlers(router, handler)

	router.Run(":" + configs.PORT)
}
