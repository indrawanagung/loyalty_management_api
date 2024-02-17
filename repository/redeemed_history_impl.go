package repository

import (
	"github.com/gofiber/fiber/v2/log"
	"github.com/indrawanagung/loyalty_management_api/model/domain"
	"github.com/indrawanagung/loyalty_management_api/model/web"
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

func (r RedeemedHistoryImpl) FindAllByMemberID(db *gorm.DB, memberID int, filter web.FilterRedeemedPoint) []domain.RedeemedPointHistory {
	var histories []domain.RedeemedPointHistory
	query := db.Joins("Member").Where("redeemed_point_histories.member_id = ?", memberID)

	if filter.DateRange.StartDate != "" && filter.DateRange.EndDate != "" {
		query.Where("transaction_date between ? and ?", filter.DateRange.StartDate, filter.DateRange.EndDate)
	}

	err := query.Find(&histories).Error
	if err != nil {
		log.Error(err)
		panic(err)
	}
	return histories
}
