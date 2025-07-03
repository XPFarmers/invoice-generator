package invoices

import (
	"encoding/json"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type LineItem struct {
	gorm.Model
	InvoiceID   uint   `json:"invoiceId"`
	Qty         string `json:"qty"`
	Description string `json:"description"`
	Amount      string `json:"amount"`
	Total       string `json:"total"`
}

type Invoice struct {
	gorm.Model
	ClientName        string     `json:"clientName"`
	ClientAddress     string     `json:"clientAddress"`
	IssueDate         string     `json:"issueDate"`
	TimeToOrder       string     `json:"timeToOrder"`
	DepositPercentage string     `json:"depositPercentage"`
	Deposit           string     `json:"deposit,omitempty"`
	Balance           string     `json:"balance,omitempty"`
	Total             string     `json:"total,omitempty"`
	Quotation         bool       `json:"quotation,omitempty"`
	Archived          bool       `json:"archived,omitempty"`
	LineItems         []LineItem `gorm:"foreignKey:InvoiceID" json:"lineItems"`
}

func GenerateInvoice(body []byte) (*Invoice, error) {
	var invoice Invoice

	// Parse JSON body into Invoice struct
	if err := json.Unmarshal(body, &invoice); err != nil {
		return nil, fmt.Errorf("invalid invoice JSON: %w", err)
	}

	var total float64
	var depositPercentage float64

	// Convert DepositPercentage from string
	fmt.Sscanf(invoice.DepositPercentage, "%f", &depositPercentage)

	// Calculate total from line items
	for _, item := range invoice.LineItems {
		var itemTotal float64
		fmt.Sscanf(item.Total, "%f", &itemTotal)
		total += itemTotal
	}

	deposit := total * (depositPercentage / 100)
	balance := total - deposit

	invoice.DepositPercentage = fmt.Sprintf("%.0f", depositPercentage)
	invoice.Total = fmt.Sprintf("%.2f", total)
	invoice.Deposit = fmt.Sprintf("%.2f", deposit)
	invoice.Balance = fmt.Sprintf("%.2f", balance)
	invoice.IssueDate = time.Now().Format("2006-01-02")
	invoice.Quotation = true
	invoice.Archived = false

	return &invoice, nil
}
