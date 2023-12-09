package repositories

import (
	"esb-invoice/server/repositories/filters"
	"esb-invoice/server/repositories/schemas"
)

type TypeRepo interface {
	Create(typeItem *schemas.Type) (int, error)
	FindAll() ([]schemas.Type, error)
	FindById(id int) (*schemas.Type, error)
	UpdateById(id int, update *schemas.Type) (*schemas.Type, error)
	SoftDelete(id int) error
}

type ItemRepo interface {
	Create(item *schemas.Item) (int, error)
	FindById(id int) (*schemas.Item, error)
	FindAll() ([]schemas.Item, error)
	UpdateById(id int, update *schemas.Item) (*schemas.Item, error)
	SoftDelete(id int) error
}

type CustomerRepo interface {
	Create(item *schemas.Customer) (int, error)
	FindById(id int) (*schemas.Customer, error)
	FindAll(filter filters.CustomerFilter, page int, pageSize int) ([]schemas.Customer, int64, error)
	UpdateById(id int, update *schemas.Customer) (*schemas.Customer, error)
	SoftDelete(id int) error
}

type InvoiceRepo interface {
	Create(invoice *schemas.Invoice) (int, error)
	FindById(id int) (*schemas.Invoice, error)
	FindAll(filter filters.InvoiceFilter, page int, pageSize int) ([]schemas.Invoice, int64, error)
	UpdateById(id int, update *schemas.Invoice) (*schemas.Invoice, error)
	UpdateStatusPaid(id int, paymentStatus int) (*schemas.Invoice, error)
	SoftDelete(id int) error
}

type InvoiceItemRepo interface {
	Create(invoiceItem *[]schemas.InvoiceItem) ([]int, error)
	FindById(id int) (*schemas.InvoiceItem, error)
	FindByInvoiceId(idInvoice int) ([]schemas.InvoiceItem, error)
	UpdateById(id int, update *schemas.InvoiceItem) (*schemas.InvoiceItem, error)
	SoftDelete(id int) error
}
