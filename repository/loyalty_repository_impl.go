package repository

import (
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2/log"
	"github.com/indrawanagung/loyalty_management_api/model/domain"
	"gorm.io/gorm"
)

type LoyaltyRepositoryImpl struct {
}

func NewLoyaltyRepository() LoyaltyRepositoryInterface {
	return &LoyaltyRepositoryImpl{}
}

func (l LoyaltyRepositoryImpl) FindByID(db *gorm.DB, loyaltyID int) (domain.LoyaltyProgram, error) {
	var loyalty domain.LoyaltyProgram
	err := db.Joins("TransactionPolicy").Joins("CommunityPolicy").Take(&loyalty, "loyalty_programs.id = ?", loyaltyID).Error
	if err != nil {
		if err.Error() != "record not found" {
			log.Error(err)
			panic(err)
		}
		return loyalty, errors.New(fmt.Sprintf("loyalty is not found"))
	}
	return loyalty, nil
}
