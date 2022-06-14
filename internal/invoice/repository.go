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

func InsertInvoice(invoice *entity.Invoice) *entity.Invoice {
	// consider beforeCreate hook for being in control of ID
	db.Create(&invoice)
	return invoice
}

func FindInvoices() []entity.Invoice {
	// consider querying the entities
	var invoices []entity.Invoice
	// db.Joins("Items").Joins("Address").Find(&invoices)
	db.Find(&invoices)
	return invoices
}

func FindInvoiceById() {}

func UpdateInvoice() {}

func RemoveInvoice() {}
