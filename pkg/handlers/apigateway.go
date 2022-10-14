package handlers

import (
	"context"
	"errors"
	"github.com/abdulloh76/invoice-dashboard/pkg/utils"
	"net/http"
	"strings"

	"github.com/abdulloh76/invoice-dashboard/pkg/domain"
	"github.com/abdulloh76/invoice-dashboard/pkg/infrastructure"
	"github.com/abdulloh76/invoice-dashboard/pkg/types"
	"github.com/aws/aws-lambda-go/events"
)

type APIGatewayHandler struct {
	invoices       *domain.Invoices
	userGrpcClient *infrastructure.UserGrpcClient
}

func NewAPIGatewayHandler(d *domain.Invoices, userGrpcClient *infrastructure.UserGrpcClient) *APIGatewayHandler {
	return &APIGatewayHandler{
		invoices:       d,
		userGrpcClient: userGrpcClient,
	}
}

func (g *APIGatewayHandler) AllHandler(ctx context.Context, event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	allInvoices, err := g.invoices.AllInvoices()
	if err != nil {
		return utils.ErrResponse(http.StatusInternalServerError, err.Error()), nil
	}

	return utils.Response(http.StatusOK, allInvoices), nil
}

func (g *APIGatewayHandler) GetHandler(ctx context.Context, event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	id, ok := event.PathParameters["id"]
	if !ok {
		return utils.ErrResponse(http.StatusBadRequest, "missing 'id' parameter in path"), nil
	}

	invoice, err := g.invoices.GetSingleInvoice(id)

	if errors.Is(err, utils.ErrInvoiceNotFound) {
		return utils.ErrResponse(http.StatusInternalServerError, err.Error()), nil
	}
	if err != nil {
		return utils.ErrResponse(http.StatusInternalServerError, err.Error()), nil
	}

	// todo will be updated when getting user_id configured
	senderAddress, err := g.userGrpcClient.GetUserAddress(ctx, "mHVxHT4VR")
	if err != nil {
		return utils.ErrResponse(http.StatusInternalServerError, err.Error()), nil
	}

	invoiceDto := types.EntityToResponseDTO(invoice, senderAddress)
	return utils.Response(http.StatusOK, invoiceDto), nil
}

func (g *APIGatewayHandler) CreateHandler(ctx context.Context, event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	if strings.TrimSpace(event.Body) == "" {
		return utils.ErrResponse(http.StatusBadRequest, "empty request body"), nil
	}

	newInvoice, err := g.invoices.Create([]byte(event.Body))
	if errors.Is(err, utils.ErrJsonUnmarshal) {
		return utils.ErrResponse(http.StatusInternalServerError, err.Error()), nil
	}
	if err != nil {
		return utils.ErrResponse(http.StatusInternalServerError, err.Error()), nil
	}

	// todo will be updated when getting user_id configured
	senderAddress, err := g.userGrpcClient.GetUserAddress(ctx, "mHVxHT4VR")
	if err != nil {
		return utils.ErrResponse(http.StatusInternalServerError, err.Error()), nil
	}

	invoiceDto := types.EntityToResponseDTO(newInvoice, senderAddress)
	return utils.Response(http.StatusOK, invoiceDto), nil
}

func (g *APIGatewayHandler) PutHandler(ctx context.Context, event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	id, ok := event.PathParameters["id"]
	if !ok {
		return utils.ErrResponse(http.StatusBadRequest, "missing 'id' parameter in path"), nil
	}

	if strings.TrimSpace(event.Body) == "" {
		return utils.ErrResponse(http.StatusBadRequest, "empty request body"), nil
	}

	updatedInvoice, err := g.invoices.ModifyInvoice(id, []byte(event.Body))
	if errors.Is(err, utils.ErrInvoiceNotFound) {
		return utils.ErrResponse(http.StatusInternalServerError, err.Error()), nil
	}
	if errors.Is(err, utils.ErrJsonUnmarshal) {
		return utils.ErrResponse(http.StatusInternalServerError, err.Error()), nil
	}
	if err != nil {
		return utils.ErrResponse(http.StatusInternalServerError, err.Error()), nil
	}

	// todo will be updated when getting user_id configured
	senderAddress, err := g.userGrpcClient.GetUserAddress(ctx, "mHVxHT4VR")
	if err != nil {
		return utils.ErrResponse(http.StatusInternalServerError, err.Error()), nil
	}

	invoice := types.EntityToResponseDTO(updatedInvoice, senderAddress)
	return utils.Response(http.StatusOK, invoice), nil
}

func (g *APIGatewayHandler) DeleteHandler(ctx context.Context, event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	id, ok := event.PathParameters["id"]
	if !ok {
		return utils.ErrResponse(http.StatusBadRequest, "missing 'id' parameter in path"), nil
	}

	err := g.invoices.DeleteInvoice(id)
	if errors.Is(err, utils.ErrInvoiceNotFound) {
		return utils.ErrResponse(http.StatusInternalServerError, err.Error()), nil
	}
	if err != nil {
		return utils.ErrResponse(http.StatusInternalServerError, err.Error()), nil
	}

	return utils.Response(http.StatusOK, nil), nil
}
