package repository

import (
	"github.com/indrawanagung/loyalty_management_api/model/domain"
	"gorm.io/gorm"
)

type EarnedHistoryInterface interface {
	Save(db *gorm.DB, earned domain.EarnedPointHistory)
	FindAllEarnedHistory(db *gorm.DB, memberID int) []domain.EarnedPointHistory
}
