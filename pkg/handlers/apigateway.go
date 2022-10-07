package handlers

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"github.com/abdulloh76/invoice-dashboard/pkg/domain"
	"github.com/abdulloh76/invoice-dashboard/pkg/types"
	"github.com/aws/aws-lambda-go/events"
)

type APIGatewayHandler struct {
	invoices *domain.Invoices
}

func NewAPIGatewayHandler(d *domain.Invoices) *APIGatewayHandler {
	return &APIGatewayHandler{
		invoices: d,
	}
}

func (g *APIGatewayHandler) AllHandler(event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	allInvoices, err := g.invoices.AllInvoices()
	if err != nil {
		return errResponse(http.StatusInternalServerError, err.Error()), nil
	}

	return response(http.StatusOK, allInvoices), nil
}

func (g *APIGatewayHandler) GetHandler(ctx context.Context, event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	id, ok := event.PathParameters["id"]
	if !ok {
		return errResponse(http.StatusBadRequest, "missing 'id' parameter in path"), nil
	}

	invoice, err := g.invoices.GetSingleInvoice(id)

	if errors.Is(err, domain.ErrInvoiceNotFound) {
		return errResponse(http.StatusInternalServerError, err.Error()), nil
	}
	if err != nil {
		return errResponse(http.StatusInternalServerError, err.Error()), nil
	}

	// todo will be updated when gRPC client implemented
	senderAddress := &types.GetAddressDto{}

	invoiceDto := types.EntityToResponseDTO(invoice, senderAddress)
	return response(http.StatusOK, invoiceDto), nil
}

func (g *APIGatewayHandler) CreateHandler(event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	if strings.TrimSpace(event.Body) == "" {
		return errResponse(http.StatusBadRequest, "empty request body"), nil
	}

	newInvoice, err := g.invoices.Create([]byte(event.Body))
	if errors.Is(err, domain.ErrJsonUnmarshal) {
		return errResponse(http.StatusInternalServerError, err.Error()), nil
	}
	if err != nil {
		return errResponse(http.StatusInternalServerError, err.Error()), nil
	}

	// todo will be updated when gRPC client implemented
	senderAddress := &types.GetAddressDto{}

	invoiceDto := types.EntityToResponseDTO(newInvoice, senderAddress)
	return response(http.StatusOK, invoiceDto), nil
}

func (g *APIGatewayHandler) PutHandler(event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	id, ok := event.PathParameters["id"]
	if !ok {
		return errResponse(http.StatusBadRequest, "missing 'id' parameter in path"), nil
	}

	if strings.TrimSpace(event.Body) == "" {
		return errResponse(http.StatusBadRequest, "empty request body"), nil
	}

	updatedInvoice, err := g.invoices.ModifyInvoice(id, []byte(event.Body))
	if errors.Is(err, domain.ErrInvoiceNotFound) {
		return errResponse(http.StatusInternalServerError, err.Error()), nil
	}
	if errors.Is(err, domain.ErrJsonUnmarshal) {
		return errResponse(http.StatusInternalServerError, err.Error()), nil
	}
	if err != nil {
		return errResponse(http.StatusInternalServerError, err.Error()), nil
	}

	// todo will be updated when gRPC client implemented
	senderAddress := &types.GetAddressDto{}

	invoice := types.EntityToResponseDTO(updatedInvoice, senderAddress)
	return response(http.StatusOK, invoice), nil
}

func (g *APIGatewayHandler) DeleteHandler(event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	id, ok := event.PathParameters["id"]
	if !ok {
		return errResponse(http.StatusBadRequest, "missing 'id' parameter in path"), nil
	}

	err := g.invoices.DeleteInvoice(id)
	if errors.Is(err, domain.ErrInvoiceNotFound) {
		return errResponse(http.StatusInternalServerError, err.Error()), nil
	}
	if err != nil {
		return errResponse(http.StatusInternalServerError, err.Error()), nil
	}

	return response(http.StatusOK, nil), nil
}

func response(code int, object interface{}) events.APIGatewayProxyResponse {
	marshalled, err := json.Marshal(object)
	if err != nil {
		return errResponse(http.StatusInternalServerError, err.Error())
	}

	return events.APIGatewayProxyResponse{
		StatusCode: code,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		Body:            string(marshalled),
		IsBase64Encoded: false,
	}
}

func errResponse(status int, body string) events.APIGatewayProxyResponse {
	message := map[string]string{
		"message": body,
	}

	messageBytes, _ := json.Marshal(&message)

	return events.APIGatewayProxyResponse{
		StatusCode: status,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		Body: string(messageBytes),
	}
}
