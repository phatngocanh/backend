package service

import "github.com/phat-ngoc-anh/backend/internal/domain/model"

type InvoiceService interface {
	GenerateInvoicePDF(request model.GenerateInvoiceRequest) ([]byte, error)
}
