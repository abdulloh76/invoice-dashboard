package main

import (
	"github.com/abdulloh76/invoice-dashboard/domain"
	"github.com/abdulloh76/invoice-dashboard/handlers"
	"github.com/abdulloh76/invoice-dashboard/internal/config"
	"github.com/abdulloh76/invoice-dashboard/store"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	configs := config.InitializeConfigs()
	cache := store.NewRedisCacheStore(configs.REDIS_URL, configs.CACHE_EXPIRATION)
	postgreDB := store.NewPostgresDBStore(configs.DATABASE_URL)
	domain := domain.NewInvoicesDomain(postgreDB, cache)
	handler := handlers.NewAPIGatewayHandler(domain)

	lambda.Start(handler.AllHandler)
}
