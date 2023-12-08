package schemas

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	Id        int            `gorm:"primaryKey;autoIncrement" json:"id"`
	Name      string         `gorm:"not null; type:varchar(255)" json:"name"`
	Email     string         `gorm:"not null; unique; type:varchar(100); index" json:"email"`
	Address   string         `gorm:"not null; unique; type:TEXT" json:"address"`
	CreatedAt time.Time      `json:"createdAt" gorm:"index;default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time      `json:"updatedAt" gorm:"index;default:CURRENT_TIMESTAMP"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deletedAt,omitempty"`
	UpdatedBy uint           `gorm:"not null; default: 0" json:"updatedBy"`
}
