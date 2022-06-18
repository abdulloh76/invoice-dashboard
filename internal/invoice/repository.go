package invoice

import (
	"fmt"
	"invoice-dashboard/config"
	"invoice-dashboard/internal/dto/invoiceDto"
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

func FindInvoices() []invoiceDto.InvoicesResponse {
	// consider querying the entities for smth like filter
	var invoices []invoiceDto.InvoicesResponse

	db.Model(&entity.Invoice{}).Find(&invoices)
	return invoices
}

func FindInvoiceById(id uint64) entity.Invoice {
	var invoice entity.Invoice
	db.Preload("Items").Preload("SenderAddress").Preload("ClientAddress").First(&invoice, id)
	return invoice
}

func UpdateInvoice() {}

func RemoveInvoice() {}
