package schemas

import (
	"time"

	"gorm.io/gorm"
)

type Type struct {
	ID        int            `gorm:"primaryKey;autoIncrement" json:"id"`
	Name      string         `gorm:"not null; type:varchar(255)" json:"name"`
	CreatedAt time.Time      `json:"createdAt" gorm:"not null"`
	UpdatedAt time.Time      `json:"updatedAt" gorm:"not null, autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deletedAt,omitempty"`
}
