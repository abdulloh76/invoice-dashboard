package invoiceDto

import (
	"time"
)

type InvoicesResponse struct {
	ID         string    `json:"id"`
	PaymentDue time.Time `json:"paymentDue"`
	ClientName string    `json:"clientName"`
	Status     string    `json:"status"`
	Total      float32   `json:"total"`
}
