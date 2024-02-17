package repository

import (
	"github.com/gofiber/fiber/v2/log"
	"github.com/indrawanagung/loyalty_management_api/model/domain"
	"gorm.io/gorm"
)

type RedeemedHistoryImpl struct {
}

func NewRedeemedHistory() RedeemedHistoryInterface {
	return &RedeemedHistoryImpl{}
}

func (r RedeemedHistoryImpl) AddRedeemed(db *gorm.DB, redeemed domain.RedeemedPointHistory) {
	err := db.Save(&redeemed).Error
	if err != nil {
		log.Error()
		panic(err)
	}

}

func (r RedeemedHistoryImpl) FindAllByMemberID(db *gorm.DB, memberID int) []domain.RedeemedPointHistory {
	var histories []domain.RedeemedPointHistory
	err := db.Where("member_id = ?", memberID).Find(&histories).Error
	if err != nil {
		log.Error(err)
		panic(err)
	}
	return histories
}
