package gorm

import (
	"esb-invoice/server/repositories"
	"esb-invoice/server/repositories/filters"
	"esb-invoice/server/repositories/schemas"

	"gorm.io/gorm"
)

type itemRepo struct {
	db *gorm.DB
}

// Create implements repositories.ItemRepo.
func (r *itemRepo) Create(item *schemas.Item) (int, error) {
	err := r.db.Create(item).Error
	return int(item.ID), err
}

// FindById implements repositories.ItemRepo.
func (r *itemRepo) FindById(id int) (*schemas.Item, error) {
	var item schemas.Item
	err := r.db.First(&item, id).Error
	return &item, err
}

// SoftDelete implements repositories.ItemRepo.
func (r *itemRepo) SoftDelete(id int) error {
	result := r.db.Delete(&schemas.Item{}, id)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

// UpdateById implements repositories.ItemRepo.
func (r *itemRepo) UpdateById(id int, update *schemas.Item) (*schemas.Item, error) {
	var existingItem schemas.Item

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

func (r *itemRepo) FindAll(filter filters.ItemFilter, page int, pageSize int) ([]schemas.Item, int64, error) {
	var items []schemas.Item
	var totalRecords int64

	query := r.db.Model(&schemas.Item{})

	if query.Error != nil {
		return nil, 0, query.Error
	}

	// Apply filters based on the ItemFilter
	if filter.ItemName != "" {
		query = query.Where("name = ?", filter.ItemName)
	}

	if filter.TypeName != "" {
		// query join to types table
		query = query.Joins("JOIN types ON types.id = items.type_id")
		query = query.Where("types.name = ?", filter.TypeName)
	}

	// end filter

	// Pagination
	offset := (page - 1) * pageSize

	if err := query.Count(&totalRecords).Error; err != nil {
		return nil, 0, err
	}

	if err := query.Offset(offset).Limit(pageSize).Find(&items).Error; err != nil {
		return nil, 0, err
	}

	return items, totalRecords, nil

}

func NewItemRepository(db *gorm.DB) repositories.ItemRepo {
	return &itemRepo{
		db: db,
	}
}
