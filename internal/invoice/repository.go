package invoice

import (
	"invoice-dashboard/internal/dto/invoiceDto"
	"invoice-dashboard/internal/entity"
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
	database := db.GetDB()

	return database.Transaction(func(tx *gorm.DB) error {
		if modifiedInvoice.ClientAddress.IsModified {
			if err := tx.Model(&entity.Address{}).Where(curInvoice.ClientAddress.ID).Updates(&entity.Address{
				Street:   modifiedInvoice.ClientAddress.Street,
				City:     modifiedInvoice.ClientAddress.City,
				PostCode: modifiedInvoice.ClientAddress.PostCode,
				Country:  modifiedInvoice.ClientAddress.Country,
			}).Error; err != nil {
				return err
			}
		}

		if modifiedInvoice.SenderAddress.IsModified {
			if err := tx.Model(&entity.Address{}).Where(curInvoice.SenderAddress.ID).Updates(&entity.Address{
				Street:   modifiedInvoice.SenderAddress.Street,
				City:     modifiedInvoice.SenderAddress.City,
				PostCode: modifiedInvoice.SenderAddress.PostCode,
				Country:  modifiedInvoice.SenderAddress.Country,
			}).Error; err != nil {
				return err
			}
		}

		for _, item := range modifiedInvoice.Items.ModifiedItems {
			if err := tx.Model(&entity.Item{}).Where(item.ID).Updates(&entity.Item{
				InvoiceID: curInvoice.ID,
				Name:      item.Name,
				Quantity:  item.Quantity,
				Price:     item.Price,
				Total:     item.Total,
			}).Error; err != nil {
				return err
			}
		}

		if len(modifiedInvoice.Items.DeletedItems) > 0 {
			// !that's very tricky coz what if client gives ids for items not in this invoice
			if err := tx.Unscoped().Delete(&entity.Item{}, modifiedInvoice.Items.DeletedItems).Error; err != nil {
				return err
			}
		}

		newItems := invoiceDto.PostItemsToEntitities(modifiedInvoice.Items.CreatedItems, curInvoice.ID)
		if err := tx.Model(&entity.Item{}).Create(newItems).Error; err != nil {
			return err
		}

		var items []entity.Item
		if err := tx.Model(&entity.Item{}).Select("total").Where("invoice_id = ?", curInvoice.ID).Find(&items).Error; err != nil {
			return err
		}
		var sumTotal float32 = 0
		for _, item := range items {
			sumTotal += item.Total
		}

		err := tx.Model(&entity.Invoice{}).Where("id = ?", curInvoice.ID).Updates(&entity.Invoice{
			PaymentDue:   modifiedInvoice.PaymentDue,
			Description:  modifiedInvoice.Description,
			PaymentTerms: modifiedInvoice.PaymentTerms,
			ClientName:   modifiedInvoice.ClientName,
			ClientEmail:  modifiedInvoice.ClientEmail,
			Status:       modifiedInvoice.Status,
			Total:        sumTotal,
		}).Error

		return err
	})
}

func RemoveInvoice(id string) error {
	database := db.GetDB()
	err := database.Unscoped().Delete(&entity.Invoice{}, "id = ?", id).Error
	// delete items, addresses
	return err
}
