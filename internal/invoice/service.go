package invoice

import (
	"invoice-dashboard/internal/entity"

	"github.com/gin-gonic/gin"
)

type Response struct {
	message string
}

var invoices []entity.Invoice = []entity.Invoice{
	{ID: "qwe123", Name: "tax for water", Price: 45},
	{ID: "asd123", Name: "tax for gas", Price: 55},
	{ID: "zxc123", Name: "tax for electricity", Price: 60},
	{ID: "wer123", Name: "college payment", Price: 354},
}

func CreateInvoice(context *gin.Context) {
	var newInvoice entity.Invoice

	if err := context.BindJSON(&newInvoice); err != nil {
		context.JSON(400, Response{message: "couldn't parse given body"})
		return
	}

	invoices = append(invoices, newInvoice)
	respose := Response{message: "successfully added"}
	context.JSON(200, respose)
}

func GetAll(context *gin.Context) {
	context.JSON(200, invoices)
}

func GetById(context *gin.Context) {
	id := context.Param("invoiceId")
	for _, invoice := range invoices {
		if invoice.ID == id {
			context.JSON(200, invoice)
			return
		}
	}
	respose := Response{message: "no invoice found with given id: " + id}
	context.JSON(404, respose)
}

func EditInvoice(context *gin.Context) {
	var invoiceObj entity.Invoice
	if err := context.BindJSON(&invoiceObj); err != nil {
		response := Response{message: "couldn't parse given body"}
		context.JSON(400, response)
		return
	}

	id := context.Param("invoiceId")
	for i, invoice := range invoices {
		if invoice.ID == id {
			invoices[i] = invoiceObj
			context.JSON(200, invoiceObj)
			return
		}
	}
	response := Response{message: "no invoice found with given id: " + id}
	context.JSON(404, response)
}

func DeleteInvoice(context *gin.Context) {
	id := context.Param("invoiceId")
	for i, invoice := range invoices {
		if invoice.ID == id {
			invoices = append(invoices[:i], invoices[i+1:]...)
			response := Response{message: "successfully deleted"}
			context.JSON(200, response)
			return
		}
	}
	response := Response{message: "no invoice found with given id: " + id}
	context.JSON(404, response)
}
