package repository

import (
	"github.com/gofiber/fiber/v2/log"
	"github.com/indrawanagung/loyalty_management_api/model/domain"
	"gorm.io/gorm"
)

type StatusRepositoryImpl struct {
}

func NewStatusRepository() StatusRepositoryInterface {
	return &StatusRepositoryImpl{}
}

func (r StatusRepositoryImpl) Save(db *gorm.DB, status domain.Status) int {
	err := db.Save(&status).Error
	if err != nil {
		log.Error(err)
		panic(err)
	}
	return status.ID
}

func (r StatusRepositoryImpl) FindAll(db *gorm.DB) []domain.Status {
	var status []domain.Status
	err := db.Find(&status).Error
	if err != nil {
		log.Error(err)
		panic(err)
	}
	return status
}
