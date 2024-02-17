package repository

import (
	"github.com/indrawanagung/loyalty_management_api/model/domain"
	"gorm.io/gorm"
)

type MembershipRepositoryInterface interface {
	FindByID(db *gorm.DB, ID int) (domain.Membership, error)
	FindByEmail(db *gorm.DB, email string) (domain.Membership, error)
	Save(db *gorm.DB, membership domain.Membership) int
	AddMemberHistory(db *gorm.DB, history domain.AddMemberHistory) string
	AddMemberActivityHistory(db *gorm.DB, history domain.AddMemberActivityHistory) string
}
