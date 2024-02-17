package repository

import (
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2/log"
	"github.com/indrawanagung/loyalty_management_api/model/domain"
	"github.com/indrawanagung/loyalty_management_api/util"
	"gorm.io/gorm"
)

type MembershipRepositoryImpl struct {
}

func NewMembershipRepository() MembershipRepositoryInterface {
	return &MembershipRepositoryImpl{}
}

func (m MembershipRepositoryImpl) FindByID(db *gorm.DB, ID int) (domain.Membership, error) {
	var member domain.Membership
	err := db.Where("id = ? ", ID).Find(&member).Error
	if err != nil {
		if err.Error() != "record not found" {
			log.Error(err)
			panic(err)
		}
		return member, errors.New(fmt.Sprintf("member id is is not found"))
	}
	return member, nil
}

func (m MembershipRepositoryImpl) FindByEmail(db *gorm.DB, email string) (domain.Membership, error) {
	var member domain.Membership
	err := db.Where("email = ? ", email).Find(&member).Error
	if err != nil {
		if err.Error() != "record not found" {
			log.Error(err)
			panic(err)
		}
		return member, errors.New(fmt.Sprintf("member id is is not found"))
	}
	return member, nil
}

func (m MembershipRepositoryImpl) Save(db *gorm.DB, membership domain.Membership) int {
	err := db.Save(&membership).Error
	if err != nil {
		log.Error(err)
		panic(err)
	}
	return membership.ID
}

func (m MembershipRepositoryImpl) AddMemberHistory(db *gorm.DB, history domain.AddMemberHistory) string {
	err := db.Save(&history).Error
	if err != nil {
		log.Error(err)
		panic(err)
	}
	history.TransactionID = util.GenerateIDGetMember(history.ID)
	err = db.Save(&history).Error
	if err != nil {
		log.Error(err)
		panic(err)
	}
	return history.TransactionID
}

func (m MembershipRepositoryImpl) AddMemberActivityHistory(db *gorm.DB, history domain.AddMemberActivityHistory) string {
	err := db.Save(&history).Error
	if err != nil {
		log.Error(err)
		panic(err)
	}

	history.TransactionID = util.GenerateIDActivityMember(history.ID)
	err = db.Save(&history).Error
	if err != nil {
		log.Error(err)
		panic(err)
	}
	return history.TransactionID
}
