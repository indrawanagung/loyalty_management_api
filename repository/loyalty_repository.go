package repository

import (
	"github.com/indrawanagung/loyalty_management_api/model/domain"
	"gorm.io/gorm"
)

type LoyaltyRepositoryInterface interface {
	FindByID(db *gorm.DB, loyaltyID int) (domain.LoyaltyProgram, error)
}
