package address

import (
	"invoice-dashboard/internal/dto/invoiceDto"
	"invoice-dashboard/internal/entity"

	"gorm.io/gorm"
)

func ModifyAddress(tx *gorm.DB, addressId uint64, modifiedAddress *invoiceDto.PutAddressDto) error {
	err := tx.Model(&entity.Address{}).Where(addressId).Updates(&entity.Address{
		Street:   modifiedAddress.Street,
		City:     modifiedAddress.City,
		PostCode: modifiedAddress.PostCode,
		Country:  modifiedAddress.Country,
	}).Error

	return err
}
