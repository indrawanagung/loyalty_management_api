package repository

import (
	"github.com/gofiber/fiber/v2/log"
	"github.com/indrawanagung/loyalty_management_api/model/domain"
	"gorm.io/gorm"
)

type EarnedHistoryImpl struct {
}

func NewEarnedHistory() EarnedHistoryInterface {
	return &EarnedHistoryImpl{}
}

func (e EarnedHistoryImpl) Save(db *gorm.DB, earned domain.EarnedPointHistory) {
	err := db.Save(&earned).Error
	if err != nil {
		log.Error(err)
		panic(err)
	}
}
