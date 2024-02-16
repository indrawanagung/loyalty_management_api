package domain

import "time"

type AddMemberActivityHistory struct {
	ID              int `gorm:"primaryKey;column:id;<-:create"`
	MemberID        int
	ActivityName    string
	TransactionDate time.Time
	TransactionID   string
	Timestamp
	Member *Membership `gorm:"foreignKey:member_id;references:id"`
}

func (a *AddMemberActivityHistory) TableName() string {
	return "add_member_activity_histories"
}
