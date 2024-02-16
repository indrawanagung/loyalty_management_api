package web

type TransactionCreateRequest struct {
	Transaction []ItemTransactionRequest `json:"item_transaction"`
}

type ItemTransactionRequest struct {
	ItemName string `validate:"required" json:"item_name"`
	Price    int64  `validate:"required, gte=1000, lte=10000000" json:"price"`
	Qty      int    `validate:"required, gt=0" json:"qty"`
}
