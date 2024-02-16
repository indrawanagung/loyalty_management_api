package repository

import (
	"github.com/indrawanagung/loyalty_management_api/model/domain"
	"gorm.io/gorm"
)

type StatusRepositoryInterface interface {
	Save(db *gorm.DB, status domain.Status) int
	FindAll(db *gorm.DB) []domain.Status
}
