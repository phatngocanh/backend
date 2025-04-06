package serviceimplement

import (
	"bytes"
	"fmt"
	"log"
	"path/filepath"
	"time"

	"github.com/jung-kurt/gofpdf"
	"github.com/phat-ngoc-anh/backend/internal/domain/model"
	"github.com/phat-ngoc-anh/backend/internal/service"
)

type invoiceService struct {
	companyName    string
	companyAddress string
	companyPhone   string
}

func NewInvoiceService() service.InvoiceService {
	return &invoiceService{
		companyName:    "CÔNG TY TNHH HÓA PHẨM PHÁT NGỌC ANH",
		companyAddress: "Địa chỉ: 430/33 đường TA28, KP 2, P. Thới An, Quận 12, TP.HCM",
		companyPhone:   "Số điện thoại: 0982868226",
	}
}

func (s *invoiceService) GenerateInvoicePDF(request model.GenerateInvoiceRequest) ([]byte, error) {
	// Initialize PDF with A5 size (148 x 210 mm)
	pdf := gofpdf.New("P", "mm", "A5", "")
	pdf.SetMargins(10, 10, 10) // Set margins to 10mm on all sides
	pdf.AddPage()

	// Add font for Vietnamese
	fontPath := filepath.Join("internal", "asset", "DejaVuSans.ttf")
	pdf.AddUTF8Font("DejaVu", "", fontPath)
	pdf.AddUTF8Font("DejaVu", "B", filepath.Join("internal", "asset", "DejaVuSans-Bold.ttf"))

	// Add logo
	logoPath := filepath.Join("internal", "asset", "brand.png")
	pdf.Image(logoPath, 10, 10, 15, 0, false, "", 0, "") // Reduced logo size to 15mm width

	// Company information - positioned to the right of the logo
	pdf.SetFont("DejaVu", "B", 11)
	pdf.SetXY(30, 10)
	pdf.CellFormat(0, 5, s.companyName, "", 1, "L", false, 0, "")
	pdf.SetFont("DejaVu", "", 8)
	pdf.SetX(30)
	pdf.CellFormat(0, 4, s.companyAddress, "", 1, "L", false, 0, "")
	pdf.SetX(30)
	pdf.CellFormat(0, 4, s.companyPhone, "", 1, "L", false, 0, "")

	// Title
	pdf.Ln(5)
	pdf.SetFont("DejaVu", "B", 13)
	pdf.CellFormat(0, 6, "PHIẾU XUẤT KHO", "", 1, "C", false, 0, "")

	// Invoice Code
	pdf.Ln(2)
	pdf.SetFont("DejaVu", "", 8)
	pdf.CellFormat(0, 1, fmt.Sprintf("Mã đơn: %s", request.InvoiceCode), "", 1, "C", false, 0, "")

	// Date and Customer information
	pdf.Ln(3)
	pdf.SetFont("DejaVu", "", 9)
	// Parse the date
	date, err := time.Parse("02-01-2006", request.InvoiceDate)
	if err != nil {
		return nil, fmt.Errorf("invalid date format: %v", err)
	}

	// Format the date for display
	formattedDate := date.Format("02-01-2006")
	pdf.CellFormat(0, 5, fmt.Sprintf("Ngày: %s", formattedDate), "", 1, "L", false, 0, "")
	pdf.CellFormat(0, 5, fmt.Sprintf("Tên khách hàng: %s", request.CustomerName), "", 1, "L", false, 0, "")
	if request.CustomerPhone != "" {
		pdf.CellFormat(0, 5, fmt.Sprintf("SĐT: %s", request.CustomerPhone), "", 1, "L", false, 0, "")
	}
	pdf.Ln(2)

	// Calculate optimal column widths for A5 paper
	pageWidth := 148.0 // A5 width
	margins := 20.0    // Total horizontal margins
	availableWidth := pageWidth - margins

	// Define column widths as percentages of available width
	colWidths := []float64{
		availableWidth * 0.08, // STT (8%)
		availableWidth * 0.42, // Tên hàng hóa (42%)
		availableWidth * 0.17, // Số thùng (17%)
		availableWidth * 0.17, // ĐV/thùng (17%)
		availableWidth * 0.16, // Tổng ĐV (16%)
	}

	// Table header
	pdf.SetFont("DejaVu", "B", 8)
	pdf.SetFillColor(240, 240, 240)
	headers := []string{"STT", "Tên hàng hóa", "Số thùng", "ĐV/thùng", "Tổng ĐV"}
	aligns := []string{"C", "C", "C", "C", "C"}

	for i, header := range headers {
		pdf.CellFormat(colWidths[i], 6, header, "1", 0, aligns[i], true, 0, "")
	}
	pdf.Ln(-1)

	// Table content
	pdf.SetFont("DejaVu", "", 8)
	rowHeight := 5.5 // Reduced row height for more compact layout

	for i, item := range request.Items {
		// STT
		pdf.CellFormat(colWidths[0], rowHeight, fmt.Sprintf("%d", i+1), "1", 0, "C", false, 0, "")
		// Tên hàng hóa
		var packagesStr, itemsPerPackageStr string
		if item.Packages != nil {
			packagesStr = fmt.Sprintf("%d", *item.Packages)
		}
		if item.ItemsPerPackage != nil {
			itemsPerPackageStr = fmt.Sprintf("%d", *item.ItemsPerPackage)
		}
		pdf.CellFormat(colWidths[1], rowHeight, item.Name, "1", 0, "L", false, 0, "")
		// Số thùng
		pdf.CellFormat(colWidths[2], rowHeight, packagesStr, "1", 0, "C", false, 0, "")
		// ĐV/thùng
		pdf.CellFormat(colWidths[3], rowHeight, itemsPerPackageStr, "1", 0, "C", false, 0, "")
		// Tổng ĐV
		pdf.CellFormat(colWidths[4], rowHeight, fmt.Sprintf("%d", item.TotalUnits), "1", 1, "C", false, 0, "")
	}

	// Total row - use the values directly from the request
	pdf.SetFont("DejaVu", "B", 8)
	pdf.CellFormat(colWidths[0], 6, "", "1", 0, "C", false, 0, "")
	pdf.CellFormat(colWidths[1], 6, "TỔNG CỘNG:", "1", 0, "R", false, 0, "")
	pdf.CellFormat(colWidths[2], 6, fmt.Sprintf("%d", request.TotalPackages), "1", 0, "C", false, 0, "")
	pdf.CellFormat(colWidths[3], 6, "", "1", 0, "C", false, 0, "")
	pdf.CellFormat(colWidths[4], 6, fmt.Sprintf("%d", request.TotalUnits), "1", 1, "C", false, 0, "")

	// Return PDF as bytes
	var buf bytes.Buffer
	err = pdf.Output(&buf)
	if err != nil {
		log.Printf("Error generating PDF output: %v", err)
		return nil, err
	}
	return buf.Bytes(), nil
}
