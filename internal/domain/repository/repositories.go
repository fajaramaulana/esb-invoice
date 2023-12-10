package repository

import (
	"esb-invoice/internal/app/handler/filters"
	"esb-invoice/internal/domain/model"

	"gorm.io/gorm"
)

type TypeRepo interface {
	Create(typeItem *model.Type) (int, error)
	FindAll() ([]model.Type, error)
	FindById(id int) (*model.Type, error)
	UpdateById(id int, update *model.Type) (*model.Type, error)
	SoftDelete(id int) error
	FindByName(name string) (*model.Type, error)
	FindByNameAndNotId(name string, id int) (*model.Type, error)
}

type ItemRepo interface {
	Create(item *model.Item) (int, error)
	FindById(id int) (*model.Item, error)
	FindAll(filter filters.ItemFilter, page int, pageSize int) ([]model.Item, int64, error)
	UpdateById(id int, update *model.Item) (*model.Item, error)
	SoftDelete(id int) error
	CountAll() (int64, error)
	FindByItemName(name string) (*model.Item, error)
	FindByNameAndNotId(name string, id int) (*model.Item, error)
}

type CustomerRepo interface {
	Create(item *model.Customer) (int, error)
	FindById(id int) (*model.Customer, error)
	FindAll(filter filters.CustomerFilter, page int, pageSize int) ([]model.Customer, int64, error)
	UpdateById(id int, update *model.Customer) (*model.Customer, error)
	SoftDelete(id int) error
	FindByEmail(email string) (*model.Customer, error)
	FindByEmailAndNotId(email string, id int) (*model.Customer, error)
	CountAll() (int64, error)
}

type InvoiceRepo interface {
	Create(tx *gorm.DB, invoice *model.Invoice) (int, error)
	FindById(id int) (*model.Invoice, error)
	FindAll(filter filters.InvoiceFilter, page int, pageSize int) ([]model.Invoice, int64, error)
	UpdateById(tx *gorm.DB, id int, update *model.Invoice) (*model.Invoice, error)
	UpdateStatusPaid(id int, paymentStatus int) (*model.Invoice, error)
	SoftDelete(id int) error
	CountAll() (int64, error)
}

type InvoiceItemRepo interface {
	Create(tx *gorm.DB, invoiceItem *[]model.InvoiceItem) ([]int, error)
	FindById(id int) (*model.InvoiceItem, error)
	FindByInvoiceId(idInvoice int) ([]model.InvoiceItem, error)
	UpdateById(tx *gorm.DB, id int, update *model.InvoiceItem) (*model.InvoiceItem, error)
	SoftDelete(tx *gorm.DB, id int) error
}
