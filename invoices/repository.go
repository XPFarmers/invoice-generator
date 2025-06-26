package invoices

import (
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Create(invoice *Invoice) error {
	return r.db.Create(invoice).Error
}

func (r *Repository) GetAll() ([]Invoice, error) {
	var invoices []Invoice
	err := r.db.Preload("LineItems").
		Where("archived = ?", false).
		Order("id desc").
		Find(&invoices).Error
	return invoices, err
}

func (r *Repository) GetByID(id uint) (*Invoice, error) {
	var invoice Invoice
	err := r.db.Preload("LineItems").
		Where("id = ? AND archived = ?", id, false).
		First(&invoice).Error
	return &invoice, err
}

func (r *Repository) SetQuotationFalse(id uint) error {
	return r.db.Model(&Invoice{}).
		Where("id = ?", id).
		Update("quotation", false).Error
}

func (r *Repository) Archive(id uint) error {
	return r.db.Model(&Invoice{}).
		Where("id = ?", id).
		Update("archived", true).Error
}

func (r *Repository) GetArchived() ([]Invoice, error) {
	var invoices []Invoice
	err := r.db.Preload("LineItems").
		Where("archived = ?", true).
		Order("id desc").
		Find(&invoices).Error
	return invoices, err
}

func (r *Repository) Update(invoice *Invoice) error {
	// Make sure it's not archived
	var existing Invoice
	err := r.db.Where("id = ? AND archived = ?", invoice.ID, false).First(&existing).Error
	if err != nil {
		return err
	}

	// Update base invoice
	err = r.db.Save(invoice).Error
	if err != nil {
		return err
	}

	// Replace line items
	err = r.db.Where("invoice_id = ?", invoice.ID).Delete(&LineItem{}).Error
	if err != nil {
		return err
	}

	for _, item := range invoice.LineItems {
		item.InvoiceID = invoice.ID
		if err := r.db.Create(&item).Error; err != nil {
			return err
		}
	}

	return nil
}
