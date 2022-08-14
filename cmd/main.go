package main

import (
	"os"

	"github.com/abdulloh76/invoice-dashboard/domain"
	"github.com/abdulloh76/invoice-dashboard/handlers"
	"github.com/abdulloh76/invoice-dashboard/internal/config"
	"github.com/abdulloh76/invoice-dashboard/middleware"
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

	postgresDSN := config.Configs.POSTGRES_URI
	postgreDB := store.NewPostgresDBStore(postgresDSN)
	domain := domain.NewInvoicesDomain(postgreDB)
	handler := handlers.NewGinAPIHandler(domain)

	handlers.RegisterHandlers(router, handler)

	router.Run(":" + config.Configs.PORT)
}
