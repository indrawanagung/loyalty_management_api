package domain

type TierManagementLoyalty struct {
	TierManagementID int             `gorm:"primary_key"`
	LoyaltyProgramID int             `gorm:"primary_key"`
	TierManagement   *TierManagement `gorm:"foreignKey:tier_management_id;references:id"`
	LoyaltyProgram   *LoyaltyProgram `gorm:"foreignKey:loyalty_program_id;references:id"`
}

func (t *TierManagementLoyalty) TableName() string {
	return "tier_management_loyalties"
}
