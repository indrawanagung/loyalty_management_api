package repository

import (
	"github.com/indrawanagung/loyalty_management_api/model/domain"
	"gorm.io/gorm"
)

type TransactionRepositoryInterface interface {
	AddTransaction(db *gorm.DB, itemTransaction domain.ItemTransaction) string
}
