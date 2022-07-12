package invoiceDto

import (
	"invoice-dashboard/internal/entity"
	"time"
)

type PutAddressDto struct {
	Street     string `json:"street"`
	City       string `json:"city"`
	PostCode   string `json:"postCode"`
	Country    string `json:"country"`
	IsModified bool   `json:"isModified"`
}

type PutItemDto struct {
	ID       uint64  `json:"id"`
	Name     string  `json:"name"`
	Quantity int     `json:"quantity"`
	Price    float32 `json:"price"`
	Total    float32 `json:"total"`
}

type PutItemsDto struct {
	CreatedItems  []PostItemDto `json:"createdItems"`
	ModifiedItems []PutItemDto  `json:"modifiedItems"`
	DeletedItems  []uint64      `json:"deletedItems"`
}

type PutInvoiceBody struct {
	PaymentDue    time.Time     `json:"paymentDue"`
	Description   string        `json:"description"`
	PaymentTerms  int           `json:"paymentTerms"`
	ClientName    string        `json:"clientName"`
	ClientEmail   string        `json:"clientEmail"`
	Status        string        `json:"status"`
	SenderAddress PutAddressDto `json:"senderAddress"`
	ClientAddress PutAddressDto `json:"clientAddress"`
	Items         PutItemsDto   `json:"items"`
}

func PostItemsToEntitities(newItems []PostItemDto, invoiceId string) []entity.Item {
	var items []entity.Item = make([]entity.Item, len(newItems))

	for i, dto := range newItems {
		items[i] = entity.Item{
			InvoiceID: invoiceId,
			Name:      dto.Name,
			Quantity:  dto.Quantity,
			Price:     dto.Price,
			Total:     dto.Total,
		}
	}

	return items
}
