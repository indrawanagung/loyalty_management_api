package domain

type ItemTransactionDetails struct {
	ID                int `gorm:"primaryKey;column:id;<-:create"`
	ItemTransactionID int
	ItemName          string
	Price             int64
	Qty               int
	ItemSubtotal      int64
	CreatedAt         string
	ItemTransaction   *ItemTransaction `gorm:"foreignKey:item_transaction_id;references:id"`
}

func (i *ItemTransactionDetails) TableName() string {
	return "item_transaction_details"
}
