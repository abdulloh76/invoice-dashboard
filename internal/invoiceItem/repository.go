package invoiceItem

import (
	"invoice-dashboard/internal/dto/invoiceDto"
	"invoice-dashboard/internal/entity"

	"gorm.io/gorm"
)

func BatchInsertItems(tx *gorm.DB, invoiceId string, createItems *[]invoiceDto.PostItemDto) error {
	newItems := invoiceDto.PostItemsToEntitities(createItems, invoiceId)
	err := tx.Model(&entity.Item{}).Create(newItems).Error

	return err
}

func ModifyItem(tx *gorm.DB, invoiceId string, modifiedItem *invoiceDto.PutItemDto) error {
	err := tx.Model(&entity.Item{}).Where(modifiedItem.ID).Updates(&entity.Item{
		InvoiceID: invoiceId,
		Name:      modifiedItem.Name,
		Quantity:  modifiedItem.Quantity,
		Price:     modifiedItem.Price,
		Total:     modifiedItem.Total,
	}).Error

	return err
}

func BatchDeleteItems(tx *gorm.DB, invoiceId string, deleteItemIds []uint64) error {
	err := tx.Unscoped().Where("invoice_id = ? AND id IN ?", invoiceId, deleteItemIds).Delete(&entity.Item{}).Error

	return err
}
