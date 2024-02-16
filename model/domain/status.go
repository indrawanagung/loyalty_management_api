package domain

import "gorm.io/gorm"

type Status struct {
	ID        int `gorm:"primaryKey;column:id;<-:create"`
	Name      string
	DeletedAt gorm.DeletedAt
}

func (receiver *Status) TableName() string {
	return "status"
}
