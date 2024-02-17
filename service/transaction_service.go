package service

import "github.com/indrawanagung/loyalty_management_api/model/web"

type TransactionServiceInterface interface {
	AddTransaction(request web.TransactionCreateRequest, memberID int) string
}
