package main

import (
	"os"

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

	postgreDB := store.NewPostgresDBStore(postgresDSN)
	domain := domain.NewInvoicesDomain(postgreDB)
	handler := handlers.NewAPIGatewayHandler(domain)

	lambda.Start(handler.GetHandler)
}
