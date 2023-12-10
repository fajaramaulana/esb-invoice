package db

import (
	"errors"
	"esb-invoice/internal/domain/model"
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Schema interface{}

func ConnectionMysqlGorm() (*gorm.DB, error) {

	dsn := os.Getenv("DB_USER") + ":" + os.Getenv("DB_PASSWORD") + "@tcp(" + os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT") + ")/" + os.Getenv("DB_DATABASE") + "?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("Error connection database %s", err)
		return nil, err
	}

	if err := migrateProcess(db); err != "nil" {
		log.Fatalf(err)
		return nil, errors.New(err)
	}

	return db, nil
}

func migrateProcess(db *gorm.DB) string {
	tableSchemas := map[string]Schema{
		"customers":     &model.Customer{},
		"items":         &model.Item{},
		"invoices":      &model.Invoice{},
		"invoice_items": &model.InvoiceItem{},
		"types":         &model.Type{},
	}

	var errMessage string

	// loop table schemas and migrate if success create .sql file
	for tableName, schema := range tableSchemas {
		// check if table already exist
		if exist := db.Migrator().HasTable(tableName); !exist {
			// If the table doesn't exist, migrate and save SQL file
			if err := db.AutoMigrate(schema); err != nil {
				log.Fatalf("Error migrating table %s: %s", tableName, err)
				errMessage = fmt.Sprintf("Error migrating table %s: %s", tableName, err)
			}
			log.Printf("Table %s created\n", tableName)
			errMessage = "nil"
		} else {
			errMessage = "nil"
		}
	}
	return errMessage
}

func ConnectionMysqlGormTest() (*gorm.DB, error) {

	dsn := os.Getenv("DB_USER") + ":" + os.Getenv("DB_PASSWORD") + "@tcp(" + os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT") + ")/" + os.Getenv("DB_DATABASE_TEST") + "?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("Error connection database %s", err)
		return nil, err
	}

	if err := migrateProcess(db); err != "nil" {
		log.Fatalf(err)
		return nil, errors.New(err)
	}

	return db, nil
}
