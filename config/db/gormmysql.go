package db

import (
	"errors"
	"esb-invoice/server/repositories/schemas"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Schema interface{}

func ConnectionMysqlGorm() (*gorm.DB, error) {
	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Error load env %s", err)
		return nil, err
	}

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
		"customers":     &schemas.Customer{},
		"items":         &schemas.Item{},
		"invoices":      &schemas.Invoice{},
		"invoice_items": &schemas.InvoiceItem{},
		"types":         &schemas.Type{},
	}

	var errMessage string

	for tableName, schema := range tableSchemas {
		// check if table already exist
		if exist := db.Migrator().HasTable(tableName); !exist {
			// if table doesnt exist, migrate
			if err := db.AutoMigrate(schema); err != nil {
				log.Fatalf("Error migrating table %s: %s", tableName, err)
				errMessage = fmt.Sprintf("Error migrating table %s: %s", tableName, err)
			}
			log.Printf("Table %s migrated successfully\n", tableName)
			errMessage = "nil"
		} else {
			log.Printf("Table %s already exists\n", tableName)
			errMessage = "nil"
		}
	}

	return errMessage
}
