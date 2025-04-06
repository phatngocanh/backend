package v1

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func MapRoutes(router *gin.Engine, studentHandler *StudentHandler, invoiceHandler *InvoiceHandler) {
	v1 := router.Group("/api/v1")
	{
		students := v1.Group("/students")
		{
			students.GET("/", studentHandler.GetAll)
		}

		invoice := v1.Group("/invoice")
		{
			invoice.POST("/generate", invoiceHandler.GenerateInvoice)
		}
	}
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
