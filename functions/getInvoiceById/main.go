package main

import (
	"os"
	"time"

	"github.com/abdulloh76/invoice-dashboard/domain"
	"github.com/abdulloh76/invoice-dashboard/handlers"
	"github.com/abdulloh76/invoice-dashboard/store"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	postgresDSN, ok := os.LookupEnv("DATABASE_URL")
	if !ok {
		panic("Need DATABASE_URL environment variable")
	}

	var cacheExpireDuration time.Duration = 600
	redisURL, ok := os.LookupEnv("REDIS_URL")
	if !ok {
		panic("Need REDIS_URL environment variable")
	}

	cache := store.NewRedisCacheStore(redisURL, cacheExpireDuration)
	postgreDB := store.NewPostgresDBStore(postgresDSN)
	domain := domain.NewInvoicesDomain(postgreDB, cache)
	handler := handlers.NewAPIGatewayHandler(domain)

	lambda.Start(handler.GetHandler)
}
