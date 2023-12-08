package schemas

import (
	"time"

	"gorm.io/gorm"
)

type Invoice struct {
	InvoiceID   uint           `gorm:"primaryKey" json:"invoiceId"`
	Subject     string         `gorm:"not null; type:varchar(255)" json:"subject"`
	IssueDate   time.Time      `gorm:"not null" json:"issueDate"`
	DueDate     time.Time      `gorm:"not null" json:"dueDate"`
	CustomerID  uint           `gorm:"not null" json:"customerId" gorm:"index"`
	Subtotal    float64        `gorm:"not null" json:"subtotal"`
	TaxRate     float64        `gorm:"not null" json:"taxRate"`
	TaxAmount   float64        `gorm:"not null" json:"taxAmount"`
	TotalAmount float64        `gorm:"not null" json:"totalAmount"`
	DetailItem  DetailItem     `gorm:"foreignKey:DetailitemID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"detailItem"`
	Customer    Customer       `gorm:"foreignKey:CustomerID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"customer"`
	CreatedAt   time.Time      `json:"createdAt" gorm:"index;default:CURRENT_TIMESTAMP"`
	UpdatedAt   time.Time      `json:"updatedAt" gorm:"index;default:CURRENT_TIMESTAMP"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deletedAt,omitempty"`
}
