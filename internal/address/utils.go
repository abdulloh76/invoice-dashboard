package address

import (
	"invoice-dashboard/internal/dto/invoiceDto"
	"invoice-dashboard/internal/entity"
)

func ModifyAddress(entityAddress *entity.Address, modifiedAddress *invoiceDto.PutAddressDto) {
	if !modifiedAddress.IsModified {
		return
	}
	entityAddress.Street = modifiedAddress.Street
	entityAddress.City = modifiedAddress.City
	entityAddress.PostCode = modifiedAddress.PostCode
	entityAddress.Country = modifiedAddress.Country
}
