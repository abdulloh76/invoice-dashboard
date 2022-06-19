package entity

import (
	"errors"
	"time"

	"github.com/teris-io/shortid"
	"gorm.io/gorm"
)

type Invoice struct {
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
	SenderAddress   Address `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	ClientAddress   Address `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Items           []Item  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Total           float32
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

type Address struct {
	gorm.Model
	ID       uint64
	Street   string
	City     string
	PostCode string
	Country  string
}

type Item struct {
	gorm.Model
	ID        uint64
	InvoiceID string
	Name      string
	Quantity  int
	Price     float32
	Total     float32
}

func (invoice *Invoice) BeforeCreate(tx *gorm.DB) (err error) {
	invoice.ID, err = shortid.Generate()

	if err != nil {
		return errors.New("can't save invalid data")
	}
	return nil
}
