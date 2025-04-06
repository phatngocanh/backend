package model

type InvoiceItem struct {
	Name            string `json:"name"`
	Packages        *int   `json:"packages"`
	ItemsPerPackage *int   `json:"itemsPerPackage"`
	TotalUnits      int    `json:"totalUnits"`
}

type GenerateInvoiceRequest struct {
	CustomerName    string        `json:"customerName"`
	CustomerPhone   string        `json:"customerPhone"`
	CustomerAddress string        `json:"customerAddress"`
	InvoiceDate     string        `json:"invoiceDate"`
	InvoiceCode     string        `json:"invoiceCode"`
	Items           []InvoiceItem `json:"items"`
	TotalPackages   int           `json:"totalPackages"`
	TotalUnits      int           `json:"totalUnits"`
}
