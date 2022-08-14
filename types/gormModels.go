package types

import (
	"errors"
	"time"

	"github.com/teris-io/shortid"
	"gorm.io/gorm"
)

type InvoiceModel struct {
	gorm.Model
	ID              string
	PaymentDue      time.Time
	Description     string
	PaymentTerms    int
	ClientName      string
	ClientEmail     string
	Status          string
	SenderAddressId *uint64
	ClientAddressId *uint64
	SenderAddress   AddressModel `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	ClientAddress   AddressModel `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Items           []ItemModel  `gorm:"foreignKey:InvoiceID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Total           float32
}

type AddressModel struct {
	gorm.Model
	ID       uint64
	Street   string
	City     string
	PostCode string
	Country  string
}

type ItemModel struct {
	gorm.Model
	ID        uint64
	InvoiceID string
	Name      string
	Quantity  int
	Price     float32
	Total     float32
}

func (InvoiceModel) TableName() string {
	return "invoices"
}
func (AddressModel) TableName() string {
	return "address"
}
func (ItemModel) TableName() string {
	return "items"
}

func (invoice *InvoiceModel) BeforeCreate(tx *gorm.DB) (err error) {
	invoice.ID, err = shortid.Generate()

	if err != nil {
		return errors.New("can't save invalid data")
	}
	return nil
}
