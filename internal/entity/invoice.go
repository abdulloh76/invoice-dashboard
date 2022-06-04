package entity

type Address struct {
	Street   string `json:"street"`
	City     string `json:"city"`
	PostCode string `json:"postCode"`
	Country  string `json:"country"`
}

type Item struct {
	Name     string  `json:"name"`
	Quantity int     `json:"quantity"`
	Price    float32 `json:"price"`
	Total    float32 `json:"total"`
}

type Invoice struct {
	ID            uint64  `gorm:"primaryKey"`
	PaymentDue    string  `json:"paymentDue"`
	Description   string  `json:"description"`
	PaymentTerms  int     `json:"paymentTerms"`
	ClientName    string  `json:"clientName"`
	ClientEmail   string  `json:"clientEmail"`
	Status        string  `json:"status"`
	SenderAddress Address `json:"senderAddress"`
	ClientAddress Address `json:"clientAddress"`
	Items         []Item  `json:"items"`
	Total         float32 `json:"total"`
}
