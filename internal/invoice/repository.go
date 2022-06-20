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

func InsertInvoice(invoice *entity.Invoice) error {
	// soomeday i think we need to consider to not duplicate address coz usually user doesn't change his address
	// consider beforeCreate hook for being in control of ID
	err := db.Create(&invoice).Error
	return err
}

func FindInvoices() ([]invoiceDto.InvoicesResponse, error) {
	// consider querying the entities for smth like filter
	var invoices []invoiceDto.InvoicesResponse
	err := db.Model(&entity.Invoice{}).Find(&invoices).Error
	return invoices, err
}

func FindInvoiceById(id string) (entity.Invoice, error) {
	var invoice entity.Invoice
	err := db.Preload("Items").Preload("SenderAddress").Preload("ClientAddress").First(&invoice, "id = ?", id).Error
	return invoice, err
}

func UpdateInvoice() {}

func RemoveInvoice(id string) error {
	err := db.Unscoped().Delete(&entity.Invoice{}, "id = ?", id).Error
	return err
}
