package invoiceDto

import (
	"invoice-dashboard/internal/entity"
	"time"
)

type Address struct {
	Street   string
	City     string
	PostCode string
	Country  string
}

type Item struct {
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
	SenderAddress Address
	ClientAddress Address
	Items         []Item
	Total         float32
}

func RequestDTOtoEntity(dto *InvoiceRequestBody) entity.Invoice {
	var items []entity.Item = make([]entity.Item, len(dto.Items))

	for i, item := range dto.Items {
		items[i] = entity.Item{
			Name:     item.Name,
			Quantity: item.Quantity,
			Price:    item.Price,
			Total:    item.Total,
		}
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
		Total: dto.Total,
	}

	return invoice
}
