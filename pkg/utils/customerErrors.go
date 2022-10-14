package utils

import "errors"

var (
	ErrJsonUnmarshal   = errors.New("failed to parse invoice from request body")
	ErrInvoiceNotFound = errors.New("invoice with given ID not found")
)
