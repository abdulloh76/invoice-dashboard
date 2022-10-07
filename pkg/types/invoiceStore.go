package types

import "context"

type InvoiceStore interface {
	InsertInvoice(invoice *InvoiceModel) error
	FindInvoices() ([]GetInvoicesResponse, error)
	FindInvoiceById(id string) (*InvoiceModel, error)
	ModifyInvoice(curInvoice *InvoiceModel, modifiedInvoice PutInvoiceBody) error
	RemoveInvoice(id string) error
}

type InvoiceCacheStore interface {
	Set(ctx context.Context, key string, invoice *InvoiceModel)
	Get(ctx context.Context, key string) *InvoiceModel
	Delete(ctx context.Context, key string) error
}
