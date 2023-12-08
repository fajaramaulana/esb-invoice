package db

import (
	"esb-invoice/server/repositories/schemas"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectionMysqlGorm() (*gorm.DB, error) {
	err := godotenv.Load()

	if err != nil {
		return nil, err
	}

	dsn := os.Getenv("DB_USER") + ":" + os.Getenv("DB_PASSWORD") + "@tcp(" + os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT") + ")/" + os.Getenv("DB_DATABASE") + "?charset=utf8mb4&parseTime=True&loc=Local"
	fmt.Printf("%# v\n", dsn)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	db.Debug().AutoMigrate(schemas.Customer{}, schemas.Item{}, schemas.Invoice{}, schemas.InvoiceItem{}, schemas.Type{})

	return db, nil
}
