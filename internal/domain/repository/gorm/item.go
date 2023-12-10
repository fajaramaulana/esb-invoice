package gorm

import (
	"esb-invoice/internal/app/handler/filters"
	"esb-invoice/internal/domain/model"
	"esb-invoice/internal/domain/repository"

	"gorm.io/gorm"
)

type itemRepo struct {
	db *gorm.DB
}

func NewItemRepository(db *gorm.DB) repository.ItemRepo {
	return &itemRepo{
		db: db,
	}
}

// Create implements repositories.ItemRepo.
func (r *itemRepo) Create(item *model.Item) (int, error) {
	err := r.db.Create(item).Error
	return int(item.ID), err
}

// FindById implements repositories.ItemRepo.
func (r *itemRepo) FindById(id int) (*model.Item, error) {
	var item model.Item
	err := r.db.Preload("Type").First(&item, id).Error
	return &item, err
}

// SoftDelete implements repositories.ItemRepo.
func (r *itemRepo) SoftDelete(id int) error {
	result := r.db.Delete(&model.Item{}, id)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

// UpdateById implements repositories.ItemRepo.
func (r *itemRepo) UpdateById(id int, update *model.Item) (*model.Item, error) {
	var existingItem model.Item

	// Find the existing Item by ID
	result := r.db.First(&existingItem, id)
	if result.Error != nil {
		return nil, result.Error
	}

	// update fields
	existingItem.Name = update.Name
	existingItem.Price = update.Price
	existingItem.TypeID = update.TypeID
	existingItem.UpdatedAt = update.UpdatedAt

	// save the changes
	result = r.db.Save(&existingItem)

	if result.Error != nil {
		return nil, result.Error
	}

	return &existingItem, nil
}

func (r *itemRepo) FindAll(filter filters.ItemFilter, page int, pageSize int) ([]model.Item, int64, error) {
	var items []model.Item
	var totalRecords int64

	query := r.db.Model(&model.Item{}).Preload("Type").Order("items.id DESC")
	if query.Error != nil {
		return nil, 0, query.Error
	}

	// Apply filters based on the Name for item name or type name
	if filter.Name != "" {
		// join table types
		query = query.Where("items.name LIKE ?", "%"+filter.Name+"%").Or("types.name LIKE ?", "%"+filter.Name+"%").Joins("JOIN types ON types.id = items.type_id")
	}
	// end filter

	// Pagination
	offset := (page - 1) * pageSize

	if err := query.Offset(offset).Limit(pageSize).Find(&items).Error; err != nil {
		return nil, 0, err
	}

	if err := query.Count(&totalRecords).Error; err != nil {
		return nil, 0, err
	}

	return items, totalRecords, nil
}

func (r *itemRepo) CountAll() (int64, error) {
	var count int64
	if err := r.db.Model(&model.Item{}).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

func (r *itemRepo) FindByItemName(name string) (*model.Item, error) {
	var item model.Item
	err := r.db.Where("name = ?", name).First(&item).Error
	return &item, err
}

func (r *itemRepo) FindByNameAndNotId(name string, id int) (*model.Item, error) {
	var item model.Item
	err := r.db.Where("name = ? AND id != ?", name, id).First(&item).Error
	return &item, err
}
