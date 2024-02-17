package repository

import (
	"github.com/indrawanagung/loyalty_management_api/model/domain"
	"gorm.io/gorm"
)

type RedeemedHistoryInterface interface {
	AddRedeemed(db *gorm.DB, redeemed domain.RedeemedPointHistory)
	FindAllByMemberID(db *gorm.DB, memberID int) []domain.RedeemedPointHistory
}
