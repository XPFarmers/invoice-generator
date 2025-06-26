package invoices

import "gorm.io/gorm"

type LineItem struct {
	gorm.Model
	InvoiceID   uint
	Qty         string
	Description string
	Amount      string
	Total       string
}

type Invoice struct {
	gorm.Model
	ClientName        string
	ClientAddress     string
	IssueDate         string
	TimeToOrder       string
	DepositPercentage string
	Deposit           string
	Balance           string
	Total             string
	Quotation         bool
	Archived          bool
	LineItems         []LineItem `gorm:"foreignKey:InvoiceID" json:"line_items"`
}
