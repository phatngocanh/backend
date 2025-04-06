package v1

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/phat-ngoc-anh/backend/internal/domain/model"
	"github.com/phat-ngoc-anh/backend/internal/service"
)

type InvoiceHandler struct {
	invoiceService service.InvoiceService
}

func NewInvoiceHandler(invoiceService service.InvoiceService) *InvoiceHandler {
	return &InvoiceHandler{invoiceService: invoiceService}
}

// @BasePath /api/v1
// @Summary Generate invoice PDF
// @Description Generate a PDF invoice based on the provided data
// @Tags Invoice
// @Accept json
// @Produce application/pdf
// @Param request body model.GenerateInvoiceRequest true "Invoice data"
// @Router /api/v1/invoice/generate [post]
// @Success 200 {file} binary "PDF file"
func (handler *InvoiceHandler) GenerateInvoice(c *gin.Context) {
	var request model.GenerateInvoiceRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		log.Printf("Error binding JSON: %v", err)
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	pdf, err := handler.invoiceService.GenerateInvoicePDF(request)
	if err != nil {
		log.Printf("Error generating PDF: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate PDF"})
		return
	}

	filename := fmt.Sprintf("invoice-%s.pdf", time.Now().Format("2006-01-02T15-04-05"))
	c.Header("Content-Disposition", fmt.Sprintf("attachment; filename=%s", filename))
	c.Data(http.StatusOK, "application/pdf", pdf)
}
