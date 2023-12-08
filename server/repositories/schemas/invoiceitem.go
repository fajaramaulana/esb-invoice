package schemas

import "time"

type InvoiceItem struct {
	ID         uint           `gorm:"primaryKey" json:"id"`
	InvoiceID  uint           `gorm:"not null" json:"invoice_id"`
	ItemID     uint           `gorm:"not null" json:"item_id"`
	Quantity   int            `gorm:"not null" json:"quantity"`
	UnitPrice  float64        `gorm:"not null" json:"unit_price"`
	TotalPrice float64        `gorm:"not null" json:"total_price"`
	CreatedAt  time.Time      `json:"created_at" gorm:"index;default:CURRENT_TIMESTAMP"`
	UpdatedAt  time.Time      `json:"updated_at" gorm:"index;default:CURRENT_TIMESTAMP"`
	DeletedAt  gorm.deletedAt `gorm:"index" json:"deletedAt,omitempty"`
}
