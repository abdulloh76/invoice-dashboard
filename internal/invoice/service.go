package invoice

import (
	"invoice-dashboard/internal/dto/invoiceDto"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateInvoice(context *gin.Context) {
	var invoice invoiceDto.InvoiceRequestBody
	if context.BindJSON(&invoice) != nil {
		context.AbortWithStatusJSON(400, gin.H{
			"status":  400,
			"message": "couldn't parse given body",
		})
		return
	}

	newInvoice := invoiceDto.RequestDTOtoEntity(&invoice)
	err := InsertInvoice(&newInvoice)
	if err != nil {
		context.AbortWithStatusJSON(500, gin.H{
			"status":  500,
			"message": "unexpected error occured",
		})
		return
	}

	context.JSON(200, gin.H{
		"status":  200,
		"message": "successfully added",
	})
}

func GetAll(context *gin.Context) {
	invoices, err := FindInvoices()

	if err != nil {
		context.AbortWithStatusJSON(500, gin.H{
			"status":  500,
			"message": "unexpected error occured",
		})
		return
	}

	context.JSON(200, gin.H{
		"status": 200,
		"data":   invoices,
	})
}

func GetById(context *gin.Context) {
	invoiceId := context.Param("invoiceId")

	id, err := strconv.ParseUint(invoiceId, 10, 64)
	if err != nil {
		context.AbortWithStatusJSON(400, gin.H{
			"status":  400,
			"message": "id is not valid",
		})
		return
	}

	invoice, err := FindInvoiceById(id)
	if err != nil {
		context.AbortWithStatusJSON(404, gin.H{
			"status":  404,
			"message": "invoice with given id not found",
		})
		return
	}

	invoiceDto := invoiceDto.EntitytoResponsetDTO(&invoice)
	context.JSON(200, gin.H{
		"status": 200,
		"data":   invoiceDto,
	})
}

func EditInvoice(context *gin.Context) {}

func DeleteInvoice(context *gin.Context) {}
