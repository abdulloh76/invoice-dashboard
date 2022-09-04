package domain

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/abdulloh76/invoice-dashboard/types"
)

var (
	ErrJsonUnmarshal   = errors.New("failed to parse invoice from request body")
	ErrInvoiceNotFound = errors.New("invoice with given ID not found")
)

type Invoices struct {
	store types.InvoiceStore
	cache types.InvoiceCacheStore
}

func NewInvoicesDomain(s types.InvoiceStore, c types.InvoiceCacheStore) *Invoices {
	return &Invoices{
		store: s,
		cache: c,
	}
}

func (i *Invoices) GetSingleInvoice(id string) (*types.InvoiceModel, error) {
	ctx := context.Background()
	var err error = nil
	var invoice *types.InvoiceModel = i.cache.Get(ctx, id)

	if invoice == nil {
		invoice, err = i.store.FindInvoiceById(id)
		if invoice.ID == "" {
			return nil, ErrInvoiceNotFound
		}
		if err != nil {
			return nil, err
		}
		i.cache.Set(ctx, id, invoice)
		return invoice, nil
	} else {
		return invoice, nil
	}
}

func (i *Invoices) AllInvoices() ([]types.GetInvoicesResponse, error) {
	allInvoices, err := i.store.FindInvoices()
	if err != nil {
		return allInvoices, err
	}

	return allInvoices, nil
}

func (i *Invoices) Create(body []byte) (*types.InvoiceModel, error) {
	ctx := context.Background()
	invoice := types.InvoiceRequestBody{}
	if err := json.Unmarshal(body, &invoice); err != nil {
		return nil, ErrJsonUnmarshal
	}

	newInvoice := types.RequestDTOtoEntity(&invoice)

	err := i.store.InsertInvoice(&newInvoice)
	if err != nil {
		return nil, err
	}

	i.cache.Set(ctx, newInvoice.ID, &newInvoice)
	return &newInvoice, nil
}

func (i *Invoices) ModifyInvoice(id string, body []byte) (*types.InvoiceModel, error) {
	ctx := context.Background()
	curInvoice, err := i.store.FindInvoiceById(id)
	if curInvoice == nil {
		return nil, ErrInvoiceNotFound
	}
	if err != nil {
		return nil, err
	}

	modifiedInvoice := types.PutInvoiceBody{}
	if err := json.Unmarshal(body, &modifiedInvoice); err != nil {
		return nil, ErrJsonUnmarshal
	}

	modifyErr := i.store.ModifyInvoice(curInvoice, modifiedInvoice)
	if modifyErr != nil {
		return nil, modifyErr
	}

	i.cache.Set(ctx, id, curInvoice)
	return curInvoice, nil
}

func (i *Invoices) DeleteInvoice(id string) error {
	ctx := context.Background()

	cacheErr := i.cache.Delete(ctx, id)
	if cacheErr != nil {
		return cacheErr
	}

	err := i.store.RemoveInvoice(id)
	if err != nil {
		return err
	}

	return nil
}
