package invoiceDto

import (
	"invoice-dashboard/internal/entity"
	"time"
)

type PostAddressDto struct {
	Street   string
	City     string
	PostCode string
	Country  string
}

type PostItemDto struct {
	Name     string
	Quantity int
	Price    float32
	Total    float32
}

type InvoiceRequestBody struct {
	PaymentDue    time.Time
	Description   string
	PaymentTerms  int
	ClientName    string
	ClientEmail   string
	Status        string
	SenderAddress PostAddressDto
	ClientAddress PostAddressDto
	Items         []PostItemDto
}

func RequestDTOtoEntity(dto *InvoiceRequestBody) entity.Invoice {
	var items []entity.Item = make([]entity.Item, len(dto.Items))
	var total float32 = 0

	for i, dtoItem := range dto.Items {
		items[i] = entity.Item{
			Name:     dtoItem.Name,
			Quantity: dtoItem.Quantity,
			Price:    dtoItem.Price,
			Total:    dtoItem.Total,
		}
		total += dtoItem.Total
	}

	invoice := entity.Invoice{
		PaymentDue:   dto.PaymentDue,
		Description:  dto.Description,
		PaymentTerms: dto.PaymentTerms,
		ClientName:   dto.ClientName,
		ClientEmail:  dto.ClientEmail,
		Status:       dto.Status,
		SenderAddress: entity.Address{
			Street:   dto.SenderAddress.Street,
			City:     dto.SenderAddress.City,
			PostCode: dto.SenderAddress.PostCode,
			Country:  dto.SenderAddress.Country,
		},
		ClientAddress: entity.Address{
			Street:   dto.ClientAddress.Street,
			City:     dto.ClientAddress.City,
			PostCode: dto.ClientAddress.PostCode,
			Country:  dto.ClientAddress.Country,
		},
		Items: items,
		Total: total,
	}

	return invoice
}
