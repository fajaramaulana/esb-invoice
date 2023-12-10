package gorm

import (
	"esb-invoice/internal/app/handler/filters"
	"esb-invoice/internal/domain/model"
	"esb-invoice/internal/domain/repository"

	"gorm.io/gorm"
)

type invoiceRepo struct {
	db *gorm.DB
}

func NewInvoiceRepository(db *gorm.DB) repository.InvoiceRepo {
	return &invoiceRepo{
		db: db,
	}
}

// Create implements repositories.InvoiceRepo.
func (r *invoiceRepo) Create(tx *gorm.DB, invoice *model.Invoice) (int, error) {

	if err := tx.Create(invoice).Error; err != nil {
		return 0, err
	}

	return int(invoice.InvoiceID), nil
}

// FindAll implements repositories.InvoiceRepo.
func (r *invoiceRepo) FindAll(filter filters.InvoiceFilter, page int, pageSize int) ([]model.Invoice, int64, error) {
	var invoices []model.Invoice
	var totalRecords int64

	query := r.db.Model(&model.Invoice{})

	if query.Error != nil {
		return nil, 0, query.Error
	}

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
func (r *invoiceRepo) FindById(id int) (*model.Invoice, error) {
	var invoice model.Invoice
	err := r.db.First(&invoice, id).Error
	return &invoice, err
}

// SoftDelete implements repositories.InvoiceRepo.
func (r *invoiceRepo) SoftDelete(id int) error {
	result := r.db.Delete(&model.Invoice{}, id)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

// UpdateById implements repositories.InvoiceRepo.
func (r *invoiceRepo) UpdateById(id int, update *model.Invoice) (*model.Invoice, error) {
	var existingInvoice model.Invoice

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
func (r *invoiceRepo) UpdateStatusPaid(id int, paymentStatus int) (*model.Invoice, error) {
	var existingInvoice model.Invoice

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
