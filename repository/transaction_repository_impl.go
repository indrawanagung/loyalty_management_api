package repository

import (
	"github.com/gofiber/fiber/v2/log"
	"github.com/indrawanagung/loyalty_management_api/model/domain"
	"github.com/indrawanagung/loyalty_management_api/util"
	"gorm.io/gorm"
)

type TransactionRepositoryImpl struct {
}

func NewTransactionRepository() TransactionRepositoryInterface {
	return &TransactionRepositoryImpl{}
}

func (t TransactionRepositoryImpl) AddTransaction(db *gorm.DB, itemTransaction domain.ItemTransaction) string {
	err := db.Save(&itemTransaction).Error
	if err != nil {
		log.Error()
		panic(err)
	}

	itemTransaction.TransactionID = util.GenerateIDTransaction(itemTransaction.ID)
	err = db.Save(&itemTransaction).Error
	if err != nil {
		log.Error()
		panic(err)
	}
	return itemTransaction.TransactionID
}
