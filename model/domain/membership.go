package domain

import "time"

type Membership struct {
	ID            int `gorm:"primaryKey;column:id;<-:create"`
	Name          string
	Email         string
	Password      string
	Phone         string
	BirthDate     time.Time
	Address       string
	JoinDate      time.Time
	Referral      string
	EarnedPoint   int64
	RedeemedPoint int64
	RemainedPoint int64
	StatusID      int
	Timestamp
	Status *Status `gorm:"foreignKey:status_id;references:id"`
}

func (m *Membership) TableName() string {
	return "memberships"
}
