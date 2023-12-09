package gorm

import (
	"esb-invoice/internal/domain/model"
	"esb-invoice/internal/domain/repository"

	"gorm.io/gorm"
)

type invoiceitemRepo struct {
	db *gorm.DB
}

func NewInvoiceItemRepository(db *gorm.DB) repository.InvoiceItemRepo {
	return &invoiceitemRepo{
		db: db,
	}
}

// Create implements repositories.InvoiceItemRepo.
func (r *invoiceitemRepo) Create(invoiceItem *[]model.InvoiceItem) ([]int, error) {
	// input multiple invoice item and return multiple invoice item id
	var invoiceItemId []int
	for _, item := range *invoiceItem {
		err := r.db.Create(item).Error
		if err != nil {
			return nil, err
		}
		invoiceItemId = append(invoiceItemId, int(item.ID))
	}
	return invoiceItemId, nil

}

// FindById implements repositories.InvoiceItemRepo.
func (r *invoiceitemRepo) FindById(id int) (*model.InvoiceItem, error) {
	// find invoice item by id
	var invoiceItem model.InvoiceItem
	err := r.db.First(&invoiceItem, id).Error
	return &invoiceItem, err
}

// FindByInvoiceId implements repositories.InvoiceItemRepo.
func (r *invoiceitemRepo) FindByInvoiceId(idInvoice int) ([]model.InvoiceItem, error) {
	// find invoice item by invoice id
	var invoiceItem []model.InvoiceItem
	err := r.db.Where("invoice_id = ?", idInvoice).Find(&invoiceItem).Error
	return invoiceItem, err
}

// SoftDelete implements repositories.InvoiceItemRepo.
func (r *invoiceitemRepo) SoftDelete(id int) error {
	result := r.db.Delete(&model.InvoiceItem{}, id)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

// UpdateById implements repositories.InvoiceItemRepo.
func (r *invoiceitemRepo) UpdateById(id int, update *model.InvoiceItem) (*model.InvoiceItem, error) {
	// find invoice item by id
	var existingInvoiceItem model.InvoiceItem
	result := r.db.First(&existingInvoiceItem, id)
	if result.Error != nil {
		return nil, result.Error
	}

	// update fields
	existingInvoiceItem.InvoiceID = update.InvoiceID
	existingInvoiceItem.ItemID = update.ItemID
	existingInvoiceItem.Quantity = update.Quantity
	existingInvoiceItem.UnitPrice = update.UnitPrice
	existingInvoiceItem.TotalPrice = update.TotalPrice
	existingInvoiceItem.UpdatedAt = update.UpdatedAt

	// save the changes
	result = r.db.Save(&existingInvoiceItem)

	if result.Error != nil {
		return nil, result.Error
	}

	return &existingInvoiceItem, nil
}
