package service

import (
	"esb-invoice/internal/app/handler/filters"
	"esb-invoice/internal/app/handler/request"
	"esb-invoice/internal/app/handler/response"
	"esb-invoice/internal/domain/model"
	"esb-invoice/internal/domain/repository"
	"time"
)

type CustomerService struct {
	customerRepo repository.CustomerRepo
}

func NewCustomerService(customerRepo repository.CustomerRepo) *CustomerService {
	return &CustomerService{
		customerRepo: customerRepo,
	}
}

func (s *CustomerService) Create(req *request.CreateCustomerRequest) (int, error) {
	var customer model.Customer

	customer.Name = req.Name
	customer.Email = req.Email
	customer.Address = req.Address
	customer.CreatedAt = time.Now()
	customer.UpdatedAt = time.Time{}

	customerId, err := s.customerRepo.Create(&customer)

	if err != nil {
		return 0, err
	}

	return customerId, nil
}

func (s *CustomerService) FindById(id int) (*model.Customer, error) {
	customer, err := s.customerRepo.FindById(id)

	if err != nil {
		return nil, err
	}

	return customer, nil
}

func (s *CustomerService) FindAll(filter *filters.CustomerFilter, page int, pageSize int) ([]response.GetByIdCustomer, int64, error) {
	var customers []model.Customer
	var res []response.GetByIdCustomer

	customers, totalRecords, err := s.customerRepo.FindAll(*filter, page, pageSize)

	if err != nil {
		return nil, 0, err
	}

	for _, customer := range customers {

		res = append(res, response.GetByIdCustomer{
			Id:      customer.ID,
			Name:    customer.Name,
			Email:   customer.Email,
			Address: customer.Address,
		})
	}

	return res, totalRecords, nil
}

func (s *CustomerService) CountAll() (int64, error) {
	totalRecords, err := s.customerRepo.CountAll()

	if err != nil {
		return 0, err
	}

	return totalRecords, nil
}

func (s *CustomerService) UpdateById(id int, req *request.UpdateCustomerRequest) (*model.Customer, error) {
	var customer model.Customer
	customer.Name = req.Name
	customer.Email = req.Email
	customer.Address = req.Address
	customer.UpdatedAt = time.Now()

	updatedCustomer, err := s.customerRepo.UpdateById(id, &customer)
	if err != nil {
		return nil, err
	}

	customer = *updatedCustomer

	return &customer, nil
}

func (s *CustomerService) SoftDelete(id int) error {
	err := s.customerRepo.SoftDelete(id)

	if err != nil {
		return err
	}

	return nil
}

func (s *CustomerService) FindByEmailAndNotId(email string, id int) (*model.Customer, error) {
	customer, err := s.customerRepo.FindByEmailAndNotId(email, id)

	if err != nil {
		return nil, err
	}

	return customer, nil
}

func (s *CustomerService) FindByEmail(email string) (*model.Customer, error) {
	customer, err := s.customerRepo.FindByEmail(email)

	if err != nil {
		return nil, err
	}

	return customer, nil
}
