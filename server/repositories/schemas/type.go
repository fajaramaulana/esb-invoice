package schemas

import (
	"time"

	"gorm.io/gorm"
)

type Type struct {
	ID        int            `gorm:"primaryKey;autoIncrement" json:"id"`
	Name      string         `gorm:"not null; type:varchar(255)" json:"name"`
	CreatedAt time.Time      `json:"createdAt" gorm:"index;default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time      `json:"updatedAt" gorm:"index;default:CURRENT_TIMESTAMP"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deletedAt,omitempty"`
}
