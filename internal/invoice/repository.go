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

	invoicePrototype := &entity.Invoice{}
	adsressPrototype := &entity.Address{}
	itemPrototype := &entity.Item{}

	db.AutoMigrate(adsressPrototype)
	db.AutoMigrate(itemPrototype)
	db.AutoMigrate(invoicePrototype)

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
	db.Joins("Items").Joins("Address").Find(&invoices)
	return invoices
}

func FindInvoiceById() {}

func UpdateInvoice() {}

func RemoveInvoice() {}
