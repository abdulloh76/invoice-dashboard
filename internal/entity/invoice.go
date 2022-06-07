package entity

import (
	"time"

	"gorm.io/gorm"
)

type Invoice struct {
	gorm.Model
	ID              uint64    `gorm:"primaryKey;autoincrement:true" json:"id" binding:"required"`
	PaymentDue      time.Time `json:"paymentDue" binding:"required"`
	Description     string    `json:"description" binding:"required"`
	PaymentTerms    int       `json:"paymentTerms" binding:"required"`
	ClientName      string    `json:"clientName" binding:"required"`
	ClientEmail     string    `json:"clientEmail" binding:"required"`
	Status          string    `json:"status" binding:"required"`
	SenderAddressId uint64
	ClientAddressId uint64
	SenderAddress   *Address       `gorm:"foreignKey:ID;references:SenderAddressId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"senderAddress" binding:"required"`
	ClientAddress   *Address       `gorm:"foreignKey:ID;references:ClientAddressId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"clientAddress" binding:"required"`
	Items           []Item         `gorm:"foreignKey:InvoiceID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"items" binding:"required,dive"`
	Total           float32        `json:"total" binding:"required"`
	CreatedAt       time.Time      `json:"createdAt" binding:"required"`
	UpdatedAt       time.Time      `json:"updatedAt" binding:"required"`
	DeletedAt       gorm.DeletedAt `gorm:"index" json:"deletedAt" binding:"required"`
}

type Address struct {
	gorm.Model
	ID       uint64 `gorm:"primaryKey"`
	Street   string `json:"street" binding:"required"`
	City     string `json:"city" binding:"required"`
	PostCode string `json:"postCode" binding:"required"`
	Country  string `json:"country" binding:"required"`
}

type Item struct {
	gorm.Model
	InvoiceID uint64
	ID        uint64  `gorm:"primaryKey"`
	Name      string  `json:"name" binding:"required"`
	Quantity  int     `json:"quantity" binding:"required"`
	Price     float32 `json:"price" binding:"required"`
	Total     float32 `json:"total" binding:"required"`
}
