package invoice

import (
	"invoice-dashboard/internal/dto/invoiceDto"
	"invoice-dashboard/internal/entity"
	"invoice-dashboard/pkg/db"
)

func InsertInvoice(invoice *entity.Invoice) error {
	// soomeday i think we need to consider to not duplicate address coz usually user doesn't change his address
	database := db.GetDB()
	err := database.Create(&invoice).Error
	return err
}

func FindInvoices() ([]invoiceDto.InvoicesResponse, error) {
	// consider querying the entities for smth like filter
	database := db.GetDB()
	var invoices []invoiceDto.InvoicesResponse
	err := database.Model(&entity.Invoice{}).Find(&invoices).Error
	return invoices, err
}

func FindInvoiceById(id string) (entity.Invoice, error) {
	database := db.GetDB()
	var invoice entity.Invoice
	err := database.Preload("Items").Preload("SenderAddress").Preload("ClientAddress").First(&invoice, "id = ?", id).Error
	return invoice, err
}

func UpdateInvoice() {}

func RemoveInvoice(id string) error {
	database := db.GetDB()
	err := database.Unscoped().Delete(&entity.Invoice{}, "id = ?", id).Error
	return err
}
