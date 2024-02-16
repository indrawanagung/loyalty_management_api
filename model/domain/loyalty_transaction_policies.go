package domain

import "gorm.io/gorm"

type LoyaltyTransactionPolicies struct {
	ID                int `gorm:"primaryKey;column:id;<-:create"`
	LoyaltyProgramID  int
	TransactionAmount int64
	Qty               int
	FirstPurchase     int64
	FixedPoint        int64
	Percentage        int64
	DeletedAt         gorm.DeletedAt
	LoyaltyProgram    *LoyaltyProgram `gorm:"foreignKey:loyalty_program_id;references:id"`
}

func (l *LoyaltyTransactionPolicies) TableName() string {
	return "loyalty_transaction_policies"
}
