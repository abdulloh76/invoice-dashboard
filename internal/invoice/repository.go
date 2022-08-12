package invoice

import (
	"invoice-dashboard/internal/address"
	"invoice-dashboard/internal/dto/invoiceDto"
	"invoice-dashboard/internal/entity"
	"invoice-dashboard/internal/invoiceItem"
	"invoice-dashboard/pkg/db"

	"gorm.io/gorm"
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

func ModifyInvoice(curInvoice *entity.Invoice, modifiedInvoice invoiceDto.PutInvoiceBody) error {

	return db.GetDB().Transaction(func(tx *gorm.DB) error {
		address.ModifyAddress(&curInvoice.ClientAddress, &modifiedInvoice.ClientAddress)
		address.ModifyAddress(&curInvoice.SenderAddress, &modifiedInvoice.SenderAddress)

		invoiceItem.BatchUpdateItems(curInvoice.Items, modifiedInvoice.Items.ModifiedItems)

		deletedItems := invoiceItem.BatchDeleteItems(&curInvoice.Items, modifiedInvoice.Items.DeletedItems)
		tx.Model(&curInvoice).Association("Items").Delete(deletedItems)

		newItems := invoiceDto.PostItemsToEntitities(&modifiedInvoice.Items.CreatedItems, curInvoice.ID)
		curInvoice.Items = append(curInvoice.Items, newItems...)

		var sumTotal float32 = 0
		for _, item := range curInvoice.Items {
			sumTotal += item.Total
		}
		curInvoice.Total = sumTotal
		curInvoice.PaymentDue = modifiedInvoice.PaymentDue
		curInvoice.Description = modifiedInvoice.Description
		curInvoice.PaymentTerms = modifiedInvoice.PaymentTerms
		curInvoice.ClientName = modifiedInvoice.ClientName
		curInvoice.ClientEmail = modifiedInvoice.ClientEmail
		curInvoice.Status = modifiedInvoice.Status

		err := tx.Session(&gorm.Session{FullSaveAssociations: true}).Updates(&curInvoice).Error

		return err
	})
}

func RemoveInvoice(id string) error {
	database := db.GetDB()
	err := database.Unscoped().Delete(&entity.Invoice{}, "id = ?", id).Error
	// delete items, addresses
	return err
}
