package repository

import (
	"github.com/indrawanagung/loyalty_management_api/model/domain"
	"github.com/indrawanagung/loyalty_management_api/model/web"
	"gorm.io/gorm"
)

type RedeemedHistoryInterface interface {
	AddRedeemed(db *gorm.DB, redeemed domain.RedeemedPointHistory)
	FindAllByMemberID(db *gorm.DB, memberID int, filter web.FilterRedeemedPoint) []domain.RedeemedPointHistory
}
