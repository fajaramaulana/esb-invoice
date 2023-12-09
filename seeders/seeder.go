package seeders

import (
	"esb-invoice/internal/domain/model"
	"log"
	"time"

	"gorm.io/gorm"
)

// run seeder
func SeederDB(dbInstance *gorm.DB) error {
	// run seeder for table customers
	if isEmpty(dbInstance, "customers") {
		if err := seedCustomers(dbInstance); err != nil {
			log.Fatal("Error Seeding Table Customer: ", err)
			return err
		}
	}

	// run seeder for table types
	if isEmpty(dbInstance, "types") {
		if err := seedType(dbInstance); err != nil {
			log.Fatal("Error Seeding Table Type: ", err)
			return err
		}
	}

	// run seeder for table items
	if isEmpty(dbInstance, "items") {
		if err := seedItem(dbInstance); err != nil {
			log.Fatal("Error Seeding Table Items: ", err)
			return err
		}
	}

	return nil
}

func seedCustomers(dbdbInstance *gorm.DB) error {
	customers := []model.Customer{
		{
			Name:      "Barrington Publishers",
			Email:     "barringtonpublisher@gmail.com",
			Address:   "17 Great Suffolk Street London SE1 0NS United Kingdom",
			CreatedAt: time.Now(),
			UpdatedAt: time.Time{},
		},
		{
			Name:      "Fajar Agus Maulana",
			Email:     "fajaragusmaulana@gmail.com",
			Address:   "Kelapa Dua, Tangerang",
			CreatedAt: time.Now(),
			UpdatedAt: time.Time{},
		},
	}

	if err := dbdbInstance.Create(&customers).Error; err != nil {
		log.Fatal("Error Insert Seeding Customer: ", err)
		return err
	}

	return nil
}

func seedType(dbdbInstance *gorm.DB) error {
	typeItem := []model.Type{
		{
			Name:      "Service",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			Name:      "Hardware",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	if err := dbdbInstance.Create(&typeItem).Error; err != nil {
		log.Fatal("Error Insert Seeding Type: ", err)
		return err
	}

	return nil
}

func seedItem(dbdbInstance *gorm.DB) error {
	item := []model.Item{
		{
			Name:   "Design",
			Price:  41.00,
			TypeID: 1,
		},
		{
			Name:   "Development",
			Price:  57.00,
			TypeID: 1,
		},
		{
			Name:   "Meetings",
			Price:  4.50,
			TypeID: 1,
		},
		{
			Name:   "Printer",
			Price:  22.00,
			TypeID: 2,
		},
		{
			Name:   "Monitor",
			Price:  29.70,
			TypeID: 2,
		},
	}

	if err := dbdbInstance.Create(&item).Error; err != nil {
		log.Fatal("Error Insert Seeding Item: ", err)
		return err
	}

	return nil
}

// checks if a table is empty
func isEmpty(db *gorm.DB, tableName string) bool {
	var count int64
	db.Table(tableName).Count(&count)
	return count == 0
}
