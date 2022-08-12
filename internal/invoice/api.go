package invoice

import (
	"github.com/gin-gonic/gin"
)

func RegisterHandlers(router *gin.Engine) {
	router.POST("/invoice", CreateInvoice)
	router.GET("/invoice", GetAll)
	router.GET("/invoice/:invoiceId", GetById)
	router.PUT("/invoice/:invoiceId", UpdateInvoice)
	router.DELETE("/invoice/:invoiceId", DeleteInvoice)
}
