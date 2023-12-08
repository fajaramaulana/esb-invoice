package schemas

import (
	"time"

	"gorm.io/gorm"
)

type Item struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	Name      string         `gorm:"not null" json:"name"`
	Price     float64        `gorm:"not null" json:"price"`
	TypeID    uint           `gorm:"not null" json:"typeId" gorm:"index"`
	Type      Type           `gorm:"foreignKey:TypeID" json:"type"`
	CreatedAt time.Time      `json:"createdAt" gorm:"index;default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time      `json:"updatedAt" gorm:"index;default:CURRENT_TIMESTAMP"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deletedAt,omitempty"`
}
