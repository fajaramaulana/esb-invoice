package service

import (
	"esb-invoice/internal/app/handler/filters"
	"esb-invoice/internal/app/handler/request"
	"esb-invoice/internal/app/handler/response"
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

	// count purpose for total item

	var totalItem float64
	var subTotal float64
	// count totalPrice per item

	for _, item := range req.InvoiceItems {
		totalItem += item.Quantity
		subTotal += float64(item.Quantity) * item.Price
	}

	// calculate tax amount
	taxAmount := (req.Tax / 100) * subTotal

	// calculate total amount
	totalAmount := subTotal + taxAmount

	invoice.Subject = req.Subject
	invoice.IssueDate = reqIssueDate
	invoice.DueDate = reqDueDate
	invoice.CustomerID = uint(req.CustomerId)
	invoice.PaymentStatus = 0
	invoice.TotalItem = totalItem
	invoice.Subtotal = subTotal
	invoice.TaxRate = req.Tax
	invoice.TaxAmount = taxAmount
	invoice.TotalAmount = totalAmount
	invoice.CreatedAt = time.Now()
	invoice.UpdatedAt = time.Time{}

	invoiceId, err := s.invoiceRepo.Create(tx, &invoice)

	if err != nil {
		tx.Rollback()
		return 0, err
	}

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
		invoiceItem.TotalPrice = float64(item.Quantity) * item.Price
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

func (s *InvoiceService) GetInvoice(id int) (*response.InvoiceResponse, error) {
	var responseInvoice response.InvoiceResponse
	var ItemInsideInvoice []response.ItemInsideInvoice

	invoice, err := s.invoiceRepo.FindById(id)
	if err != nil {
		return nil, err
	}

	if invoice == nil {
		return nil, fmt.Errorf("invoice id %d not found", id)
	}

	for _, v := range invoice.InvoiceItem {
		ItemInsideInvoice = append(ItemInsideInvoice, response.ItemInsideInvoice{
			ID:         v.ID,
			InvoiceID:  v.InvoiceID,
			ItemID:     v.ItemID,
			Quantity:   v.Quantity,
			UnitPrice:  v.UnitPrice,
			TotalPrice: v.TotalPrice,
			Item: response.DetailItemInsideInvoice{
				ID:     v.Item.ID,
				Name:   v.Item.Name,
				Price:  v.Item.Price,
				TypeID: v.Item.TypeID,
				Type: response.TypeInsideInvoice{
					ID:   uint(v.Item.Type.ID),
					Name: v.Item.Type.Name,
				},
			},
		})
	}

	responseInvoice = response.InvoiceResponse{
		InvoiceId:     invoice.InvoiceID,
		Subject:       invoice.Subject,
		IssueDate:     invoice.IssueDate.Format("02/01/2006"),
		DueDate:       invoice.DueDate.Format("02/01/2006"),
		CustomerId:    invoice.CustomerID,
		PaymentStatus: invoice.PaymentStatus,
		TotalItem:     invoice.TotalItem,
		Subtotal:      invoice.Subtotal,
		TaxRate:       invoice.TaxRate,
		TaxAmount:     invoice.TaxAmount,
		TotalAmount:   invoice.TotalAmount,
		InvoiceItem:   ItemInsideInvoice,
		Customer: response.CustomerInsideInvoice{
			ID:      uint(invoice.Customer.ID),
			Name:    invoice.Customer.Name,
			Email:   invoice.Customer.Email,
			Address: invoice.Customer.Address,
		},
	}

	return &responseInvoice, nil
}

func (s *InvoiceService) FindAll(filter *filters.InvoiceFilter, page int, pageSize int) ([]*response.InvoiceResponse, int64, error) {
	var responseInvoice []response.InvoiceResponse
	var ItemInsideInvoice []response.ItemInsideInvoice

	if filter.IssueDate != "" {
		_, err := time.Parse("02/01/2006", filter.IssueDate)
		if err != nil {
			filter.IssueDate = "01/01/1970"
		}
	}

	if filter.DueDate != "" {
		_, err := time.Parse("02/01/2006", filter.DueDate)
		if err != nil {
			filter.DueDate = "01/01/1970"
		}
	}

	invoices, total, err := s.invoiceRepo.FindAll(*filter, page, pageSize)

	if err != nil {
		return nil, 0, err
	}

	for _, invoice := range invoices {
		for _, v := range invoice.InvoiceItem {
			ItemInsideInvoice = append(ItemInsideInvoice, response.ItemInsideInvoice{
				ID:         v.ID,
				InvoiceID:  v.InvoiceID,
				ItemID:     v.ItemID,
				Quantity:   v.Quantity,
				UnitPrice:  v.UnitPrice,
				TotalPrice: v.TotalPrice,
				Item: response.DetailItemInsideInvoice{
					ID:     v.Item.ID,
					Name:   v.Item.Name,
					Price:  v.Item.Price,
					TypeID: v.Item.TypeID,
					Type: response.TypeInsideInvoice{
						ID:   uint(v.Item.Type.ID),
						Name: v.Item.Type.Name,
					},
				},
			})
		}

		responseInvoice = append(responseInvoice, response.InvoiceResponse{
			InvoiceId:     invoice.InvoiceID,
			Subject:       invoice.Subject,
			IssueDate:     invoice.IssueDate.Format("02/01/2006"),
			DueDate:       invoice.DueDate.Format("02/01/2006"),
			CustomerId:    invoice.CustomerID,
			PaymentStatus: invoice.PaymentStatus,
			TotalItem:     invoice.TotalItem,
			Subtotal:      invoice.Subtotal,
			TaxRate:       invoice.TaxRate,
			TaxAmount:     invoice.TaxAmount,
			TotalAmount:   invoice.TotalAmount,
			InvoiceItem:   ItemInsideInvoice,
			Customer: response.CustomerInsideInvoice{
				ID:      uint(invoice.Customer.ID),
				Name:    invoice.Customer.Name,
				Email:   invoice.Customer.Email,
				Address: invoice.Customer.Address,
			},
		})
	}

	var responseInvoicePtr []*response.InvoiceResponse
	for _, invoice := range responseInvoice {
		responseInvoicePtr = append(responseInvoicePtr, &invoice)
	}
	return responseInvoicePtr, total, nil
}

func (s *InvoiceService) CountAll() (int64, error) {
	totalRecords, err := s.invoiceRepo.CountAll()

	if err != nil {
		return 0, err
	}

	return totalRecords, nil
}

func (s *InvoiceService) Update(id int, req *request.UpdateInvoiceRequest) (*response.InvoiceResponse, error) {
	var invoice model.Invoice
	var invoiceItem model.InvoiceItem
	var invoiceItems []model.InvoiceItem
	var newInvoiceItems []model.InvoiceItem

	// Start a database transaction
	tx := s.db.Begin()

	// Defer the rollback in case of error
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// check invoice id
	invoiceData, err := s.invoiceRepo.FindById(id)
	if err != nil {
		tx.Rollback()
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("invoice id %d not found", id)
		}
		return nil, err
	}

	if invoiceData == nil {
		tx.Rollback()
		return nil, fmt.Errorf("invoice id %d not found", id)
	}

	// check customer id
	customer, err := s.customerRepo.FindById(req.CustomerId)
	if err != nil {
		tx.Rollback()
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("customer not found: %v", req.CustomerId)
		}
		return nil, err
	}

	if customer == nil {
		tx.Rollback()
		return nil, fmt.Errorf("customer not found: %v", req.CustomerId)
	}

	// convert request string date to time

	// convert string date to time
	// from 06/01/2021 to time.Time
	reqIssueDate, err := time.Parse("02/01/2006", req.IssueDate)
	if err != nil {
		return nil, fmt.Errorf("error parse issue date: %s", err.Error())
	}

	reqDueDate, err := time.Parse("02/01/2006", req.DueDate)
	if err != nil {
		return nil, fmt.Errorf("error parse due date: %s", err.Error())
	}

	// count purpose for total item

	var totalItem float64
	var subTotal float64
	// count totalPrice per item

	for _, item := range req.InvoiceItems {
		if item.IsDeleted == false 
			totalItem += item.Quantity
			subTotal += float64(item.Quantity) * item.Price
		}
	}

	// calculate tax amount
	taxAmount := (req.Tax / 100) * subTotal

	// calculate total amount
	totalAmount := subTotal + taxAmount

	invoice.Subject = req.Subject
	invoice.IssueDate = reqIssueDate
	invoice.DueDate = reqDueDate
	invoice.CustomerID = uint(req.CustomerId)
	invoice.PaymentStatus = 0
	invoice.TotalItem = totalItem
	invoice.Subtotal = subTotal
	invoice.TaxRate = req.Tax
	invoice.TaxAmount = taxAmount
	invoice.TotalAmount = totalAmount
	invoice.UpdatedAt = time.Now()

	_, err = s.invoiceRepo.UpdateById(tx, id, &invoice)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	for _, item := range req.InvoiceItems {
		// check item id
		itemData, err := s.itemRepo.FindById(item.ProductId)

		if err != nil {
			tx.Rollback()
			if err == gorm.ErrRecordNotFound {
				return nil, fmt.Errorf("item id %d not found", item.ProductId)
			}

			return nil, err
		}

		if itemData == nil {
			tx.Rollback()
			return nil, fmt.Errorf("item id %d not found", item.ProductId)
		}

		if item.Id != 0 { // update invoice item
			if item.IsDeleted == true { // soft delete invoice item
				err = s.invoiceItemRepo.SoftDelete(tx, item.Id)

				if err != nil {
					if err != gorm.ErrRecordNotFound {
						tx.Rollback()
						return nil, err
					}
				}
			} else {
				invoiceItem.InvoiceID = uint(item.Id)
				invoiceItem.ItemID = uint(item.ProductId)
				invoiceItem.Quantity = item.Quantity
				invoiceItem.UnitPrice = item.Price
				invoiceItem.TotalPrice = float64(item.Quantity) * item.Price
				invoiceItem.CreatedAt = time.Now()
				invoiceItem.UpdatedAt = time.Time{}

				_, err = s.invoiceItemRepo.UpdateById(tx, item.Id, &invoiceItem)
				if err != nil {
					tx.Rollback()
					return nil, err
				}
			}
		} else { // create invoice item
			invoiceItem.InvoiceID = uint(id)
			invoiceItem.ItemID = uint(item.ProductId)
			invoiceItem.Quantity = item.Quantity
			invoiceItem.UnitPrice = item.Price
			invoiceItem.TotalPrice = float64(item.Quantity) * item.Price
			invoiceItem.CreatedAt = time.Now()
			invoiceItem.UpdatedAt = time.Time{}

			// apped invoice item to invoice items
			newInvoiceItems = append(newInvoiceItems, invoiceItem)
		}

		invoiceItems = append(invoiceItems, invoiceItem)
	}

	_, err = s.invoiceItemRepo.Create(tx, &newInvoiceItems)

	if err != nil {
		tx.Rollback()
		return nil, err
	}

	// Commit the transaction
	err = tx.Commit().Error
	if err != nil {
		return nil, err
	}

	var responseInvoice response.InvoiceResponse
	var ItemInsideInvoice []response.ItemInsideInvoice

	invoiceUpdate, err := s.invoiceRepo.FindById(id)
	if err != nil {
		return nil, err
	}

	if invoiceUpdate == nil {
		return nil, fmt.Errorf("invoice id %d not found", id)
	}

	for _, v := range invoice.InvoiceItem {
		ItemInsideInvoice = append(ItemInsideInvoice, response.ItemInsideInvoice{
			ID:         v.ID,
			InvoiceID:  v.InvoiceID,
			ItemID:     v.ItemID,
			Quantity:   v.Quantity,
			UnitPrice:  v.UnitPrice,
			TotalPrice: v.TotalPrice,
			Item: response.DetailItemInsideInvoice{
				ID:     v.Item.ID,
				Name:   v.Item.Name,
				Price:  v.Item.Price,
				TypeID: v.Item.TypeID,
				Type: response.TypeInsideInvoice{
					ID:   uint(v.Item.Type.ID),
					Name: v.Item.Type.Name,
				},
			},
		})
	}

	responseInvoice = response.InvoiceResponse{
		InvoiceId:     invoiceUpdate.InvoiceID,
		Subject:       invoiceUpdate.Subject,
		IssueDate:     invoiceUpdate.IssueDate.Format("02/01/2006"),
		DueDate:       invoiceUpdate.DueDate.Format("02/01/2006"),
		CustomerId:    invoiceUpdate.CustomerID,
		PaymentStatus: invoiceUpdate.PaymentStatus,
		TotalItem:     invoiceUpdate.TotalItem,
		Subtotal:      invoiceUpdate.Subtotal,
		TaxRate:       invoiceUpdate.TaxRate,
		TaxAmount:     invoiceUpdate.TaxAmount,
		TotalAmount:   invoiceUpdate.TotalAmount,
		InvoiceItem:   ItemInsideInvoice,
		Customer: response.CustomerInsideInvoice{
			ID:      uint(invoiceUpdate.Customer.ID),
			Name:    invoiceUpdate.Customer.Name,
			Email:   invoiceUpdate.Customer.Email,
			Address: invoiceUpdate.Customer.Address,
		},
	}

	return &responseInvoice, nil

}
