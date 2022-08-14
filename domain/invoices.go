package domain

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/abdulloh76/invoice-dashboard/types"
)

var (
	ErrJsonUnmarshal = errors.New("failed to parse invoice from request body")
	ErrUserNotFound  = errors.New("invoice with given ID not found")
)

type Invoices struct {
	store types.InvoiceStore
}

func NewInvoicesDomain(s types.InvoiceStore) *Invoices {
	return &Invoices{
		store: s,
	}
}

func (i *Invoices) GetSingleInvoice(id string) (*types.InvoiceModel, error) {
	invoice, err := i.store.FindInvoiceById(id)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	return invoice, nil
}

func (i *Invoices) AllInvoices() ([]types.GetInvoicesResponse, error) {
	allInvoices, err := i.store.FindInvoices()
	if err != nil {
		return allInvoices, fmt.Errorf("%w", err)
	}

	return allInvoices, nil
}

func (i *Invoices) Create(body []byte) (*types.InvoiceModel, error) {
	invoice := types.InvoiceRequestBody{}
	if err := json.Unmarshal(body, &invoice); err != nil {
		return nil, fmt.Errorf("%w", ErrJsonUnmarshal)
	}

	newInvoice := types.RequestDTOtoEntity(&invoice)

	err := i.store.InsertInvoice(&newInvoice)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	return &newInvoice, nil
}

func (i *Invoices) ModifyInvoice(id string, body []byte) (*types.InvoiceModel, error) {
	curInvoice, notFoundErr := i.store.FindInvoiceById(id)
	if notFoundErr != nil {
		return nil, fmt.Errorf("%w", ErrJsonUnmarshal)
	}

	modifiedInvoice := types.PutInvoiceBody{}
	if err := json.Unmarshal(body, &modifiedInvoice); err != nil {
		return nil, fmt.Errorf("%w", ErrJsonUnmarshal)
	}

	err := i.store.ModifyInvoice(curInvoice, modifiedInvoice)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	return curInvoice, nil
}

func (i *Invoices) DeleteInvoice(id string) error {
	err := i.store.RemoveInvoice(id)
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	return nil
}
