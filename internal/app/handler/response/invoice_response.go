package response

type InvoiceResponse struct {
	InvoiceId     uint                  `json:"invoice_id"`
	Subject       string                `json:"subject"`
	IssueDate     string                `json:"issue_date"`
	DueDate       string                `json:"due_date"`
	CustomerId    uint                  `json:"customer_id"`
	PaymentStatus int                   `json:"payment_status"`
	TotalItem     int                   `json:"total_item"`
	Subtotal      float64               `json:"subtotal"`
	TaxRate       float64               `json:"tax_rate"`
	TaxAmount     float64               `json:"tax_amount"`
	TotalAmount   float64               `json:"total_amount"`
	InvoiceItem   []ItemInsideInvoice   `json:"invoice_item"`
	Customer      CustomerInsideInvoice `json:"customer"`
}

type ItemInsideInvoice struct {
	ID         uint                    `json:"id"`
	InvoiceID  uint                    `json:"invoice_id"`
	ItemID     uint                    `json:"item_id"`
	Quantity   int                     `json:"quantity"`
	UnitPrice  float64                 `json:"unit_price"`
	TotalPrice float64                 `json:"total_price"`
	Item       DetailItemInsideInvoice `json:"item"`
}

type DetailItemInsideInvoice struct {
	ID     uint              `json:"id"`
	Name   string            `json:"name"`
	Price  float64           `json:"price"`
	TypeID uint              `json:"typeId"`
	Type   TypeInsideInvoice `json:"type"`
}

type TypeInsideInvoice struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type CustomerInsideInvoice struct {
	ID      uint   `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Address string `json:"address"`
}
