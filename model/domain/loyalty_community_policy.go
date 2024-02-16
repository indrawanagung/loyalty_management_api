package domain

type LoyaltyCommunityPolicies struct {
	ID               int `gorm:"primaryKey;column:id;<-:create"`
	LoyaltyProgramID int
	Name             string
	Timestamp
	LoyaltyProgram *LoyaltyProgram `gorm:"foreignKey:loyalty_program_id;references:id"`
}
