package invoice

import (
	"invoice-dashboard/config"
	"invoice-dashboard/internal/entity"

	"gorm.io/gorm"
)

var db *gorm.DB = config.GetDB()

func Create(invoice *entity.Invoice) *entity.Invoice {
	// consider beforeCreate hook for being in control of ID
	db.Create(&invoice)
	return invoice
}

func FindInvoices() []entity.Invoice {
	// consider querying the entities
	var invoices []entity.Invoice
	db.Find(&invoices)
	return invoices
}

func FindInvoiceById() {

}

func UpdateInvoice() {

}

func Delete() {

}
