package handlers

import (
	"errors"
	"io"
	"log"
	"net/http"

	"github.com/abdulloh76/invoice-dashboard/domain"
	"github.com/abdulloh76/invoice-dashboard/internal/config"
	"github.com/abdulloh76/invoice-dashboard/store"
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

func RegisterHandlers(router *gin.Engine) {
	postgresDSN := config.Configs.POSTGRES_URI
	postgreDB := store.NewPostgresDBStore(postgresDSN)
	domain := domain.NewInvoicesDomain(postgreDB)
	handler := NewGinAPIHandler(domain)

	router.POST("/invoice", handler.CreateHandler)
	router.GET("/invoice", handler.AllHandler)
	router.GET("/invoice/:id", handler.GetHandler)
	router.PUT("/invoice/:id", handler.PutHandler)
	router.DELETE("/invoice/:id", handler.DeleteHandler)
}

func (g *GinAPIHandler) AllHandler(context *gin.Context) {
	allInvoices, err := g.invoices.AllInvoices()
	if err != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	context.JSON(http.StatusOK, allInvoices)
}

func (g *GinAPIHandler) GetHandler(context *gin.Context) {
	id := context.Param("id")

	invoice, err := g.invoices.GetSingleInvoice(id)

	if errors.Is(err, domain.ErrUserNotFound) {
		context.AbortWithStatusJSON(http.StatusNotFound, err)
		return
	}
	if err != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}
	context.JSON(http.StatusOK, invoice)
}

func (g *GinAPIHandler) CreateHandler(context *gin.Context) {
	body, err := io.ReadAll(context.Request.Body)
	if err != nil {
		log.Fatal(err)
	}

	newInvoice, err := g.invoices.Create(body)
	if errors.Is(err, domain.ErrJsonUnmarshal) {
		context.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}
	if err != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	context.AbortWithStatusJSON(http.StatusOK, newInvoice)
}

func (g *GinAPIHandler) PutHandler(context *gin.Context) {
	id := context.Param("id")

	body, err := io.ReadAll(context.Request.Body)
	if err != nil {
		log.Fatal(err)
	}

	updatedInvoice, err := g.invoices.ModifyInvoice(id, body)
	if errors.Is(err, domain.ErrJsonUnmarshal) {
		context.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}
	if err != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	invoice := types.EntitytoResponsetDTO(updatedInvoice)
	context.JSON(http.StatusOK, invoice)
}

func (g *GinAPIHandler) DeleteHandler(context *gin.Context) {
	id := context.Param("id")

	err := g.invoices.DeleteInvoice(id)
	if errors.Is(err, domain.ErrUserNotFound) {
		context.AbortWithStatusJSON(http.StatusNotFound, err)
		return
	}
	if err != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	context.JSON(http.StatusOK, map[string]string{
		"message": "successfully deleted",
	})
}
