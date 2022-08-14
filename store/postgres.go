package store

import (
	"github.com/abdulloh76/invoice-dashboard/types"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgresDBStore struct {
	db *gorm.DB
}

var _ types.InvoiceStore = (*PostgresDBStore)(nil)

func NewPostgresDBStore(dsn string) *PostgresDBStore {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&types.AddressModel{})
	db.AutoMigrate(&types.InvoiceModel{})
	db.AutoMigrate(&types.ItemModel{})

	return &PostgresDBStore{
		db,
	}
}

func (d *PostgresDBStore) InsertInvoice(invoice *types.InvoiceModel) error {
	// soomeday i think we need to consider to not duplicate address coz usually user doesn't change his address
	err := d.db.Create(&invoice).Error

	return err
}

func (d *PostgresDBStore) FindInvoices() ([]types.GetInvoicesResponse, error) {
	// consider querying the entities for smth like filter
	var invoices []types.GetInvoicesResponse
	err := d.db.Model(&types.InvoiceModel{}).Find(&invoices).Error
	return invoices, err
}

func (d *PostgresDBStore) FindInvoiceById(id string) (*types.InvoiceModel, error) {
	var invoice types.InvoiceModel
	err := d.db.Preload("Items").Preload("SenderAddress").Preload("ClientAddress").First(&invoice, "id = ?", id).Error
	return &invoice, err
}

func (d *PostgresDBStore) ModifyInvoice(curInvoice *types.InvoiceModel, modifiedInvoice types.PutInvoiceBody) error {
	return d.db.Transaction(func(tx *gorm.DB) error {
		modifyAddress(&curInvoice.ClientAddress, &modifiedInvoice.ClientAddress)
		modifyAddress(&curInvoice.SenderAddress, &modifiedInvoice.SenderAddress)

		batchUpdateItems(curInvoice.Items, modifiedInvoice.Items.ModifiedItems)

		deletedItems := batchDeleteItems(&curInvoice.Items, modifiedInvoice.Items.DeletedItems)
		tx.Model(&curInvoice).Association("Items").Delete(deletedItems)

		newItems := types.PostItemsToEntitities(&modifiedInvoice.Items.CreatedItems, curInvoice.ID)
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

func (d *PostgresDBStore) RemoveInvoice(id string) error {
	err := d.db.Unscoped().Delete(&types.InvoiceModel{}, "id = ?", id).Error
	// delete items, addresses
	return err
}

func modifyAddress(address *types.AddressModel, modifiedAddress *types.PutAddressDto) {
	if !modifiedAddress.IsModified {
		return
	}
	address.Street = modifiedAddress.Street
	address.City = modifiedAddress.City
	address.PostCode = modifiedAddress.PostCode
	address.Country = modifiedAddress.Country
}

func batchUpdateItems(curItems []types.ItemModel, modifiedItems []types.PutItemDto) {
	// !consider refactoring the search coz both slices are sorted
	for modIdx := range modifiedItems {
		for curIdx := range curItems {
			if curItems[curIdx].ID == modifiedItems[modIdx].ID {
				curItems[curIdx].Name = modifiedItems[modIdx].Name
				curItems[curIdx].Quantity = modifiedItems[modIdx].Quantity
				curItems[curIdx].Price = modifiedItems[modIdx].Price
				curItems[curIdx].Total = modifiedItems[modIdx].Total
			}
		}
	}

}

func batchDeleteItems(curItems *[]types.ItemModel, deleteItemIds []uint64) []types.ItemModel {
	var deleteItems []types.ItemModel

	// !consider refactoring the search coz both slices are sorted?
	for _, deleteId := range deleteItemIds {
		newLength := 0
		for _, curItem := range *curItems {
			if curItem.ID != deleteId {
				(*curItems)[newLength] = curItem
				newLength++
			} else {
				deleteItems = append(deleteItems, curItem)
			}
		}
		(*curItems) = (*curItems)[:newLength]
	}

	return deleteItems
}
