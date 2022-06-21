package invoiceDto

import (
	"invoice-dashboard/internal/entity"
	"time"
)

type GetAddressDto struct {
	Street   string `json:"street"`
	City     string `json:"city"`
	PostCode string `json:"postCode"`
	Country  string `json:"country"`
}

type GetItemDto struct {
	ID       uint64  `json:"id"`
	Name     string  `json:"name"`
	Quantity int     `json:"quantity"`
	Price    float32 `json:"price"`
	Total    float32 `json:"total"`
}

type InvoiceResponse struct {
	ID            string        `json:"id"`
	PaymentDue    time.Time     `json:"paymentDue"`
	Description   string        `json:"description"`
	PaymentTerms  int           `json:"paymentTerms"`
	ClientName    string        `json:"clientName"`
	ClientEmail   string        `json:"clientEmail"`
	Status        string        `json:"status"`
	SenderAddress GetAddressDto `json:"senderAddress"`
	ClientAddress GetAddressDto `json:"clientAddress"`
	Items         []GetItemDto  `json:"items"`
	Total         float32       `json:"total"`
	CreatedAt     time.Time     `json:"createdAt"`
	UpdatedAt     time.Time     `json:"updatedAt"`
}

func EntitytoResponsetDTO(invoice *entity.Invoice) InvoiceResponse {
	var items []GetItemDto = make([]GetItemDto, len(invoice.Items))

	for i, item := range invoice.Items {
		items[i] = GetItemDto{
			ID:       item.ID,
			Name:     item.Name,
			Quantity: item.Quantity,
			Price:    item.Price,
			Total:    item.Total,
		}
	}

	getInvoiceDto := InvoiceResponse{
		ID:           invoice.ID,
		PaymentDue:   invoice.PaymentDue,
		Description:  invoice.Description,
		PaymentTerms: invoice.PaymentTerms,
		ClientName:   invoice.ClientName,
		ClientEmail:  invoice.ClientEmail,
		Status:       invoice.Status,
		SenderAddress: GetAddressDto{
			Street:   invoice.SenderAddress.Street,
			City:     invoice.SenderAddress.City,
			PostCode: invoice.SenderAddress.PostCode,
			Country:  invoice.SenderAddress.Country,
		},
		ClientAddress: GetAddressDto{
			Street:   invoice.ClientAddress.Street,
			City:     invoice.ClientAddress.City,
			PostCode: invoice.ClientAddress.PostCode,
			Country:  invoice.ClientAddress.Country,
		},
		Items:     items,
		Total:     invoice.Total,
		CreatedAt: invoice.CreatedAt,
		UpdatedAt: invoice.UpdatedAt,
	}

	return getInvoiceDto
}
