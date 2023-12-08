package schemas

import (
	"time"

	"gorm.io/gorm"
)

type Invoice struct {
	InvoiceID  uint           `gorm:"primaryKey" json:"invoiceId"`
	Subject    string         `gorm:"not null; type:varchar(255)" json:"subject"`
	IssueDate  time.Time      `gorm:"not null" json:"issueDate"`
	DueDate    time.Time      `gorm:"not null" json:"dueDate"`
	CustomerID uint           `gorm:"not null" json:"customerId" gorm:"index"`
	UserID     uint           `gorm:"not null" json:"userId" gorm:"index"`
	CreatedAt  time.Time      `json:"createdAt" gorm:"index;default:CURRENT_TIMESTAMP"`
	UpdatedAt  time.Time      `json:"updatedAt" gorm:"index;default:CURRENT_TIMESTAMP"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"deletedAt,omitempty"`
	UpdatedBy  uint           `gorm:"not null; default: 0" json:"updatedBy"`
	Customer   Customer       `gorm:"foreignKey:CustomerID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"customer"`
}
