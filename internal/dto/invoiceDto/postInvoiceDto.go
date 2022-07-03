package invoiceDto

import (
	"invoice-dashboard/internal/entity"
	"time"
)

type PostAddressDto struct {
	Street   string `json:"street"`
	City     string `json:"city"`
	PostCode string `json:"postCode"`
	Country  string `json:"country"`
}

type PostItemDto struct {
	Name     string  `json:"name"`
	Quantity int     `json:"quantity"`
	Price    float32 `json:"price"`
	Total    float32 `json:"total"`
}

type InvoiceRequestBody struct {
	PaymentDue    time.Time      `json:"paymentDue"`
	Description   string         `json:"description"`
	PaymentTerms  int            `json:"paymentTerms"`
	ClientName    string         `json:"clientName"`
	ClientEmail   string         `json:"clientEmail"`
	Status        string         `json:"status"`
	SenderAddress PostAddressDto `json:"senderAddress"`
	ClientAddress PostAddressDto `json:"clientAddress"`
	Items         []PostItemDto  `json:"items"`
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
