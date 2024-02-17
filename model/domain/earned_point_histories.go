package domain

import "time"

type EarnedPointHistory struct {
	ID                     int `gorm:"primaryKey;column:id;<-:create"`
	MemberID               int
	TransactionDate        time.Time
	ReferenceTransactionID string
	LoyaltyProgramID       int
	ExistingPoint          int64
	EarnedPoint            int64
	BalancePoint           int64
	Timestamp
	Member         *Membership     `gorm:"foreignKey:member_id;references:id"`
	LoyaltyProgram *LoyaltyProgram `gorm:"foreignKey:loyalty_program_id;references:id"`
}

func (e *EarnedPointHistory) TableName() string {
	return "earned_point_histories"
}
