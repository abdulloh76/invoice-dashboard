package handlers

import (
	"errors"
	"io"
	"log"
	"net/http"

	"github.com/abdulloh76/invoice-dashboard/domain"
	"github.com/abdulloh76/invoice-dashboard/types"
	"github.com/gin-gonic/gin"
)

type GinAPIHandler struct {
	invoices *domain.Invoices
}

func NewGinAPIHandler(d *domain.Invoices) *GinAPIHandler {
	return &GinAPIHandler{
		invoices: d,
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

	if errors.Is(err, domain.ErrInvoiceNotFound) {
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

	invoiceDto := types.EntitytoResponsetDTO(invoice)
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
	if errors.Is(err, domain.ErrJsonUnmarshal) {
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

	invoiceDto := types.EntitytoResponsetDTO(newInvoice)
	context.JSON(http.StatusOK, invoiceDto)
}

func (g *GinAPIHandler) PutHandler(context *gin.Context) {
	id := context.Param("id")

	body, err := io.ReadAll(context.Request.Body)
	if err != nil {
		log.Fatal(err)
	}

	updatedInvoice, err := g.invoices.ModifyInvoice(id, body)
	if errors.Is(err, domain.ErrInvoiceNotFound) {
		context.AbortWithStatusJSON(http.StatusNotFound, map[string]string{
			"message": err.Error(),
		})
		return
	}
	if errors.Is(err, domain.ErrJsonUnmarshal) {
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

	invoice := types.EntitytoResponsetDTO(updatedInvoice)
	context.JSON(http.StatusOK, invoice)
}

func (g *GinAPIHandler) DeleteHandler(context *gin.Context) {
	id := context.Param("id")

	err := g.invoices.DeleteInvoice(id)
	if errors.Is(err, domain.ErrInvoiceNotFound) {
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