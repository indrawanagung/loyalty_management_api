package domain

import "gorm.io/gorm"

type Timestamp struct {
	CreatedAt string         `gorm:"column:created_at"`
	UpdatedAt string         `gorm:"column:updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at"`
}
