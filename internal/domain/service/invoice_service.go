package service

import (
	"esb-invoice/internal/app/handler/request"
	"esb-invoice/internal/domain/model"
	"esb-invoice/internal/domain/repository"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type InvoiceService struct {
	db              *gorm.DB
	invoiceRepo     repository.InvoiceRepo
	invoiceItemRepo repository.InvoiceItemRepo
	customerRepo    repository.CustomerRepo
	itemRepo        repository.ItemRepo
}

func NewInvoiceService(db *gorm.DB, invoiceRepo repository.InvoiceRepo, invoiceItemRepo repository.InvoiceItemRepo, customerRepo repository.CustomerRepo, itemRepo repository.ItemRepo) *InvoiceService {
	return &InvoiceService{
		db:              db,
		invoiceRepo:     invoiceRepo,
		invoiceItemRepo: invoiceItemRepo,
		customerRepo:    customerRepo,
		itemRepo:        itemRepo,
	}
}

func (s *InvoiceService) Create(req *request.CreateInvoiceRequest) (int, error) {
	var invoice model.Invoice
	var invoiceItem model.InvoiceItem
	var invoiceItems []model.InvoiceItem

	// Start a database transaction
	tx := s.db.Begin()

	// Defer the rollback in case of error
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// check customer id
	customer, err := s.customerRepo.FindById(req.CustomerId)
	if err != nil {
		tx.Rollback()
		if err == gorm.ErrRecordNotFound {
			return 0, fmt.Errorf("customer id %d not found", req.CustomerId)
		}
		return 0, err
	}

	if customer == nil {
		tx.Rollback()
		return 0, fmt.Errorf("customer id %d not found", req.CustomerId)
	}

	// convert request string date to time

	// convert string date to time
	// from 06/01/2021 to time.Time
	reqIssueDate, err := time.Parse("02/01/2006", req.IssueDate)
	if err != nil {
		return 0, fmt.Errorf("error parse issue date: %s", err.Error())
	}

	reqDueDate, err := time.Parse("02/01/2006", req.DueDate)
	if err != nil {
		return 0, fmt.Errorf("error parse due date: %s", err.Error())
	}

	invoice.Subject = req.Subject
	invoice.IssueDate = reqIssueDate
	invoice.DueDate = reqDueDate
	invoice.CustomerID = uint(req.CustomerId)
	invoice.PaymentStatus = 0
	invoice.TotalItem = req.TotalItem
	invoice.Subtotal = req.SubTotal
	invoice.TaxRate = req.Tax
	invoice.TaxAmount = req.TaxAmount
	invoice.TotalAmount = req.TotalAmount
	invoice.CreatedAt = time.Now()
	invoice.UpdatedAt = time.Time{}

	invoiceId, err := s.invoiceRepo.Create(tx, &invoice)

	if err != nil {
		tx.Rollback()
		return 0, err
	}

	fmt.Printf("%# v\n", req.InvoiceItems)

	for _, item := range req.InvoiceItems {
		// check item id
		itemData, err := s.itemRepo.FindById(item.ProductId)
		if err != nil {
			tx.Rollback()
			if err == gorm.ErrRecordNotFound {
				return 0, fmt.Errorf("item id %d not found", item.ProductId)
			}

			return 0, err
		}

		if itemData == nil {
			tx.Rollback()
			return 0, fmt.Errorf("item id %d not found", item.ProductId)
		}

		invoiceItem.InvoiceID = uint(invoiceId)
		invoiceItem.ItemID = uint(item.ProductId)
		invoiceItem.Quantity = item.Quantity
		invoiceItem.UnitPrice = item.Price
		invoiceItem.TotalPrice = item.TotalPrice
		invoiceItem.CreatedAt = time.Now()
		invoiceItem.UpdatedAt = time.Time{}

		// apped invoice item to invoice items
		invoiceItems = append(invoiceItems, invoiceItem)
	}

	_, err = s.invoiceItemRepo.Create(tx, &invoiceItems)
	if err != nil {
		fmt.Printf("%# v\n", "error disini")
		tx.Rollback()
		return 0, err
	}

	// Commit the transaction
	err = tx.Commit().Error
	if err != nil {
		return 0, err
	}

	return invoiceId, nil
}
