package domain

type LoyaltyProgram struct {
	ID   int `gorm:"primaryKey;column:id;<-:create"`
	Name string
	Timestamp
	CommunityPolicy   LoyaltyCommunityPolicies   `gorm:"foreignKey:loyalty_program_id;references:id"`
	TransactionPolicy LoyaltyTransactionPolicies `gorm:"foreignKey:loyalty_program_id;references:id"`
	TierManagements   []TierManagement           `gorm:"many2many:tier_management_loyalties;foreignKey:id;joinForeignKey:loyalty_program_id;references:id;joinReferences:tier_management_id"`
}

func (l *LoyaltyProgram) TableName() string {
	return "loyalty_programs"
}
