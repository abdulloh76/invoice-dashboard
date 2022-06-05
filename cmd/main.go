package main

import (
	"invoice-dashboard/config"
	"invoice-dashboard/internal/invoice"

	"github.com/gin-gonic/gin"
)

func init() {
	// load env vars to global variable first
	config.LoadConfig(".")
	config.ConnectDB()
}

func main() {
	// 	if db, err := Connect(); err != nil {
	// 		fmt.Printf("Dude! I could not connect to the database. This happened: %s. Please fix everything and try again", err)
	//  } else {
	// 		defer db.Close()
	// 		Migrate(db)
	// 		Seed(db) // just initializing the db with data
	// 		ListEverything(db)
	// 		ClearEverything(db)
	//  }

	router := gin.Default()

	invoice.RegisterHandlers(router)

	router.Run(":" + config.EnvVariables.PORT)
}
