package domain

import "time"

type EarnedPointHistory struct {
	ID                     int `gorm:"primaryKey;column:id;<-:create"`
	MemberID               int
	TransactionDate        time.Time
	ReferenceTransactionID int
	LoyaltyProgramID       int
	ExistingPoint          int64
	EarnedPoint            int64
	BalancePoint           int64
	Timestamp
	Member          *Membership      `gorm:"foreignKey:member_id;references:id"`
	ItemTransaction *ItemTransaction `gorm:"foreignKey:item_transaction_id;references:id"`
	LoyaltyProgram  *LoyaltyProgram  `gorm:"foreignKey:loyalty_program_id;references:id"`
}

func (e *EarnedPointHistory) TableName() string {
	return "earned_point_histories"
}
