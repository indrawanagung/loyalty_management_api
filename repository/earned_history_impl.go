package repository

import (
	"fmt"
	"github.com/gofiber/fiber/v2/log"
	"github.com/indrawanagung/loyalty_management_api/model/domain"
	"github.com/indrawanagung/loyalty_management_api/model/web"
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

func (e EarnedHistoryImpl) FindAllEarnedHistory(db *gorm.DB, memberID int, filter web.FilterEarnedPoint) []domain.EarnedPointHistory {
	var earnedHistories []domain.EarnedPointHistory
	fmt.Println(filter.LoyaltyName)
	//query := db.Joins("LoyaltyProgram").Joins("Member").Where("earned_point_histories.member_id = ?", memberID)
	query := db.Joins("join loyalty_programs on loyalty_programs.id = earned_point_histories.loyalty_program_id").Joins("LoyaltyProgram").Joins("Member").Where("earned_point_histories.member_id = ?", memberID)

	if filter.LoyaltyName != "" {
		query.Where("loyalty_programs.name = ?", filter.LoyaltyName)
	}

	if filter.ReferenceTransactionID != "" {
		query.Where("earned_point_histories.reference_transaction_id = ?", filter.ReferenceTransactionID)
	}

	if filter.DateRange.StartDate != "" && filter.DateRange.EndDate != "" {
		query.Where("transaction_date between ? and ?", filter.DateRange.StartDate, filter.DateRange.EndDate)
	}

	err := query.Find(&earnedHistories).Error
	if err != nil {
		log.Error(err)
		panic(err)
	}
	fmt.Println(earnedHistories)
	return earnedHistories
}
