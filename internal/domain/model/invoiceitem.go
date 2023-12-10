package model

import (
	"time"

	"gorm.io/gorm"
)

type InvoiceItem struct {
	ID         uint           `gorm:"primaryKey" json:"id"`
	InvoiceID  uint           `gorm:"not null" json:"invoice_id"`
	ItemID     uint           `gorm:"not null" json:"item_id"`
	Quantity   int            `gorm:"not null" json:"quantity"`
	UnitPrice  float64        `gorm:"not null; type:decimal(10,2)" json:"unit_price"`
	TotalPrice float64        `gorm:"not null; type:decimal(10,2)" json:"total_price"`
	Item       Item           `gorm:"foreignKey:ItemID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"item"`
	CreatedAt  time.Time      `json:"createdAt" gorm:"not null"`
	UpdatedAt  time.Time      `json:"updatedAt" gorm:"not null, autoUpdateTime"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"deletedAt,omitempty"`
}
