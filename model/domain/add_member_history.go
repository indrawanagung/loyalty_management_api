package domain

import "time"

type AddMemberHistory struct {
	ID              int `gorm:"primaryKey;column:id;<-:create"`
	MemberID        int
	PersonID        int
	TransactionDate time.Time
	TransactionID   string
	Timestamp
	Member *Membership `gorm:"foreignKey:member_id;references:id"`
	Person *Membership `gorm:"foreignKey:person_id;references:id"`
}

func (a *AddMemberHistory) TableName() string {
	return "add_member_histories"
}
