package repository

import (
	"github.com/indrawanagung/loyalty_management_api/model/domain"
	"gorm.io/gorm"
)

type TierManagementRepositoryInterface interface {
	FindAll(db *gorm.DB) []domain.TierManagement
	FindByID(db *gorm.DB, ID int) (domain.TierManagement, error)
	FindByPoint(db *gorm.DB, point int64) ([]domain.LoyaltyProgram, error)
	SaveOrUpdate(db *gorm.DB, tier domain.TierManagement) int
	Delete(db *gorm.DB, ID int)
}
