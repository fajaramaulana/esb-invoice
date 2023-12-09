package gorm

import (
	"esb-invoice/server/repositories"
	"esb-invoice/server/repositories/filters"
	"esb-invoice/server/repositories/schemas"

	"gorm.io/gorm"
)

type invoiceRepo struct {
	db *gorm.DB
}

// Create implements repositories.InvoiceRepo.
func (r *invoiceRepo) Create(invoice *schemas.Invoice) (int, error) {
	err := r.db.Create(invoice).Error
	return int(invoice.InvoiceID), err
}

// FindAll implements repositories.InvoiceRepo.
func (r *invoiceRepo) FindAll(filter filters.InvoiceFilter, page int, pageSize int) ([]schemas.Invoice, int64, error) {
	var invoices []schemas.Invoice
	var totalRecords int64

	query := r.db.Model(&schemas.Customer{})

	// Apply filters based on the InvoiceFilter
	if filter.InvoiceID != 0 {
		query = query.Where("invoice_id = ?", filter.InvoiceID)
	}

	if filter.IssueDate != "" {
		query = query.Where("issue_date = ?", filter.IssueDate)
	}

	if filter.DueDate != "" {
		query = query.Where("due_date = ?", filter.DueDate)
	}

	if filter.Subject != "" {
		query = query.Where("subject = ?", filter.Subject)
	}

	if filter.TotalItems != 0 {
		query = query.Where("total_items = ?", filter.TotalItems)
	}

	if filter.CustomerID != 0 {
		query = query.Where("customer_id = ?", filter.CustomerID)
	}

	if filter.UserID != 0 {
		query = query.Where("user_id = ?", filter.UserID)
	}

	if filter.Status != "" {
		query = query.Where("status = ?", filter.Status)
	}

	// Pagination
	offset := (page - 1) * pageSize

	if err := query.Count(&totalRecords).Error; err != nil {
		return nil, 0, err
	}

	if err := query.Offset(offset).Limit(pageSize).Find(&invoices).Error; err != nil {
		return nil, 0, err
	}

	return invoices, totalRecords, nil
}

// FindById implements repositories.InvoiceRepo.
func (r *invoiceRepo) FindById(id int) (*schemas.Invoice, error) {
	var invoice schemas.Invoice
	err := r.db.First(&invoice, id).Error
	return &invoice, err
}

// SoftDelete implements repositories.InvoiceRepo.
func (r *invoiceRepo) SoftDelete(id int) error {
	result := r.db.Delete(&schemas.Invoice{}, id)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

// UpdateById implements repositories.InvoiceRepo.
func (r *invoiceRepo) UpdateById(id int, update *schemas.Invoice) (*schemas.Invoice, error) {
	var existingInvoice schemas.Invoice

	// Find the existing invoice by ID
	result := r.db.First(&existingInvoice, id)
	if result.Error != nil {
		return nil, result.Error
	}

	// update fields
	existingInvoice.Subject = update.Subject
	existingInvoice.IssueDate = update.IssueDate
	existingInvoice.DueDate = update.DueDate
	existingInvoice.CustomerID = update.CustomerID
	existingInvoice.PaymentStatus = update.PaymentStatus
	existingInvoice.TotalItem = update.TotalItem
	existingInvoice.Subtotal = update.Subtotal
	existingInvoice.TaxRate = update.TaxRate
	existingInvoice.TaxAmount = update.TaxAmount
	existingInvoice.TotalAmount = update.TotalAmount
	existingInvoice.UpdatedAt = update.UpdatedAt

	// save the changes
	result = r.db.Save(&existingInvoice)

	if result.Error != nil {
		return nil, result.Error
	}

	return &existingInvoice, nil
}

// UpdateStatusPaid implements repositories.InvoiceRepo.
func (r *invoiceRepo) UpdateStatusPaid(id int, paymentStatus int) (*schemas.Invoice, error) {
	var existingInvoice schemas.Invoice

	result := r.db.First(&existingInvoice, id)

	if result.Error != nil {
		return nil, result.Error
	}

	// update paid status
	existingInvoice.PaymentStatus = paymentStatus

	result = r.db.Save(&existingInvoice)

	if result.Error != nil {
		return nil, result.Error
	}

	return &existingInvoice, nil
}

func NewInvoiceRepository(db *gorm.DB) repositories.InvoiceRepo {
	return &invoiceRepo{
		db: db,
	}
}
