package domain

import "time"

type ItemTransaction struct {
	ID              int `gorm:"primaryKey;column:id;<-:create"`
	MemberID        int
	TotalAmount     int64
	TransactionDate time.Time
	TransactionID   string
	Timestamp
	Membership *Membership `gorm:"foreignKey:member_id;references:id"`
}

func (i *ItemTransaction) TableName() string {
	return "item_transactions"
}
