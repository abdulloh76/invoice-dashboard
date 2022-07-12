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
			"message": "unexpected error occured: " + err.Error(),
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
			"message": "unexpected error occured: " + err.Error(),
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
			"message": "invoice wiht given id not found: " + err.Error(),
		})
		return
	}

	invoiceDto := invoiceDto.EntitytoResponsetDTO(&invoice)
	context.JSON(200, invoiceDto)
}

func UpdateInvoice(context *gin.Context) {
	invoiceId := context.Param("invoiceId")

	curInvoice, notFoundErr := FindInvoiceById(invoiceId)
	if notFoundErr != nil {
		context.AbortWithStatusJSON(404, map[string]string{
			"message": "invoice wiht given id not found: " + notFoundErr.Error(),
		})
		return
	}

	var invoice invoiceDto.PutInvoiceBody
	if context.BindJSON(&invoice) != nil {
		context.AbortWithStatusJSON(400, map[string]string{
			"message": "couldn't parse given body",
		})
		return
	}

	err := ModifyInvoice(&curInvoice, invoice)
	if err != nil {
		context.AbortWithStatusJSON(500, map[string]string{
			"message": "can't modify invoice object: " + err.Error(),
		})
		return
	}

	// ? is ignoring the err ok for this expression
	updatedInvoice, _ := FindInvoiceById(invoiceId)
	invoiceDto := invoiceDto.EntitytoResponsetDTO(&updatedInvoice)

	context.JSON(200, invoiceDto)
}

func DeleteInvoice(context *gin.Context) {
	invoiceId := context.Param("invoiceId")

	// !shordid ususally returns string with length 9||10
	if len(invoiceId) != 9 && len(invoiceId) != 10 {
		context.AbortWithStatusJSON(400, map[string]string{
			"message": "id is not valid",
		})
		return
	}

	err := RemoveInvoice(invoiceId)
	if err != nil {
		context.AbortWithStatusJSON(404, map[string]string{
			"message": "invoice wiht given id not found: " + err.Error(),
		})
		return
	}

	context.JSON(201, map[string]string{
		"message": "successfully deleted",
	})
}
