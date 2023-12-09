package filters

type InvoiceFilter struct {
	InvoiceID  uint   `json:"invoiceID"`
	IssueDate  string `json:"issueDate"`
	DueDate    string `json:"dueDate"`
	Subject    string `json:"subject"`
	TotalItems int    `json:"totalItems"`
	CustomerID uint   `json:"customerID"`
	UserID     uint   `json:"userID"`
	Status     string `json:"status"`
}

func NewInvoiceFilter() *InvoiceFilter {
	return &InvoiceFilter{}
}
