package types

type InvoiceStore interface {
	InsertInvoice(invoice *InvoiceModel) error
	FindInvoices() ([]GetInvoicesResponse, error)
	FindInvoiceById(id string) (*InvoiceModel, error)
	ModifyInvoice(curInvoice *InvoiceModel, modifiedInvoice PutInvoiceBody) error
	RemoveInvoice(id string) error
}
