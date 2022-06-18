package invoice

import (
	"invoice-dashboard/internal/dto/invoiceDto"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateInvoice(context *gin.Context) {
	var invoice invoiceDto.InvoiceRequestBody
	if context.BindJSON(&invoice) != nil {
		context.AbortWithStatusJSON(400, map[string]string{
			"status":  "400",
			"message": "couldn't parse given body",
		})
		return
	}

	newInvoice := invoiceDto.RequestDTOtoEntity(&invoice)
	InsertInvoice(&newInvoice)
	context.JSON(200, map[string]string{
		"status":  "200",
		"message": "successfully added",
	})
}

func GetAll(context *gin.Context) {
	invoices := FindInvoices()
	context.JSON(200, invoices)
}

func GetById(context *gin.Context) {
	invoiceId := context.Param("invoiceId")

	id, err := strconv.ParseUint(invoiceId, 10, 64)
	if err != nil {
		context.AbortWithStatusJSON(400, map[string]string{
			"status":  "400",
			"message": "id is not valid",
		})
		return
	}

	invoice := FindInvoiceById(id)
	// need to check for 404
	invoiceDto := invoiceDto.EntitytoResponsetDTO(&invoice)

	context.JSON(200, invoiceDto)
}

func EditInvoice(context *gin.Context) {}

func DeleteInvoice(context *gin.Context) {}
