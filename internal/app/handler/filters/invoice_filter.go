package filters

type InvoiceFilter struct {
	InvoiceID     string `form:"invoice_id" validate:"omitempty"`
	IssueDate     string `form:"issue_date" validate:"omitempty"`
	DueDate       string `form:"due_date" validate:"omitempty"`
	Subject       string `form:"subject" validate:"omitempty"`
	TotalItem     string `form:"total_item" validate:"omitempty"`
	CustomerName  string `form:"customer_name" validate:"omitempty"`
	PaymentStatus int    `form:"payment_status" validate:"omitempty"`
}

func NewInvoiceFilter() *InvoiceFilter {
	return &InvoiceFilter{}
}
