package gorm

import (
	"esb-invoice/server/repositories"
	"esb-invoice/server/repositories/filters"
	"esb-invoice/server/repositories/schemas"

	"gorm.io/gorm"
)

type customerRepo struct {
	db *gorm.DB
}

func NewCustomerRepository(db *gorm.DB) repositories.CustomerRepo {
	return &customerRepo{
		db: db,
	}
}

// Create implements repositories.CustomerRepo.
func (r *customerRepo) Create(comment *schemas.Customer) (int, error) {
	err := r.db.Create(comment).Error
	return comment.ID, err
}

// FindAll implements repositories.CustomerRepo.
func (r *customerRepo) FindAll(filter filters.CustomerFilter, page int, pageSize int) ([]schemas.Customer, int64, error) {
	var customers []schemas.Customer
	var totalRecords int64

	query := r.db.Model(&schemas.Customer{})

	if query.Error != nil {
		return nil, 0, query.Error
	}

	// Apply filters based on the CustomerFilter
	if filter.Name != "" {
		query = query.Where("invoice_id = ?", filter.Name)
	}

	// Pagination
	offset := (page - 1) * pageSize

	if err := query.Count(&totalRecords).Error; err != nil {
		return nil, 0, err
	}

	if err := query.Offset(offset).Limit(pageSize).Find(&customers).Error; err != nil {
		return nil, 0, err
	}

	return customers, totalRecords, nil
}

// FindById implements repositories.CustomerRepo.
func (r *customerRepo) FindById(id int) (*schemas.Customer, error) {
	var customer schemas.Customer
	err := r.db.First(&customer, id).Error
	return &customer, err
}

// SoftDelete implements repositories.CustomerRepo.
func (r *customerRepo) SoftDelete(id int) error {
	result := r.db.Delete(&schemas.Customer{}, id)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

// UpdateById implements repositories.CustomerRepo.
func (r *customerRepo) UpdateById(id int, update *schemas.Customer) (*schemas.Customer, error) {
	var existingCustomer schemas.Customer

	// Find the existing invoice by ID
	result := r.db.First(&existingCustomer, id)
	if result.Error != nil {
		return nil, result.Error
	}

	// update fields
	existingCustomer.Name = update.Name
	existingCustomer.UpdatedAt = update.UpdatedAt

	// save the changes
	result = r.db.Save(&existingCustomer)

	if result.Error != nil {
		return nil, result.Error
	}

	return &existingCustomer, nil
}
