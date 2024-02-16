package domain

type TierManagement struct {
	ID       int `gorm:"primaryKey;column:id;<-:create"`
	Name     string
	MinPoint int64
	MaxPoint int64
	Timestamp
	LoyaltyPrograms []LoyaltyProgram `gorm:"many2many:tier_management_loyalties;foreignKey:id;joinForeignKey:tier_management_id;references:id;joinReferences:loyalty_program_id"`
}

func (t *TierManagement) TableName() string {
	return "tier_managements"
}
