package domain

import "time"

type RedeemedPointHistory struct {
	ID              int `gorm:"primaryKey;column:id;<-:create"`
	MemberID        int
	EarnedPoint     int64
	RedeemedPoint   int64
	RemainingPoint  int64
	TransactionDate time.Time
	Timestamp
	Member *Membership `gorm:"foreignKey:member_id;references:id"`
}

func (r *RedeemedPointHistory) TableName() string {
	return "redeemed_point_histories"
}
