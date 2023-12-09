package schemas

import (
	"time"

	"gorm.io/gorm"
)

type Invoice struct {
	InvoiceID     uint           `gorm:"primaryKey" json:"invoiceId"`
	Subject       string         `gorm:"not null; type:varchar(255)" json:"subject"`
	IssueDate     time.Time      `gorm:"not null" json:"issueDate"`
	DueDate       time.Time      `gorm:"not null" json:"dueDate"`
	CustomerID    uint           `gorm:"not null" json:"customerId" gorm:"index"`
	PaymentStatus int            `gorm:"not null; comment:'1 = paid, 0 = unpaid'; type:int(1); default:0" json:"paymentStatus"`
	TotalItem     int            `gorm:"not null;" json:"totalItem"`
	Subtotal      float64        `gorm:"not null; type:decimal(10,2)" json:"subtotal"`
	TaxRate       float64        `gorm:"not null; type:decimal(10,2)" json:"taxRate"`
	TaxAmount     float64        `gorm:"not null; type:decimal(10,2)" json:"taxAmount"`
	TotalAmount   float64        `gorm:"not null; type:decimal(10,2)" json:"totalAmount"`
	InvoiceItem   InvoiceItem    `gorm:"foreignKey:InvoiceID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"invoiceItem"`
	Customer      Customer       `gorm:"foreignKey:CustomerID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"customer"`
	CreatedAt     time.Time      `json:"createdAt" gorm:"not null"`
	UpdatedAt     time.Time      `json:"updatedAt" gorm:"not null, autoUpdateTime"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"deletedAt,omitempty"`
}
