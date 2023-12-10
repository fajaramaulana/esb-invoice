package gorm

import (
	"esb-invoice/internal/app/handler/filters"
	"esb-invoice/internal/domain/model"
	"esb-invoice/internal/domain/repository"

	"gorm.io/gorm"
)

type customerRepo struct {
	db *gorm.DB
}

func NewCustomerRepository(db *gorm.DB) repository.CustomerRepo {
	return &customerRepo{
		db: db,
	}
}

// Create implements repositories.CustomerRepo.
func (r *customerRepo) Create(comment *model.Customer) (int, error) {
	err := r.db.Create(comment).Error
	return comment.ID, err
}

// FindAll implements repositories.CustomerRepo.
func (r *customerRepo) FindAll(filter filters.CustomerFilter, page int, pageSize int) ([]model.Customer, int64, error) {
	var customers []model.Customer
	var totalRecords int64

	query := r.db.Model(&model.Customer{})

	if query.Error != nil {
		return nil, 0, query.Error
	}

	// Apply filters based on the CustomerFilter
	if filter.Name != "" {
		query = query.Where("name LIKE ?", "%"+filter.Name+"%")
	}

	if filter.Email != "" {
		query = query.Where("email LIKE ?", "%"+filter.Email+"%")
	}

	if filter.Address != "" {
		// where like
		query = query.Where("address LIKE ?", "%"+filter.Address+"%")
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
func (r *customerRepo) FindById(id int) (*model.Customer, error) {
	var customer model.Customer
	err := r.db.First(&customer, id).Error
	return &customer, err
}

// SoftDelete implements repositories.CustomerRepo.
func (r *customerRepo) SoftDelete(id int) error {
	result := r.db.Delete(&model.Customer{}, id)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

// UpdateById implements repositories.CustomerRepo.
func (r *customerRepo) UpdateById(id int, update *model.Customer) (*model.Customer, error) {
	var existingCustomer model.Customer

	// Find the existing invoice by ID
	result := r.db.First(&existingCustomer, id)
	if result.Error != nil {
		return nil, result.Error
	}

	// update fields
	existingCustomer.Name = update.Name
	existingCustomer.Email = update.Email
	existingCustomer.UpdatedAt = update.UpdatedAt

	// save the changes
	result = r.db.Save(&existingCustomer)

	if result.Error != nil {
		return nil, result.Error
	}

	return &existingCustomer, nil
}

func (r *customerRepo) FindByEmailAndNotId(email string, id int) (*model.Customer, error) {
	var customer model.Customer
	err := r.db.Where("email = ?", email).Not("id = ?", id).First(&customer).Error
	return &customer, err
}

func (r *customerRepo) FindByEmail(email string) (*model.Customer, error) {
	var customer model.Customer
	err := r.db.Where("email = ?", email).First(&customer).Error
	return &customer, err
}

func (r *customerRepo) CountAll() (int64, error) {
	var totalRecords int64

	query := r.db.Model(&model.Customer{})

	if query.Error != nil {
		return 0, query.Error
	}

	if err := query.Count(&totalRecords).Error; err != nil {
		return 0, err
	}

	return totalRecords, nil
}
