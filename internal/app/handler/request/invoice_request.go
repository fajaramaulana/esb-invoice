package request

type CreateInvoiceRequest struct {
	Subject string `json:"subject" form:"subject" binding:"required" validate:"required,min=3,max=255"`
	// i want to validate this date format 01/12/2021 for issuedate
	// 01/12/2021 for duedate
	// 01/12/2021 for invoice_items
	IssueDate    string                     `json:"issue_date" form:"issue_date" binding:"required" validate:"required,customDate=01/12/2021"`
	DueDate      string                     `json:"due_date" form:"due_date" binding:"required" validate:"required,customDate=01/12/2021"`
	CustomerId   int                        `json:"customer_id" form:"customer_id" binding:"required" validate:"required,numeric"`
	TotalItem    int                        `json:"total_item" form:"total_item" binding:"required" validate:"required,numeric"`
	SubTotal     float64                    `json:"sub_total" form:"sub_total" binding:"required" validate:"required,numeric"`
	Tax          float64                    `json:"tax" form:"tax" binding:"required" validate:"required,numeric"`
	TaxAmount    float64                    `json:"tax_amount" form:"tax_amount" binding:"required" validate:"required,numeric"`
	TotalAmount  float64                    `json:"total_amount" form:"total_amount" binding:"required" validate:"required,numeric"`
	InvoiceItems []CreateInvoiceItemRequest `json:"invoice_items" form:"invoice_items" binding:"required" validate:"required"`
}

type CreateInvoiceItemRequest struct {
	ProductId   int     `json:"product_id" form:"product_id" binding:"required" validate:"required"`
	ProductName string  `json:"product_name" form:"product_name" binding:"required" validate:"required"`
	Quantity    int     `json:"quantity" form:"quantity" binding:"required" validate:"required"`
	Price       float64 `json:"price" form:"price" binding:"required" validate:"required"`
	TotalPrice  float64 `json:"total_price" form:"total_price" binding:"required" validate:"required"`
}
