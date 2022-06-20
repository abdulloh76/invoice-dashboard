package invoice

import (
	"invoice-dashboard/internal/dto/invoiceDto"

	"github.com/gin-gonic/gin"
)

func CreateInvoice(context *gin.Context) {
	var invoice invoiceDto.InvoiceRequestBody
	if context.BindJSON(&invoice) != nil {
		context.AbortWithStatusJSON(400, map[string]string{
			"message": "couldn't parse given body",
		})
		return
	}

	newInvoice := invoiceDto.RequestDTOtoEntity(&invoice)
	err := InsertInvoice(&newInvoice)
	if err != nil {
		context.AbortWithStatusJSON(500, map[string]string{
			"message": "unexpected error occured",
		})
		return
	}

	context.AbortWithStatusJSON(200, map[string]string{
		"message": "successfully added",
	})
}

func GetAll(context *gin.Context) {
	invoices, err := FindInvoices()

	if err != nil {
		context.AbortWithStatusJSON(500, map[string]string{
			"message": "unexpected error occured",
		})
		return
	}

	context.JSON(200, invoices)
}

func GetById(context *gin.Context) {
	invoiceId := context.Param("invoiceId")

	// shordid ususally returns string with length 9||10
	if len(invoiceId) != 9 && len(invoiceId) != 10 {
		context.AbortWithStatusJSON(400, map[string]string{
			"message": "id is not valid",
		})
		return
	}

	invoice, err := FindInvoiceById(invoiceId)
	if err != nil {
		context.AbortWithStatusJSON(404, map[string]string{
			"message": "invoice wiht given id not found",
		})
		return
	}

	invoiceDto := invoiceDto.EntitytoResponsetDTO(&invoice)
	context.JSON(200, invoiceDto)
}

func EditInvoice(context *gin.Context) {}

func DeleteInvoice(context *gin.Context) {
	invoiceId := context.Param("invoiceId")

	// shordid ususally returns string with length 9||10
	if len(invoiceId) != 9 && len(invoiceId) != 10 {
		context.AbortWithStatusJSON(400, map[string]string{
			"message": "id is not valid",
		})
		return
	}

	err := RemoveInvoice(invoiceId)
	if err != nil {
		context.AbortWithStatusJSON(404, map[string]string{
			"message": "invoice wiht given id not found",
		})
		return
	}

	context.JSON(201, map[string]string{
		"message": "successfully deleted",
	})
}
