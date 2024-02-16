package domain

type LoyaltyTransactionPolicies struct {
	ID                int `gorm:"primaryKey;column:id;<-:create"`
	LoyaltyProgramID  int
	TransactionAmount int64
	Qty               int
	FirstPurchase     int64
	Timestamp
	LoyaltyProgram *LoyaltyProgram `gorm:"foreignKey:loyalty_program_id;references:id"`
}

func (l *LoyaltyTransactionPolicies) TableName() string {
	return "loyalty_transaction_policies"
}
