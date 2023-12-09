package response

type CreateCustomerResponse struct {
	Id int `json:"id"`
}

type ReturnResponseCreate struct {
	Meta         Meta                   `json:"meta"`
	Data         CreateCustomerResponse `json:"data"`
	TotalRecords int                    `json:"totalRecords"`
}

type UpdateCustomerResponse struct {
	Id int `json:"id"`
}

type ReturnResponseUpdate struct {
	Meta         Meta                   `json:"meta"`
	Data         UpdateCustomerResponse `json:"data"`
	TotalRecords int                    `json:"totalRecords"`
}

type DeleteCustomerResponse struct {
	Id int `json:"id"`
}

type ReturnResponseDelete struct {
	Meta         Meta                   `json:"meta"`
	Data         DeleteCustomerResponse `json:"data"`
	TotalRecords int                    `json:"totalRecords"`
}

type GetByIdCustomer struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Address string `json:"address"`
}

type ReturnResponseGetById struct {
	Meta         Meta            `json:"meta"`
	Data         GetByIdCustomer `json:"data"`
	TotalRecords int             `json:"totalRecords"`
}

type GetAllCustomer struct {
	Data []GetByIdCustomer `json:"data"`
}

type ReturnResponseGetAll struct {
	Meta         Meta           `json:"meta"`
	Data         GetAllCustomer `json:"data"`
	TotalRecords int            `json:"totalRecords"`
}
