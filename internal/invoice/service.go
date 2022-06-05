package invoice

import (
	"invoice-dashboard/internal/entity"

	"github.com/gin-gonic/gin"
)

type Response struct {
	message string
}

func CreateInvoice(context *gin.Context) {
	var newInvoice entity.Invoice

	if context.BindJSON(&newInvoice) == nil {
		context.JSON(400, Response{message: "couldn't parse given body"})
		return
	}

	InsertInvoice(&newInvoice)
	respose := Response{message: "successfully added"}
	context.JSON(200, respose)
}

func GetAll(context *gin.Context) {
	invoices := FindInvoices()
	context.JSON(200, invoices)
}

func GetById(context *gin.Context) {}

func EditInvoice(context *gin.Context) {}

func DeleteInvoice(context *gin.Context) {}
