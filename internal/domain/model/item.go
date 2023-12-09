package model

import (
	"time"

	"gorm.io/gorm"
)

type Item struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	Name      string         `gorm:"not null" json:"name"`
	Price     float64        `gorm:"not null; type:decimal(10,2)" json:"price"`
	TypeID    uint           `gorm:"not null" json:"typeId" gorm:"index"`
	Type      Type           `gorm:"foreignKey:TypeID" json:"type"`
	CreatedAt time.Time      `json:"createdAt" gorm:"not null"`
	UpdatedAt time.Time      `json:"updatedAt" gorm:"not null, autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deletedAt,omitempty"`
}
