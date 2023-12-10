package gorm

import (
	"esb-invoice/internal/domain/model"
	"esb-invoice/internal/domain/repository"

	"gorm.io/gorm"
)

type typeRepo struct {
	db *gorm.DB
}

func NewTypeRepository(db *gorm.DB) repository.TypeRepo {
	return &typeRepo{
		db: db,
	}
}

// Create implements repositories.TypeRepo.
func (r *typeRepo) Create(item *model.Type) (int, error) {
	err := r.db.Create(item).Error
	return int(item.ID), err
}

// FindById implements repositories.TypeRepo.
func (r *typeRepo) FindById(id int) (*model.Type, error) {
	var item model.Type
	err := r.db.First(&item, id).Error
	return &item, err
}

// SoftDelete implements repositories.TypeRepo.
func (r *typeRepo) SoftDelete(id int) error {
	result := r.db.Delete(&model.Type{}, id)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

// UpdateById implements repositories.TypeRepo.
func (r *typeRepo) UpdateById(id int, update *model.Type) (*model.Type, error) {
	var existingType model.Type

	// Find the existing Item by ID
	result := r.db.First(&existingType, id)
	if result.Error != nil {
		return nil, result.Error
	}

	// update fields
	existingType.Name = update.Name
	existingType.UpdatedAt = update.UpdatedAt

	// save the changes
	result = r.db.Save(&existingType)

	if result.Error != nil {
		return nil, result.Error
	}

	return &existingType, nil
}

func (r *typeRepo) FindAll() ([]model.Type, error) {
	var items []model.Type

	result := r.db.Find(&items)
	if result.Error != nil {
		return nil, result.Error
	}

	return items, nil
}

func (r *typeRepo) FindByName(name string) (*model.Type, error) {
	var item model.Type
	err := r.db.Where("name = ?", name).First(&item).Error
	return &item, err
}

func (r *typeRepo) FindByNameAndNotId(name string, id int) (*model.Type, error) {
	var item model.Type
	err := r.db.Where("name = ? AND id != ?", name, id).First(&item).Error
	return &item, err
}
