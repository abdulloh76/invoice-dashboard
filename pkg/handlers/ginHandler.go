package handlers

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/abdulloh76/invoice-dashboard/pkg/domain"
	"github.com/abdulloh76/invoice-dashboard/pkg/infrastructure"
	"github.com/abdulloh76/invoice-dashboard/pkg/types"
	"github.com/abdulloh76/invoice-dashboard/pkg/utils"

	"github.com/awslabs/aws-lambda-go-api-proxy/core"
	"github.com/gin-gonic/gin"
)

type GinAPIHandler struct {
	invoices       *domain.Invoices
	userGrpcClient *infrastructure.UserGrpcClient
}

func NewGinAPIHandler(d *domain.Invoices, userGrpcClient *infrastructure.UserGrpcClient) *GinAPIHandler {
	return &GinAPIHandler{
		invoices:       d,
		userGrpcClient: userGrpcClient,
	}
}

func RegisterHandlers(router *gin.Engine, handler *GinAPIHandler) {
	router.POST("/invoice", handler.CreateHandler)
	router.GET("/invoice", handler.AllHandler)
	router.GET("/invoice/:id", handler.GetHandler)
	router.PUT("/invoice/:id", handler.PutHandler)
	router.DELETE("/invoice/:id", handler.DeleteHandler)
}

func (g *GinAPIHandler) AllHandler(context *gin.Context) {
	apiGWRequestContext, ok := core.GetAPIGatewayV2ContextFromContext(context.Request.Context())
	fmt.Println("context request context", context.Request.Context())
	fmt.Println("apiGWRequestContext.Authorizer.Lambda", apiGWRequestContext.Authorizer.Lambda)
	fmt.Println("apiGWRequestContext.Authorizer.IAM", apiGWRequestContext.Authorizer.IAM)
	fmt.Println("apiGWRequestContext.Authorizer.IAM.UserID", apiGWRequestContext.Authorizer.IAM.UserID)
	fmt.Println("ok got with core aws lambda package", ok)
	allInvoices, err := g.invoices.AllInvoices()
	if err != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
		return
	}

	context.JSON(http.StatusOK, allInvoices)
}

func (g *GinAPIHandler) GetHandler(context *gin.Context) {
	id := context.Param("id")

	invoice, err := g.invoices.GetSingleInvoice(id)

	if errors.Is(err, utils.ErrInvoiceNotFound) {
		context.AbortWithStatusJSON(http.StatusNotFound, map[string]string{
			"message": err.Error(),
		})
		return
	}
	if err != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
		return
	}

	// todo will be updated when getting user_id configured
	senderAddress, err := g.userGrpcClient.GetUserAddress(context, "mHVxHT4VR")
	if err != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
		return
	}

	invoiceDto := types.EntityToResponseDTO(invoice, senderAddress)
	context.JSON(http.StatusOK, invoiceDto)
}

func (g *GinAPIHandler) CreateHandler(context *gin.Context) {
	body, err := io.ReadAll(context.Request.Body)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
		return
	}

	newInvoice, err := g.invoices.Create(body)
	if errors.Is(err, utils.ErrJsonUnmarshal) {
		context.AbortWithStatusJSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
		})
		return
	}
	if err != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
		return
	}

	// todo will be updated when getting user_id configured
	senderAddress, err := g.userGrpcClient.GetUserAddress(context, "mHVxHT4VR")
	if err != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
		return
	}

	invoiceDto := types.EntityToResponseDTO(newInvoice, senderAddress)
	context.JSON(http.StatusOK, invoiceDto)
}

func (g *GinAPIHandler) PutHandler(context *gin.Context) {
	id := context.Param("id")

	body, err := io.ReadAll(context.Request.Body)
	if err != nil {
		log.Fatal(err)
	}

	updatedInvoice, err := g.invoices.ModifyInvoice(id, body)
	if errors.Is(err, utils.ErrInvoiceNotFound) {
		context.AbortWithStatusJSON(http.StatusNotFound, map[string]string{
			"message": err.Error(),
		})
		return
	}
	if errors.Is(err, utils.ErrJsonUnmarshal) {
		context.AbortWithStatusJSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
		})
		return
	}
	if err != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
		return
	}

	// todo will be updated when getting user_id configured
	senderAddress, err := g.userGrpcClient.GetUserAddress(context, "mHVxHT4VR")
	if err != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
		return
	}

	invoice := types.EntityToResponseDTO(updatedInvoice, senderAddress)
	context.JSON(http.StatusOK, invoice)
}

func (g *GinAPIHandler) DeleteHandler(context *gin.Context) {
	id := context.Param("id")

	err := g.invoices.DeleteInvoice(id)
	if errors.Is(err, utils.ErrInvoiceNotFound) {
		context.AbortWithStatusJSON(http.StatusNotFound, map[string]string{
			"message": err.Error(),
		})
		return
	}
	if err != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
		return
	}

	context.JSON(http.StatusOK, map[string]string{
		"message": "successfully deleted",
	})
}
