package invoice

import (
	"fmt"
	"invoice-dashboard/config"
	"invoice-dashboard/internal/entity"

	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	config.ConnectDB()
	db = config.GetDB()

	db.AutoMigrate(&entity.Address{})
	db.AutoMigrate(&entity.Invoice{})
	db.AutoMigrate(&entity.Item{})

	fmt.Println(db)
}

func InsertInvoice(invoice *entity.Invoice) {
	// soomeday i think we need to consider to not duplicate address coz usually user doesn't change his address
	// consider beforeCreate hook for being in control of ID
	db.Create(&invoice)
}

func FindInvoices() []entity.Invoice {
	// consider querying the entities
	var invoices []entity.Invoice
	db.Preload("Items").Preload("SenderAddress").Preload("ClientAddress").Find(&invoices)
	return invoices
}

func FindInvoiceById() {}

func UpdateInvoice() {}

func RemoveInvoice() {}
