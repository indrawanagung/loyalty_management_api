package repository

import (
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2/log"
	"github.com/indrawanagung/loyalty_management_api/model/domain"
	"gorm.io/gorm"
)

type TierManagementRepositoryImpl struct {
}

func NewTierManagementRepository() TierManagementRepositoryInterface {
	return &TierManagementRepositoryImpl{}
}

func (t TierManagementRepositoryImpl) FindAll(db *gorm.DB) []domain.TierManagement {
	var tierManagements []domain.TierManagement
	err := db.Find(&tierManagements).Error
	if err != nil {
		log.Error(err)
		panic(err)
	}
	return tierManagements
}

func (t TierManagementRepositoryImpl) FindByPoint(db *gorm.DB, point int64) ([]domain.LoyaltyProgram, error) {
	var loyalties []domain.LoyaltyProgram
	var tier domain.TierManagement
	err := db.Take(&tier, "min_point <= ? and ? <= max_point", point, point).Error
	if err != nil {
		if err.Error() != "record not found" {
			log.Error(err)
			panic(err)
		}
		return loyalties, errors.New(fmt.Sprintf("tier range point is not found"))
	}

	err = db.Model(&tier).Association("LoyaltyPrograms").Find(&loyalties)
	if err != nil {
		if err.Error() != "record not found" {
			log.Error(err)
			panic(err)
		}
		return loyalties, errors.New(fmt.Sprintf("tier range point is not found"))
	}

	//Joins("LoyaltyProgram.LoyaltyTransactionPolicies")

	return loyalties, nil
}

func (t TierManagementRepositoryImpl) FindByID(db *gorm.DB, ID int) (domain.TierManagement, error) {
	var tierManagement domain.TierManagement
	err := db.First(&tierManagement, "id = ?", ID).Error
	if err != nil {
		if err.Error() != "record not found" {
			log.Error(err)
			panic(err)
		}
		return tierManagement, errors.New(fmt.Sprintf("tier id is is not found"))
	}

	return tierManagement, nil
}

func (t TierManagementRepositoryImpl) SaveOrUpdate(db *gorm.DB, tier domain.TierManagement) int {
	err := db.Save(&tier).Error
	if err != nil {
		log.Error(err)
		panic(err)
	}
	return tier.ID
}

func (t TierManagementRepositoryImpl) Delete(db *gorm.DB, ID int) {
	err := db.Delete(&domain.TierManagement{ID: ID}).Error
	if err != nil {
		log.Error(err)
		panic(err)
	}

}
