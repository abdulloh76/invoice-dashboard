package invoice

import (
	"invoice-dashboard/internal/entity"

	"github.com/gin-gonic/gin"
)

func CreateInvoice(context *gin.Context) {
	var newInvoice entity.Invoice

	if context.BindJSON(&newInvoice) == nil {
		context.JSON(400, map[string]string{
			"message": "couldn't parse given body",
		})
	}

	InsertInvoice(&newInvoice)
	context.JSON(200, map[string]string{
		"message": "successfully added",
	})
}

func GetAll(context *gin.Context) {
	invoices := FindInvoices()
	context.JSON(200, invoices)
}

func GetById(context *gin.Context) {}

func EditInvoice(context *gin.Context) {}

func DeleteInvoice(context *gin.Context) {}
