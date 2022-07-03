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

func ModifyInvoice(id string, modifiedInvoice invoiceDto.PutInvoiceBody) error {
	database := db.GetDB()

	if modifiedInvoice.ClientAddress.IsModified {
		database.Model(&entity.Address{}).Where("invoice_id = ?", id).Updates(modifiedInvoice.ClientAddress)
	}
	if modifiedInvoice.SenderAddress.IsModified {
		database.Model(&entity.Address{}).Where("invoice_id = ?", id).Updates(modifiedInvoice.SenderAddress)
	}

	database.Model(&entity.Item{}).Updates(modifiedInvoice.Items.ModifiedItems)
	database.Unscoped().Delete(&entity.Item{}, modifiedInvoice.Items.DeletedItems)

	newItems := invoiceDto.PostItemToEntity(modifiedInvoice.Items.CreatedItems, id)
	database.Model(&entity.Item{}).Create(newItems)

	err := database.Model(&entity.Invoice{}).Where("id = ?", id).Updates(modifiedInvoice).Error

	return err
}

func RemoveInvoice(id string) error {
	database := db.GetDB()
	err := database.Unscoped().Delete(&entity.Invoice{}, "id = ?", id).Error
	// delete items, addresses
	return err
}
