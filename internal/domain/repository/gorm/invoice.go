package gorm

import (
	"esb-invoice/internal/app/handler/filters"
	"esb-invoice/internal/domain/model"
	"esb-invoice/internal/domain/repository"
	"fmt"
	"time"

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

	query := r.db.Model(&model.Invoice{}).Preload("Customer").Preload("InvoiceItem").Preload("InvoiceItem.Item").Preload("InvoiceItem.Item.Type").Order("invoice_id DESC")

	if query.Error != nil {
		return nil, 0, query.Error
	}

	// Apply filters based on the InvoiceFilter
	if filter.InvoiceID != "" {
		query = query.Where("invoice_id = ?", filter.InvoiceID)
	}

	if filter.IssueDate != "" {
		issueDate, err := time.Parse("02/01/2006", filter.IssueDate)
		if err != nil {
			return nil, 0, fmt.Errorf("error parse issue date: %s", err.Error())
		}
		query = query.Where("issue_date = ?", issueDate)
	}

	if filter.DueDate != "" {
		dueDate, err := time.Parse("02/01/2006", filter.DueDate)
		if err != nil {
			return nil, 0, fmt.Errorf("error parse due date: %s", err.Error())
		}
		query = query.Where("due_date = ?", dueDate)
	}

	if filter.Subject != "" {
		// where like
		query = query.Where("subject LIKE ?", "%"+filter.Subject+"%")
	}

	if filter.TotalItem != "" {
		query = query.Where("total_item = ?", filter.TotalItem)
	}

	if filter.CustomerName != "" {
		query = query.Joins("JOIN customers ON invoices.customer_id = customers.id").Where("customers.name LIKE ?", "%"+filter.CustomerName+"%")
	}

	if filter.PaymentStatus != 0 {
		query = query.Where("payment_status = ?", filter.PaymentStatus)
	}

	// Pagination
	offset := (page - 1) * pageSize

	if err := query.Offset(offset).Limit(pageSize).Find(&invoices).Error; err != nil {
		return nil, 0, err
	}

	if err := query.Count(&totalRecords).Error; err != nil {
		return nil, 0, err
	}

	return invoices, totalRecords, nil
}

// FindById implements repositories.InvoiceRepo.
func (r *invoiceRepo) FindById(id int) (*model.Invoice, error) {
	var invoice model.Invoice
	err := r.db.Preload("Customer").Preload("InvoiceItem").Preload("InvoiceItem.Item").Preload("InvoiceItem.Item.Type").First(&invoice, id).Error
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
func (r *invoiceRepo) UpdateById(tx *gorm.DB, id int, update *model.Invoice) (*model.Invoice, error) {
	var existingInvoice model.Invoice

	result := tx.First(&existingInvoice, id)

	if result.Error != nil {
		return nil, result.Error
	}

	// update invoice
	existingInvoice.Subject = update.Subject
	existingInvoice.IssueDate = update.IssueDate
	existingInvoice.DueDate = update.DueDate
	existingInvoice.CustomerID = update.CustomerID
	existingInvoice.TotalItem = update.TotalItem
	existingInvoice.Subtotal = update.Subtotal
	existingInvoice.TaxRate = update.TaxRate
	existingInvoice.TaxAmount = update.TaxAmount
	existingInvoice.TotalAmount = update.TotalAmount
	existingInvoice.PaymentStatus = update.PaymentStatus
	existingInvoice.UpdatedAt = time.Now()

	result = tx.Save(&existingInvoice)

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

func (r *invoiceRepo) CountAll() (int64, error) {
	var totalInvoice int64
	result := r.db.Model(&model.Invoice{}).Count(&totalInvoice)
	if result.Error != nil {
		return 0, result.Error
	}
	return totalInvoice, nil
}
